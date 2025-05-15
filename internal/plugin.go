package internal

import (
	"fmt"

	"github.com/mach-composer/mach-composer-plugin-helpers/helpers"
	"github.com/mach-composer/mach-composer-plugin-sdk/plugin"
	"github.com/mach-composer/mach-composer-plugin-sdk/schema"
	"github.com/mitchellh/mapstructure"
)

type Plugin struct {
	environment  string
	provider     string
	globalConfig *ContentfulConfig
	siteConfigs  map[string]*ContentfulConfig
	enabled      bool
}

func NewContentfulPlugin() schema.MachComposerPlugin {
	state := &Plugin{
		provider:    "5.3.0",
		siteConfigs: map[string]*ContentfulConfig{},
	}
	return plugin.NewPlugin(&schema.PluginSchema{
		Identifier: "contentful",

		Configure: state.Configure,
		IsEnabled: state.IsEnabled,

		GetValidationSchema: state.GetValidationSchema,

		// Config
		SetGlobalConfig: state.SetGlobalConfig,
		SetSiteConfig:   state.SetSiteConfig,

		// Renders
		RenderTerraformProviders: state.TerraformRenderProviders,
		RenderTerraformResources: state.TerraformRenderResources,
		RenderTerraformComponent: state.RenderTerraformComponent,
	})
}

func (p *Plugin) Configure(environment string, provider string) error {
	p.environment = environment
	if provider != "" {
		p.provider = provider
	}
	return nil
}

func (p *Plugin) IsEnabled() bool {
	return p.enabled
}

func (p *Plugin) GetValidationSchema() (*schema.ValidationSchema, error) {
	result := getSchema()
	return result, nil
}

func (p *Plugin) SetGlobalConfig(data map[string]any) error {
	cfg := ContentfulConfig{}
	if err := mapstructure.Decode(data, &cfg); err != nil {
		return err
	}
	p.globalConfig = &cfg
	p.enabled = true
	return nil
}

func (p *Plugin) SetSiteConfig(site string, data map[string]any) error {
	if data == nil {
		return nil
	}

	cfg := ContentfulConfig{}
	if err := mapstructure.Decode(data, &cfg); err != nil {
		return err
	}
	p.siteConfigs[site] = &cfg
	p.enabled = true
	return nil
}

func (p *Plugin) TerraformRenderStateBackend(site string) (string, error) {
	return "", nil
}

func (p *Plugin) TerraformRenderProviders(site string) (string, error) {
	cfg := p.getSiteConfig(site)
	if cfg == nil {
		return "", nil
	}

	result := fmt.Sprintf(`
		contentful = {
			source = "labd/contentful"
			version = "%s"
		}`, helpers.VersionConstraint(p.provider))
	return result, nil
}

func (p *Plugin) TerraformRenderResources(site string) (string, error) {
	cfg := p.getSiteConfig(site)
	if cfg == nil {
		return "", nil
	}

	template := `
		  provider "contentful" {
			cma_token       = {{ .CMAToken|printf "%q" }}
			organization_id = {{ .OrganizationID|printf "%q" }}
			environment     = {{ .Environment|printf "%q" }}
		  }
	`
	return helpers.RenderGoTemplate(template, cfg)
}

func (p *Plugin) RenderTerraformComponent(site string, component string) (*schema.ComponentSchema, error) {
	cfg := p.getSiteConfig(site)
	if cfg == nil {
		return nil, nil
	}

	result := &schema.ComponentSchema{}
	return result, nil
}

func (p *Plugin) getSiteConfig(site string) *ContentfulConfig {
	if p.globalConfig == nil {
		return nil
	}
	cfg, ok := p.siteConfigs[site]
	if !ok {
		cfg = &ContentfulConfig{}
	}
	return cfg.extendConfig(p.globalConfig)
}
