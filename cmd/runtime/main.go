package main

import (
	"github.com/nitrictech/nitric/cloud/gcp/runtime"
	"github.com/nitrictech/nitric/cloud/gcp/runtime/resource"
	"github.com/nitrictech/nitric/core/pkg/logger"
	"github.com/nitrictech/nitric/core/pkg/server"
)

func main() {
	resolver, err := resource.New()
	if err != nil {
		logger.Fatalf("could not create aws resource resolver: %v", err)
		return
	}

	m, err := runtime.NewGcpRuntimeServer(resolver)
	if err != nil {
		logger.Fatalf("there was an error initializing the AWS runtime server: %v", err)
	}

	server.Run(m)
}
