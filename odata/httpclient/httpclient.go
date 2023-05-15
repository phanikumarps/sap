package httpclient

import (
	"context"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Host struct {
	host      string
	port      string
	sapClient string
}

type Client struct {
	Host
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
		c.Host.host = host
	}
}

func WithPort(port string) optionalFunc {
	return func(c *Client) {
		c.Host.port = port
	}
}

func WithSapClient(sapClnt string) optionalFunc {
	return func(c *Client) {
		c.Host.sapClient = sapClnt
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
		WithHost(defaultHost.host),
		WithPort(defaultHost.port),
		WithSapClient(defaultHost.sapClient),
		WithAuthToken(authToken),
	)
}

func (c *Client) Call(ctx context.Context, httpMethod string, rootResource string, resource string, body io.Reader) (*http.Response, error) {
	h := c.Host.host + ":" + c.Host.port
	u := h + rootResource + resource

	request, err := http.NewRequest(httpMethod, u, body)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer func() {
		_ = request.Body.Close()
	}()

	request.Header.Add("sap-client", c.sapClient)

	request.Header.Add("Authorization", c.authToken)

	if httpMethod == http.MethodGet {
		request.Header.Add("X-Csrf-Token", "fetch")
	} else {
		request.Header.Add("X-Csrf-Token", c.csrfToken)
	}

	//calling the URL
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL(c.proxy)),
	}

	//adding the Transport object to the http Client
	client := &http.Client{
		Transport: transport,
	}

	resp, err := client.Do(request)
	if err != nil {
		//log.Println(err)
		return nil, err
	}

	return resp, nil

}

type RequestOptions struct {
	Path string
}

func proxyURL(Url string) *url.URL {
	u, err := url.Parse(Url)
	if err != nil {
		log.Println(err)
	}
	return u
}

var defaultHost Host

type optionalFunc func(*Client)
