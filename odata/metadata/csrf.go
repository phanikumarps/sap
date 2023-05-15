package metadata

import (
	"context"
	"net/http"

	"github.com/phanikumarps/sap/odata"
	"github.com/phanikumarps/sap/odata/httpclient"
)

func NewCSRFToken(ctx context.Context, host, port, clnt, auth string) (*string, error) {
	s := NewService(
		host,
		port,
		clnt,
		auth,
	)
	t, err := s.Get()
	if err != nil {
		return nil, err
	}
	return t, nil
}

type Service struct {
	httpclient.Client
}

func NewService(host, port, sapClient, authToken string) *Service {
	s := new(Service)
	s.Client = *httpclient.New(
		httpclient.WithHost(host),
		httpclient.WithPort(port),
		httpclient.WithSapClient(sapClient),
		httpclient.WithAuthToken(authToken),
	)
	return s
}
func (s *Service) Get() (*string, error) {
	ctx := context.TODO()
	resource := httpclient.RequestOptions{Path: Url}
	resp, err := s.Call(ctx, http.MethodHead, odata.RootResource, resource.Path, nil, "")
	if err != nil {
		return nil, err
	}
	t := resp.Header.Get("X-Csrf-Token")
	return &t, nil
}

const (
	Url = "$metadata"
)
