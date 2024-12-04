package deploy

import (
	"embed"
	"io/fs"

	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/nitrictech/nitric/cloud/common/deploy/provider"
	"github.com/nitrictech/nitric/cloud/gcp/deploytf"
	deploymentspb "github.com/nitrictech/nitric/core/pkg/proto/deployments/v1"
	"github.com/nitrictech/terraform-gcp-extension/deploy/generated/extapi"
)

type ExtendedGcpProvider struct {
	deploytf.NitricGcpTerraformProvider

	// Record deployed apis
	Apis map[string]extapi.Extapi
}

type mergedFS struct {
	fses []fs.FS
}

func (m mergedFS) Open(name string) (fs.File, error) {
	for _, f := range m.fses {
		file, err := f.Open(name)
		if err == nil {
			return file, nil
		}
	}
	return nil, fs.ErrNotExist
}

// embed the modules directory here
//
//go:embed .ext/modules/**/*
var modules embed.FS

// TODO: Need to update this extension so we can merge the modules directory with the base provider
func (a *ExtendedGcpProvider) CdkTfModules() ([]provider.ModuleDirectory, error) {
	origModules, err := a.CdkTfModules()
	if err != nil {
		return nil, err
	}

	// Merge the original modules with the embedded modules
	mergedModules := append(origModules, provider.ModuleDirectory{
		ParentDir: ".ext/modules/extapi",
		Modules:   modules,
	})

	return mergedModules, nil
}

func (e *ExtendedGcpProvider) Api(stack cdktf.TerraformStack, name string, config *deploymentspb.Api) error {

	e.Apis[name] = extapi.NewExtapi(stack, jsii.String(name), &extapi.ExtapiConfig{})
	// Replace the implementation of the `Api` method with your own logic

	// May also need to extend the gateway plugin implementation to support the new API routing
	// OR have the API routing conform to the base GCP gateway plugin

	return nil
}

func NewExtendedAwsProvider() *ExtendedGcpProvider {
	baseProvider := deploytf.NewNitricGcpProvider()

	return &ExtendedGcpProvider{
		Apis:                       make(map[string]*extapi.Extapi),
		NitricGcpTerraformProvider: *baseProvider,
	}
}
