binaries: deploybin

# build runtime binary directly into the deploy director so it can be embedded directly into the deployment engine binary
runtimebin:
	@echo Building GCP Runtime Server
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/runtime-gcp -ldflags="-s -w -extldflags=-static" ./cmd/runtime

predeploybin: runtimebin
	@cp bin/runtime-gcp cmd/deploy/runtime

# There appears to be an old namespace conflict with the protobuf definitions
deploybin: predeploybin
	@echo Building GCP Deployment Server
	@CGO_ENABLED=0 go build -o bin/deploy-gcp -ldflags="-s -w -extldflags=-static" -ldflags="-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=ignore" ./cmd/deploy

install: deploybin
	@echo installing deployment server to ${HOME}/.nitric/providers/nitric/gcptfext-0.0.1
	@mkdir -p ${HOME}/.nitric/providers/nitric/
	@if [ "$(OS)" == "Windows_NT" ]; then \
		rm -f "${HOME}/.nitric/providers/nitric/gcptfext-0.0.1"; \
		cp bin/deploy-gcp "${HOME}/.nitric/providers/nitric/gcptfext-0.0.1.exe"; \
	else \
		rm -f "${HOME}/.nitric/providers/nitric/gcptfext-0.0.1"; \
		cp bin/deploy-gcp "${HOME}/.nitric/providers/nitric/gcptfext-0.0.1"; \
	fi

generate-terraform:
	@cd deploy && npx -y cdktf-cli@0.20.10 get
