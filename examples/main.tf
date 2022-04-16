terraform {
  required_providers {
    tailscale = {
      version = "0.9.3"
      source  = "github.com/pellegrino/tailscale"
    }
  }
}

provider "tailscale" {
  # api_key = "my-api-key"
  # tailnet = "my tailnet"
}


# resource "tailscale_device" "example" {
#   name = "device.example.com"
# }

#resource "tailscale_acl" "sample_acl" {
#  acl = jsonencode({
#    acls : [
#      {
#        // Allow all users access to all ports.
#        action = "accept",
#        users  = ["*"],
#        ports  = ["*:*"],
#    }],
#  })
#}
#
#
#resource "tailscale_dns_nameservers" "sample_nameservers" {
#  nameservers = [
#    "8.8.8.8",
#    "8.8.4.4",
#  ]
#}
#
#resource "tailscale_dns_preferences" "sample_preferences" {
#  depends_on = [
#    tailscale_dns_nameservers.sample_nameservers,
#  ]
#
#  magic_dns = true
#}
#

#output "sample_acl" {
#  value = tailscale_acl.sample_acl.acl
#}
#
#output "sample_nameservers" {
#  value = tailscale_dns_nameservers.sample_nameservers.nameservers
#}
#
#output "sample_preferences_magic_dns" {
#  value = tailscale_dns_preferences.sample_preferences.magic_dns
#}
