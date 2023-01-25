resource "solaris_webhook" "identification" {
  event_type = "IDENTIFICATION"
  url        = "https://example.com/identification"
}

output "identification_webhook_secret" {
  value     = solaris_webhook.identification.secret
  sensitive = true
}
