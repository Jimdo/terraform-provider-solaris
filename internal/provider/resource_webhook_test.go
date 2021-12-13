package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceWebhook(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWebhook,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"solarisbank_webhook.identification", "event_type", "IDENTIFICATION"),
					resource.TestCheckResourceAttr(
						"solarisbank_webhook.identification", "url", "https://httpbin.org/status/200"),
					resource.TestMatchResourceAttr(
						"solarisbank_webhook.identification", "secret", regexp.MustCompile(`\w+`)),
				),
			},
		},
	})
}

const testAccResourceWebhook = `
resource "solarisbank_webhook" "identification" {
  event_type = "IDENTIFICATION"
  url        = "https://httpbin.org/status/200"
}
`
