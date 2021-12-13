# Terraform Provider Solarisbank

Terraform provider to provision static resources using the Solarisbank API.

Currently supported resources are:
- Webhooks

## Using the provider

First require the provider - TODO: find out how we do this, since the provider will be private.

Now add a provider declaration and configure it:

```terraform
provider "solarisbank" {
  endpoint = var.solarisbank_endpoint[terraform.workspace] # or define SOLARISBANK_ENDPOINT
  client_id = var.solarisbank_client_id # or define SOLARISBANK_CLIENT_ID
  client_secret = var.solarisbank_client_secret # or define SOLARISBANK_CLIENT_SECRET
}
```

Reference the individual resource docs for further information.

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 0.13.x
- [Go](https://golang.org/doc/install) >= 1.17

## Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command: 
```sh
$ go install
```

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `make docs`.

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources. Usually they clean up after *themselves but if they fail, and warn about potential dirty states remotely, *please make sure to clean up manually.

```sh
$ make testacc
```
