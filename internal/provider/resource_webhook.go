package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jimdo/terraform-provider-solaris/internal/solaris"
)

func resourceWebhook() *schema.Resource {
	return &schema.Resource{
		Description: "Solaris webhook subscription. Reference [the Solaris documentation](https://docs.solarisgroup.com/api-reference/webhooks/) for more information.",

		CreateContext: resourceWebhookCreate,
		ReadContext:   resourceWebhookRead,
		DeleteContext: resourceWebhookDelete,

		Schema: map[string]*schema.Schema{
			"event_type": {
				Description: "The type of event, you want to subscribe to. Values listed [here](https://docs.solarisgroup.com/api-reference/webhooks/#full-list-of-webhook-events).",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"url": {
				Description: "The recipient URL of the event notification.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"secret": {
				Description: "A key phrase used to verify the authenticity of received webhook messages. See [Verification](https://docs.solarisgroup.com/api-reference/webhooks/#content-verification) for further information",
				Type:        schema.TypeString,
				Sensitive:   true,
				Computed:    true,
			},
		},
	}
}

func resourceWebhookCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*solaris.Client)
	eventType := d.Get("event_type").(string)
	url := d.Get("url").(string)

	webhook, err := client.CreateWebhook(ctx, &solaris.CreateWebhookRequest{
		EventType: eventType,
		URL:       url,
	})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(webhook.ID)
	d.Set("event_type", webhook.EventType)
	d.Set("url", webhook.URL)
	d.Set("secret", webhook.Secret)

	return nil
}

func resourceWebhookRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*solaris.Client)
	webhookID := d.Id()

	webhook, err := client.GetWebhook(ctx, webhookID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(webhook.ID)
	d.Set("event_type", webhook.EventType)
	d.Set("url", webhook.URL)

	return nil
}

func resourceWebhookDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*solaris.Client)
	webhookID := d.Id()

	if err := client.DeleteWebhook(ctx, webhookID); err != nil {
		return diag.FromErr(err)
	}
	return nil
}
