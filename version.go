package sdk

import (
	"context"
	"fmt"
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

type versionCheckingHTTPClient struct {
	client   http.Client
	toolName string
}

func (c *versionCheckingHTTPClient) Do(req *http.Request) (*http.Response, error) {
	ret, err := c.client.Do(req)
	if ret.StatusCode == http.StatusPreconditionFailed {
		return nil, fmt.Errorf("API version is not supported by the server, please update your %s", c.toolName)
	}
	return ret, err
}

func NewClientWithVersionCheck(apiUrl, userAgent, toolName string, opts ...ClientOption) (*ClientWithResponses, error) {
	opts = append(opts,
		WithHTTPClient(
			&versionCheckingHTTPClient{
				client:   http.Client{},
				toolName: toolName,
			}),
		WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			ver := GetVersion()
			if ver != "" {
				req.Header.Set(versionHeaderName, ver)
			}
			return nil
		}),
		WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			if userAgent != "" {
				req.Header.Set("User-Agent", userAgent)
			}
			return nil
		}),
	)
	return NewClientWithResponses(apiUrl, opts...)
}
