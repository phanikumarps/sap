package account

import (
	"context"
	"fmt"
	"net/http"

	"github.com/phanikumarps/sap/odata/config"
	"github.com/phanikumarps/sap/odata/httpclient"
	"github.com/phanikumarps/sap/odata/isu/umc"
)

func (s *acctService) GetAccount(ctx context.Context, id string) (*http.Response, error) {
	resource := httpclient.RequestOptions{Path: fmt.Sprintf(getAccountUrl, id)}
	r := umc.DefaultUmcRootPath()
	resp, err := s.Call(ctx, http.MethodGet, string(*r)+"/", resource.Path, nil, "json")
	if err != nil {
		return nil, err
	}
	return resp, nil
}

const (
	getAccountUrl = "/Accounts('%s')"
)

type acctService config.Service
