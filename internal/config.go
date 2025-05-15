package internal

type ContentfulConfig struct {
	CMAToken       string `mapstructure:"cma_token"`
	OrganizationID string `mapstructure:"organization_id"`
	Environment    string `mapstructure:"environment"`
}

func (c *ContentfulConfig) extendConfig(o *ContentfulConfig) *ContentfulConfig {
	cfg := &ContentfulConfig{
		CMAToken:       o.CMAToken,
		OrganizationID: o.OrganizationID,
	}
	if c.Environment != "" {
		cfg.Environment = c.Environment
	}
	if c.CMAToken != "" {
		cfg.CMAToken = c.CMAToken
	}
	if c.OrganizationID != "" {
		cfg.OrganizationID = c.OrganizationID
	}
	return cfg
}
