package account

import (
	"context"
	"fmt"
	"net/http"

	"github.com/phanikumarps/sap/odata/httpclient"
	"github.com/phanikumarps/sap/odata/isu/umc"
)

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
func (s *Service) GetAccount(ctx context.Context, id string) (*http.Response, error) {
	resource := httpclient.RequestOptions{Path: fmt.Sprintf(GetAccountUrl, id)}
	r := umc.DefaultUmcRootPath()
	resp, err := s.Call(ctx, http.MethodGet, string(*r)+"/", resource.Path, nil, "json")
	if err != nil {
		return nil, err
	}
	return resp, nil
}

const (
	GetAccountUrl = "/Accounts('%s')"
)
