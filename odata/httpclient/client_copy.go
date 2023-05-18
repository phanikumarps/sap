package httpclient

/*
import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
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
	Get(ctx context.Context, path string, resp any) error
}

type Inserter interface {
	Post(ctx context.Context, path string, body io.Reader, resp any) error
}

type Updater interface {
	Put(ctx context.Context, path string, body io.Reader, resp any) error
	Patch(ctx context.Context, path string, body io.Reader, resp any) error
}
type Deleter interface {
	Delete(ctx context.Context, path string, body io.Reader, resp any) error
}

type Options struct {
	hostUrl  string
	rootPath string
	verbose  bool
}

type Clnt struct {
	httpclient *http.Client
	options    *Options
	csrfToken  string
}

func New(clnt *http.Client, options *Options) *Clnt {
	return &Clnt{
		httpclient: clnt,
		options:    options,
	}
}


func (c *Clnt) Get(ctx context.Context, path string, receiver any) error {

	u := c.options.hostUrl + c.options.rootPath + path

	// build response format
	f := "?" + "$format=" + "json"
	u = u + f

	request, err := c.newRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return fmt.Errorf("failed to create GET request: %w", err)
	}

	request.Header.Add("X-Csrf-Token", "fetch")

	if err := c.doRequest(request, receiver); err != nil {
		return err
	}

	return nil
}

func (c *Clnt) newRequest(ctx context.Context, method, path string, payload any) (*http.Request, error) {
	var reqBody io.Reader
	if payload != nil {
		bodyBytes, err := json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewReader(bodyBytes)
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.options.hostUrl, path), reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	if c.options.verbose {
		body, _ := httputil.DumpRequest(req, true)
		log.Println(fmt.Sprintf("%s", string(body)))
	}

	req = req.WithContext(ctx)
	return req, nil
}

func (c *Clnt) doRequest(r *http.Request, receiver any) error {
	resp, err := c.do(r)
	if err != nil {
		return err
	}

	if resp == nil {
		return nil
	}
	defer resp.Body.Close()

	if receiver == nil {
		return nil
	}

	var buf bytes.Buffer
	dec := json.NewDecoder(io.TeeReader(resp.Body, &buf))
	if err := dec.Decode(receiver); err != nil {
		return fmt.Errorf("could not parse response body: %w [%s:%s] %s", err, r.Method, r.URL.String(), buf.String())
	}

	return nil
}

func (c *Clnt) do(r *http.Request) (*http.Response, error) {

	var (
		ErrUserAccessDenied = errors.New("you do not have access to the requested resource")
		ErrNotFound         = errors.New("the requested resource not found")
		ErrTooManyRequests  = errors.New("you have exceeded throttle")
	)

	resp, err := c.httpclient.Do(r)
	if err != nil {
		return nil, fmt.Errorf("failed to make request [%s:%s]: %w", r.Method, r.URL.String(), err)
	}

	if c.options.verbose {
		body, _ := httputil.DumpResponse(resp, true)
		log.Println(fmt.Sprintf("%s", string(body)))
	}

	switch resp.StatusCode {
	case http.StatusOK,
		http.StatusCreated,
		http.StatusNoContent:
		return resp, nil
	}

	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusNotFound:
		return nil, ErrNotFound
	case http.StatusUnauthorized,
		http.StatusForbidden:
		return nil, ErrUserAccessDenied
	case http.StatusTooManyRequests:
		return nil, ErrTooManyRequests
	}

	return nil, fmt.Errorf("failed to do request, %d status code received", resp.StatusCode)
}
*/
