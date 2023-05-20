package csrf

import (
	"context"
	"net/http"
	"os"

	"github.com/phanikumarps/sap/odata/httpclient"
)

func NewTk(host, auth string) (*string, error) {

	s := newService(host, auth)
	ctx := context.TODO()
	tk, err := s.getToken(ctx)
	if err != nil {
		return nil, err
	}
	return tk, nil

}

func (s *service) getToken(ctx context.Context) (*string, error) {

	resp, err := s.Get(ctx, "")
	if err != nil {
		return nil, err
	}
	t := resp.Header.Get("X-Csrf-Token")

	return &t, nil
}

type service struct {
	httpclient.Client
}

func newService(host, auth string) *service {

	ser := service{
		*httpclient.NewClnt(
			http.DefaultClient,
			httpclient.Options{
				HostUrl:   host,
				RootPath:  os.Getenv("UMC_SERVICE"),
				SapClient: os.Getenv("SAP_CLIENT"),
				AuthToken: auth,
				Verbose:   false,
			}),
	}
	return &ser
}
