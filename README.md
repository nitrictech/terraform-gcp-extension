# terraform-gcp-extension

This is a demonstration of how to extend the existing Nitric GCP Terraform provider to make changes to the base deployment.

## Prerequisites

To build this project you will need
 - Go
 - Node.js (for terraform CDK/jsii)
 - Make (for running build scripts)


## Nitric module overrides

If your module can have exactly the same variables and outputs as the original nitric terraform modules then you make provide drop in replacements for the existing modules.

See the `deploy/.nitric/modules/roles` module as an example of how this is done.

For net new infrastructure that we want to replace you can use the method hooks provided by the base terraform provider to provide your own custom implementation of a type of nitric infrastructure. This example provides an empty override for nitric APIs. See `deploy/api.go` and `deploy/.ext/modules/extapi`. 

## Generating Terraform bindings


