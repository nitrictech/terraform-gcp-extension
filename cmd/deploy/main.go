package main

import (
	_ "embed"

	"github.com/nitrictech/nitric/cloud/common/deploy/provider"
	"github.com/nitrictech/terraform-gcp-extension/deploy"
)

//go:embed runtime
var runtimeBin []byte

var runtimeProvider = func() []byte {
	return runtimeBin
}

// Start the deployment server
func main() {
	stack := deploy.NewExtendedGcpProvider()

	providerServer := provider.NewTerraformProviderServer(stack, runtimeProvider)

	providerServer.Start()
}
