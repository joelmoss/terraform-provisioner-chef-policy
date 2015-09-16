package main

import (
  "github.com/joelmoss/terraform-provisioner-chef-policy/chef-policy"
  "github.com/hashicorp/terraform/plugin"
  "github.com/hashicorp/terraform/terraform"
)

func main() {
  plugin.Serve(&plugin.ServeOpts{
    ProvisionerFunc: func() terraform.ResourceProvisioner {
      return new(chef.ResourceProvisioner)
    },
  })
}
