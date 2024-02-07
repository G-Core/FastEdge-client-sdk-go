generate: pkg/sdk/sdk.gen.go

pkg/sdk/sdk.gen.go: pkg/sdk/api.yml
	oapi-codegen -config oapi-gen.yml pkg/sdk/api.yml
