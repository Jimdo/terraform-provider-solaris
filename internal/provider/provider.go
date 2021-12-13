package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jimdo-fs/terraform-provider-solarisbank/internal/solaris"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"endpoint": {
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("SOLARISBANK_ENDPOINT", nil),
				},
				"client_id": {
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("SOLARISBANK_CLIENT_ID", nil),
				},
				"client_secret": {
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("SOLARISBANK_CLIENT_SECRET", nil),
				},
			},
			ResourcesMap: map[string]*schema.Resource{
				"solarisbank_webhook": resourceWebhook(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, data *schema.ResourceData) (interface{}, diag.Diagnostics) {
		endpoint := data.Get("endpoint").(string)
		clientID := data.Get("client_id").(string)
		clientSecret := data.Get("client_secret").(string)

		// we pass the background context here because the current
		// context gets cancelled after the provider is configured.
		// However the context that we pass here gets persistet into the
		// solaris clients oauth token source, which needs to be alive
		// (not cancelled) at resource creation time.
		client := solaris.NewClient(context.Background(), solaris.Config{
			Endpoint:     endpoint,
			ClientID:     clientID,
			ClientSecret: clientSecret,
		})

		return client, nil
	}
}
