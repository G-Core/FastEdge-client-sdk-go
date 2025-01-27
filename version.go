package sdk

import (
	"context"
	"net/http"
	"runtime/debug"
	"strings"
)

const (
	versionHeaderName = "Fastedge-Sdk-Version"
	SDKpackage        = "github.com/G-Core/FastEdge-client-sdk-go"
)

var version *string

func GetVersion() string {
	if version != nil {
		return *version
	}
	bi, ok := debug.ReadBuildInfo()
	if ok {
		for _, dep := range bi.Deps {
			if dep.Path == SDKpackage {
				ver := strings.SplitN(dep.Version, "-", 2) // drop revision info
				version = &ver[0]
				return ver[0]
			}
		}
	}
	empty := ""
	version = &empty
	return empty
}

func AddVersionHeader(ctx context.Context, req *http.Request) error {
	ver := GetVersion()
	if ver != "" {
		req.Header.Set(versionHeaderName, ver)
	}
	return nil
}
