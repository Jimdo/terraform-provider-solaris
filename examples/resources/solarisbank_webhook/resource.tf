resource "solarisbank_webhook" "identification" {
  event_type = "IDENTIFICATION"
  url        = "https://example.com/identification"
}

output "identification_webhook_secret" {
  value     = solarisbank_webhook.identification.secret
  sensitive = true
}
