package httpclient

import (
	"context"
	"io"
	"log"
	"net/http"
)

type Client struct {
	host
	proxy     string
	authToken string
	csrfToken string
}

func New(options ...optionalFunc) *Client {
	clnt := &Client{}
	for _, opt := range options {
		opt(clnt)
	}
	return clnt
}

func WithHost(host string) optionalFunc {
	return func(c *Client) {
		c.host.hostUrl = host
	}
}

func WithSapClient(sapClnt string) optionalFunc {
	return func(c *Client) {
		c.host.sapClient = sapClnt
	}
}

func WithProxy(proxy string) optionalFunc {
	return func(c *Client) {
		c.proxy = proxy
	}
}

func WithAuthToken(authToken string) optionalFunc {
	return func(c *Client) {
		c.authToken = authToken
	}
}

func WithCsrfToken(csrfToken string) optionalFunc {
	return func(c *Client) {
		c.csrfToken = csrfToken
	}
}

func Default(authToken string) *Client {
	return New(
		WithHost(defaultHost.hostUrl),
		WithSapClient(defaultHost.sapClient),
		WithAuthToken(authToken),
	)
}

func (c *Client) Call(ctx context.Context, httpMethod string, rootResource string, resource string, body io.Reader, format string) (*http.Response, error) {
	u := c.host.hostUrl + rootResource + resource
	// build response format
	if format == "json" {
		f := "?" + "$format=" + "json"
		u = u + f
	}

	request, err := http.NewRequest(httpMethod, u, body)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	request.Header.Add("sap-client", c.sapClient)

	request.Header.Add("Authorization", c.authToken)

	if httpMethod == http.MethodGet || httpMethod == http.MethodHead {
		request.Header.Add("X-Csrf-Token", "fetch")
	} else {
		request.Header.Add("X-Csrf-Token", c.csrfToken)
	}

	request = request.WithContext(ctx)
	//calling the URL
	client := http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	return resp, nil

}

type RequestOptions struct {
	Path string
}

type optionalFunc func(*Client)

type host struct {
	hostUrl   string
	sapClient string
}

var defaultHost host

type HttpClient struct {
	host
	rootResource string
	proxy        string
	authToken    string
	csrfToken    string
}

type Doer interface {
	Do(ctx context.Context, httpMethod string, resourceUrl string, body io.Reader) (*http.Response, error)
}

func (c *HttpClient) Do(ctx context.Context, httpMethod string, resourceUrl string, body io.Reader) (*http.Response, error) {
	u := c.host.hostUrl + c.rootResource + resourceUrl
	if resourceUrl != "" {
		// build response format
		f := "?" + "$format=" + "json"
		u = u + f
	}
	request, err := http.NewRequest(httpMethod, u, body)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	request.Header.Add("sap-client", c.sapClient)

	request.Header.Add("Authorization", c.authToken)

	if httpMethod == http.MethodGet || httpMethod == http.MethodHead {
		request.Header.Add("X-Csrf-Token", "fetch")
	} else {
		request.Header.Add("X-Csrf-Token", c.csrfToken)
	}

	//calling the URL
	client := http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	return resp, nil

}
