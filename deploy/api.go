package deploy

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	deploymentspb "github.com/nitrictech/nitric/core/pkg/proto/deployments/v1"
	"github.com/nitrictech/terraform-gcp-extension/deploy/generated/extapi"
)

func (e *ExtendedGcpProvider) Api(stack cdktf.TerraformStack, name string, config *deploymentspb.Api) error {
	// Replace the implementation of the `Api` method with your own logic
	// For the reference implementation, see
	// https://github.com/nitrictech/nitric/blob/main/cloud/gcp/deploy/api.go
	// AND
	// https://github.com/nitrictech/nitric/tree/main/cloud/gcp/deploytf/.nitric/modules/api
	e.Apis[name] = extapi.NewExtapi(stack, jsii.String(name), &extapi.ExtapiConfig{})

	// May also need to extend the gateway plugin implementation to support the new API routing
	// OR have the API routing conform to the base GCP gateway plugin

	return nil
}
