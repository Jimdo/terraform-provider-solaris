# Solaris Terraform Provider

![Test](https://github.com/Jimdo/terraform-provider-solarisbank/actions/workflows/test.yml/badge.svg)

![Release](https://github.com/Jimdo/terraform-provider-solarisbank/actions/workflows/test.yml/badge.svg)

[![Terraform Registry](https://registry.terraform.io/providers/Jimdo/solaris/latest)

A Terraform provider, the tool for provisioning all your static resources using the Solaris API.

📝 Read [the documentation](https://registry.terraform.io/providers/Jimdo/solaris/latest/docs)

👀 See [example/](examples/)

## Supported Resources
- [Webhooks](https://docs.solarisgroup.com/api-reference/onboarding/webhooks/)

## Using the provider

The provider can be installed directly from the Terraform Registry. To do this, include the following block in your Terraform configuration file. This will download the provider from the Terraform Registry.

```hcl
terraform {
  required_providers {
    solaris = {
      source = "Jimdo/solaris"
      version = "1.0.3"
    }
  }
}

provider "solaris" {
  endpoint = var.solaris_endpoint
  client_id = var.solaris_client_id
  client_secret = var.solaris_client_secret
}
```
You also have the option of setting those values through their respective environment variables:

- `SOLARIS_ENDPOINT`
- `SOLARIS_CLIENT_ID`
- `SOLARIS_CLIENT_ID_SECRET`


For more information, check out the individual resource documents!

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 0.13.x
- [Go](https://golang.org/doc/install) >= 1.19

## Building The Provider

1. Clone the repository
2. Enter the repository directory
3. Build the provider using the Go `install` command: 

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

## License

This software is distributed under the terms of the MIT license, see [LICENSE](./LICENSE) for details.
