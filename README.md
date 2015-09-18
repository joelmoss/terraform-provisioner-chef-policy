# terraform-provisioner-chef-policy

A [Terraform](https://www.terraform.io) Chef provisioner that supports [Policyfiles](https://github.com/chef/chef-dk/blob/master/POLICYFILE_README.md).

This provisioner is almost identical to the Chef provisioner that is bundled with Terraform, but
with a few additional arguments that allow you to work with Policyfiles.

## Argument Reference

 - `policy_name (string)` - (Optional) The name of the policy to apply to this instance.
 - `policy_group (string)` - (Optional) The group of the policy to apply to this instance.
