package httpclient

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
)

type TokenFetcher interface {
	Head(ctx context.Context, path string) (csrf string, err error)
}
type Getter interface {
	Get(ctx context.Context, path string) (resp *http.Response, err error)
}

type Inserter interface {
	Post(ctx context.Context, path string, body io.Reader) (resp *http.Response, err error)
}

type Updater interface {
	Put(ctx context.Context, path string, body io.Reader) (resp *http.Response, err error)
	Patch(ctx context.Context, path string, body io.Reader) (resp *http.Response, err error)
}
type Deleter interface {
	Delete(ctx context.Context, path string, body io.Reader) (resp *http.Response, err error)
}

type Options struct {
	HostUrl   string
	RootPath  string
	SapClient string
	AuthToken string
	Verbose   bool
}

type Clnt struct {
	httpclient *http.Client
	options    *Options
}

func NewClnt(httpClient *http.Client, options Options) *Clnt {
	return &Clnt{
		httpclient: httpClient,
		options:    &options,
	}
}

func (c *Clnt) Get(ctx context.Context, path string) (*http.Response, error) {
	u := c.options.HostUrl + c.options.RootPath + path
	if path != "" {
		// build response format
		f := "?" + "$format=" + "json"
		u = u + f
	}
	request, err := http.NewRequest(http.MethodGet, u, nil)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	request.Header.Add("sap-client", c.options.SapClient)
	request.Header.Add("Authorization", c.options.AuthToken)
	request.Header.Add("X-Csrf-Token", "fetch")

	request = request.WithContext(ctx)

	if c.options.Verbose {
		body, _ := httputil.DumpRequest(request, true)
		log.Println(fmt.Sprintf("%s", string(body)))
	}

	//calling the URL
	resp, err := c.httpclient.Do(request)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	return resp, nil

}
