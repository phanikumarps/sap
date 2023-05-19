package csrf

import (
	"context"
	"net/http"
	"os"

	"github.com/phanikumarps/sap/odata/httpclient"
)

func (s *Srv) GetToken(ctx context.Context) (*string, error) {

	resp, err := s.client.Get(ctx, "")
	if err != nil {
		return nil, err
	}
	t := resp.Header.Get("X-Csrf-Token")

	return &t, nil
}

type Srv struct {
	client *httpclient.Clnt
}

func NewSrv(host, auth string) *Srv {
	s := new(Srv)
	s.client = httpclient.NewClnt(
		http.DefaultClient,
		httpclient.Options{
			HostUrl:   host,
			RootPath:  os.Getenv("UMC_SERVICE"),
			SapClient: os.Getenv("SAP_CLIENT"),
			AuthToken: auth,
			Verbose:   false,
		})
	return s
}

func NewTk(host, auth string) (*string, error) {

	s := NewSrv(host, auth)
	ctx := context.TODO()
	tk, err := s.GetToken(ctx)
	if err != nil {
		return nil, err
	}
	return tk, nil

}
