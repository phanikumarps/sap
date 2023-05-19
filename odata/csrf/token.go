package csrf

import (
	"context"
	"net/http"
	"os"

	"github.com/phanikumarps/sap/odata/config"
	"github.com/phanikumarps/sap/odata/httpclient"
)

func NewToken(ctx context.Context, host, clnt, auth string) (*string, error) {
	s := NewService(
		host,
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

func NewService(host, sapClient, authToken string) *Service {
	s := new(Service)
	s.Client = *httpclient.New(
		httpclient.WithHost(host),
		httpclient.WithSapClient(sapClient),
		httpclient.WithAuthToken(authToken),
	)
	return s
}

func (s *Service) Get() (*string, error) {
	ctx := context.TODO()
	resource := httpclient.RequestOptions{Path: ""}
	r := config.DefaultRootPath(os.Getenv("UMC_SERVICE"))
	resp, err := s.Call(ctx, http.MethodHead, string(*r), resource.Path, nil, "")
	if err != nil {
		return nil, err
	}
	t := resp.Header.Get("X-Csrf-Token")
	return &t, nil
}

type Srv struct {
	client *httpclient.Clnt
}

func (s *Srv) GetToken(ctx context.Context) (*string, error) {

	resp, err := s.client.Get(ctx, "")
	if err != nil {
		return nil, err
	}
	t := resp.Header.Get("X-Csrf-Token")

	return &t, nil
}

func NewTk(host, auth string) (*string, error) {

	c := httpclient.NewClnt(
		http.DefaultClient,
		httpclient.Options{
			HostUrl:   host,
			RootPath:  os.Getenv("UMC_SERVICE"),
			SapClient: os.Getenv("SAP_CLIENT"),
			AuthToken: auth,
			Verbose:   false,
		})
	s := new(Srv)
	s.client = c
	ctx := context.TODO()
	tk, err := s.GetToken(ctx)
	if err != nil {
		return nil, err
	}
	return tk, nil

}
