// Package sdk is FastEdge API client SDK. It provides a Go client to be used in CLI tool and Terraform provider.
// The SDK is generated from the FastEdge API OpenAPI specification (see Makefile for generation command).
package sdk

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
)

var ErrNotImplemented = errors.New("not implemented")
var _ ClientInterface = ClientSDKmock{} // to make sure ClientSDKmock implements ClientInterface

/*
ClientSDKmock is a mock implementation of the ClientInterface. Keep it in sync with ClientInterface!
To use this mock, you can create a new instance of it and override only the methods you need:

	type mock struct {
	    ClientSDKmock
	}

	func (mock) GetBinary(ctx context.Context, id int64, reqEditors ...RequestEditorFn) (*http.Response, error) {
	    return NewJsonHttpResponse(http.StatusOK, `{"id": 1, "checksum":"1234567890"}`), nil
	}
*/
type ClientSDKmock struct{}

func (ClientSDKmock) ListApps(ctx context.Context, params *ListAppsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	return nil, ErrNotImplemented
}
func (ClientSDKmock) AddAppWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	return nil, ErrNotImplemented
}
func (ClientSDKmock) AddApp(ctx context.Context, body AddAppJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	return nil, ErrNotImplemented
}
func (ClientSDKmock) GetAppIdByName(ctx context.Context, name string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	return nil, ErrNotImplemented
}
func (ClientSDKmock) DelApp(ctx context.Context, id int64, reqEditors ...RequestEditorFn) (*http.Response, error) {
	return nil, ErrNotImplemented
}
func (ClientSDKmock) GetApp(ctx context.Context, id int64, reqEditors ...RequestEditorFn) (*http.Response, error) {
	return nil, ErrNotImplemented
}
func (ClientSDKmock) PatchAppWithBody(ctx context.Context, id int64, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	return nil, ErrNotImplemented
}
func (ClientSDKmock) PatchApp(ctx context.Context, id int64, body PatchAppJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	return nil, ErrNotImplemented
}
func (ClientSDKmock) UpdateAppWithBody(ctx context.Context, id int64, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	return nil, ErrNotImplemented
}
func (ClientSDKmock) UpdateApp(ctx context.Context, id int64, body UpdateAppJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	return nil, ErrNotImplemented
}
func (ClientSDKmock) ListLogs(ctx context.Context, id int64, params *ListLogsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	return nil, ErrNotImplemented
}
func (ClientSDKmock) ListBinaries(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	return nil, ErrNotImplemented
}
func (ClientSDKmock) StoreBinaryWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	return nil, ErrNotImplemented
}
func (ClientSDKmock) DelBinary(ctx context.Context, id int64, reqEditors ...RequestEditorFn) (*http.Response, error) {
	return nil, ErrNotImplemented
}
func (ClientSDKmock) GetBinary(ctx context.Context, id int64, reqEditors ...RequestEditorFn) (*http.Response, error) {
	return nil, ErrNotImplemented
}
func (ClientSDKmock) GetClientMe(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	return nil, ErrNotImplemented
}
func (ClientSDKmock) ListSecrets(ctx context.Context, params *ListSecretsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	return nil, ErrNotImplemented
}
func (ClientSDKmock) AddSecretWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	return nil, ErrNotImplemented
}
func (ClientSDKmock) AddSecret(ctx context.Context, body AddSecretJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	return nil, ErrNotImplemented
}
func (ClientSDKmock) DeleteSecret(ctx context.Context, id int64, params *DeleteSecretParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	return nil, ErrNotImplemented
}
func (ClientSDKmock) GetSecret(ctx context.Context, id int64, reqEditors ...RequestEditorFn) (*http.Response, error) {
	return nil, ErrNotImplemented
}
func (ClientSDKmock) UpdateSecretWithBody(ctx context.Context, id int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	return nil, ErrNotImplemented
}
func (ClientSDKmock) UpdateSecret(ctx context.Context, id int, body UpdateSecretJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	return nil, ErrNotImplemented
}
func (ClientSDKmock) AddSecretSlotWithBody(ctx context.Context, secretId int64, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	return nil, ErrNotImplemented
}
func (ClientSDKmock) AddSecretSlot(ctx context.Context, secretId int64, body AddSecretSlotJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	return nil, ErrNotImplemented
}
func (ClientSDKmock) DeleteSecretSlot(ctx context.Context, secretId int64, slot int64, reqEditors ...RequestEditorFn) (*http.Response, error) {
	return nil, ErrNotImplemented
}
func (ClientSDKmock) UpdateSecretSlotWithBody(ctx context.Context, secretId int64, slot int64, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	return nil, ErrNotImplemented
}
func (ClientSDKmock) UpdateSecretSlot(ctx context.Context, secretId int64, slot int64, body UpdateSecretSlotJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	return nil, ErrNotImplemented
}
func (ClientSDKmock) StatsDuration(ctx context.Context, params *StatsDurationParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	return nil, ErrNotImplemented
}
func (ClientSDKmock) StatsCalls(ctx context.Context, params *StatsCallsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	return nil, ErrNotImplemented
}
func (ClientSDKmock) ListTemplates(ctx context.Context, params *ListTemplatesParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	return nil, ErrNotImplemented
}
func (ClientSDKmock) AddTemplateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	return nil, ErrNotImplemented
}
func (ClientSDKmock) AddTemplate(ctx context.Context, body AddTemplateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	return nil, ErrNotImplemented
}
func (ClientSDKmock) DelTemplate(ctx context.Context, id int64, reqEditors ...RequestEditorFn) (*http.Response, error) {
	return nil, ErrNotImplemented
}
func (ClientSDKmock) GetTemplate(ctx context.Context, id int64, reqEditors ...RequestEditorFn) (*http.Response, error) {
	return nil, ErrNotImplemented
}
func (ClientSDKmock) UpdateTemplateWithBody(ctx context.Context, id int64, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	return nil, ErrNotImplemented
}
func (ClientSDKmock) UpdateTemplate(ctx context.Context, id int64, body UpdateTemplateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	return nil, ErrNotImplemented
}

// BufCloser implements io.ReadCloser
type BufCloser struct {
	bytes.Buffer
}

func NewBufCloser(content string) *BufCloser {
	return &BufCloser{Buffer: *bytes.NewBufferString(content)}
}

func (b BufCloser) Close() error { return nil }

// NewJsonHttpResponse creates a new http.Response with the given status code and content. The content is expected to be a JSON string.
func NewJsonHttpResponse(status int, content string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Body:       NewBufCloser(content),
		Header:     http.Header{"Content-Type": []string{"application/json"}},
	}
}
