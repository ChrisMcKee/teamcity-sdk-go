package teamcity

import (
	"bytes"
	"io"
	"net/http"
)

func NewTestClient(replyResp *http.Response, err error) *Client {
	client := &Client{
		username: "username",
		password: "password",
		host:     "host.example.com",
		version:  "latest",
		retries:  8,
	}
	httpClient := &http.Client{}
	httpClient.Transport = &MockTransport{
		resp: replyResp,
		err:  err,
	}
	client.HTTPClient = httpClient
	return client
}

type MockTransport struct {
	req  *http.Request
	resp *http.Response
	err  error
}

func (b *MockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	b.req = req
	return b.resp, b.err
}

func newResponse(body string) *http.Response {
	resp := &http.Response{Body: io.NopCloser(bytes.NewBuffer([]byte(body)))}
	return resp
}

func newCodeResponse(status string, httpCode int, body string) *http.Response {
	resp := &http.Response{
		Status:     status,
		StatusCode: httpCode,
		Body:       io.NopCloser(bytes.NewBuffer([]byte(body))),
	}

	return resp
}
