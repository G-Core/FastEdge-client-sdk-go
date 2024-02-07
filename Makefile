generate: sdk.gen.go

sdk.gen.go: api.yml
	oapi-codegen -config oapi-gen.yml api.yml
