# terraform-provider-tailscale

[![Go Reference](https://pkg.go.dev/badge/github.com/pellegrino/terraform-provider-tailscale.svg)](https://pkg.go.dev/github.com/pellegrino/terraform-provider-tailscale)
[![Go Report Card](https://goreportcard.com/badge/github.com/pellegrino/terraform-provider-tailscale)](https://goreportcard.com/report/github.com/pellegrino/terraform-provider-tailscale)
![Github Actions](https://github.com/pellegrino/terraform-provider-tailscale/actions/workflows/ci.yml/badge.svg?branch=master)

This repository contains the source code for the [Terraform Tailscale Provider](https://registry.terraform.io/providers/davidsbond/tailscale)
a Terraform provider implementation for interacting with the [Tailscale](https://tailscale.com) API.

See the [documentation](https://registry.terraform.io/providers/davidsbond/tailscale/latest/docs) in the Terraform registry
for the most up-to-date information and latest release.

This provider is unofficial and not affiliated in any way with the team at Tailscale. I've been given thanks for creating
it, but do not expect anyone who actually works at Tailscale to help you with problems regarding this provider.

## Getting Started

To install this provider, copy and paste this code into your Terraform configuration. Then, run `terraform init`:

```terraform
terraform {
  required_providers {
    tailscale = {
      source = "pellegrino/tailscale"
      version = "0.12.1-device"
    }
  }
}

provider "tailscale" {
  api_key = "my_api_key"
  tailnet = "example.com"
}
```

In the `provider` block, replace `api_key` and `tailnet` with your own tailnet and API key. Alternatively, use the
`TAILSCALE_API_KEY` and `TAILSCALE_TAILNET` environment variables.

The default api endpoint is `https://api.tailscale.com`. If your coordination/control server api is at another endpoint, you can pass in `base_url` in the provider block.

```terraform
provider "tailscale" {
  api_key = "my_api_key"
  tailnet = "example.com"
  base_url = "https://api.us.tailscale.com"
}
```

## Contributing

Please review the [contributing guidelines](./CONTRIBUTING.md) and [code of conduct](.github/CODE_OF_CONDUCT.md) before
contributing to this codebase. Please create a [new issue](https://github.com/pellegrino/terraform-provider-tailscale/issues/new/choose)
for bugs and feature requests and fill in as much detail as you can.
