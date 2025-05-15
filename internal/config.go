package internal

type ContentfulConfig struct {
	CMAToken       string `mapstructure:"cma_token"`
	OrganizationID string `mapstructure:"organization_id"`
}

func (c *ContentfulConfig) extendConfig(o *ContentfulConfig) *ContentfulConfig {
	cfg := &ContentfulConfig{
		CMAToken:       o.CMAToken,
		OrganizationID: o.OrganizationID,
	}
	if c.CMAToken != "" {
		cfg.CMAToken = c.CMAToken
	}
	if c.OrganizationID != "" {
		cfg.OrganizationID = c.OrganizationID
	}
	return cfg
}
