package umc

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/phanikumarps/sap/odata"
	"github.com/phanikumarps/sap/odata/httpclient"
	"github.com/phanikumarps/sap/odata/isu/umc/account"
)

type Service struct {
	httpclient.Client
}

func NewService(host, auth string) *Service {

	ser := Service{
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

func (service *Service) GetAccount(ctx context.Context, id string) (*account.AcctOutput, error) {
	res, err := service.Get(ctx, fmt.Sprintf(accountURL, id))
	if err != nil {
		return nil, err
	}

	resp, err := odata.HandleResponse(res, err, &account.AcctResponse{})
	if err != nil {
		return nil, err
	}

	acct := account.ConvRespToOutput(resp)
	return acct, nil
}

const (
	accountURL = "/Accounts('%s)"
)
