package deploy

import (
	"embed"

	"github.com/nitrictech/nitric/cloud/common/deploy/provider"
	"github.com/nitrictech/nitric/cloud/gcp/deploytf"
	"github.com/nitrictech/terraform-gcp-extension/deploy/generated/extapi"
)

type ExtendedGcpProvider struct {
	deploytf.NitricGcpTerraformProvider

	// Record deployed apis
	Apis map[string]extapi.Extapi
}

// HCL modules located in this embed will be added to the base provider
// These modules are provided as replacements for a module in the provider but do not need to match the original module implementation
// Additional code will need to be written in order to create these modules e.g. overriding the original module method implementation e.g. Api or Bucket
//
//go:embed .ext/modules/**/*
var modules embed.FS

// HCL modules located in this embed will directly override the modules in the base provider
// therefore these modules MUST implement the same variables and output as the equivalent module in the base provider
//
//go:embed .nitric/modules/**/*
var nitricOverrides embed.FS

// TODO: Need to update this extension so we can merge the modules directory with the base provider
func (a *ExtendedGcpProvider) CdkTfModules() ([]provider.ModuleDirectory, error) {
	// Get the original modules from the provider we're extending
	origModules, err := a.NitricGcpTerraformProvider.CdkTfModules()
	if err != nil {
		return nil, err
	}

	// Merge the original modules with our new modules
	mergedModules := append(origModules, provider.ModuleDirectory{
		ParentDir: ".ext/modules",
		Modules:   modules,
	}, provider.ModuleDirectory{
		ParentDir: ".nitric/modules",
		Modules:   nitricOverrides,
	})

	return mergedModules, nil
}

func NewExtendedGcpProvider() *ExtendedGcpProvider {
	baseProvider := deploytf.NewNitricGcpProvider()

	return &ExtendedGcpProvider{
		Apis:                       make(map[string]extapi.Extapi),
		NitricGcpTerraformProvider: *baseProvider,
	}
}
