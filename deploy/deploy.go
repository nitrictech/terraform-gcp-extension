package deploy

import (
	"embed"
	"io/fs"

	"github.com/nitrictech/nitric/cloud/common/deploy/provider"
	"github.com/nitrictech/nitric/cloud/gcp/deploytf"
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

//go:embed .nitric/modules/**/*
var nitricOverrides embed.FS

// TODO: Need to update this extension so we can merge the modules directory with the base provider
func (a *ExtendedGcpProvider) CdkTfModules() ([]provider.ModuleDirectory, error) {
	origModules, err := a.NitricGcpTerraformProvider.CdkTfModules()
	if err != nil {
		return nil, err
	}

	// Merge the original modules with the embedded modules
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
