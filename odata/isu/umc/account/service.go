package account

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/phanikumarps/sap/odata"
)

type Account struct {
	AccountID string
	odata.Resource
}

type AcctGetter interface {
	AcctGet(ctx context.Context, AccountID string) (*acctResponse, error)
}

func AcctGet(ctx context.Context, AccountID string) (*acctResponse, error) {
	var acc Account
	acc.AccountID = AccountID

	resp, err := acc.Get(ctx, acc.AccountID)
	if err != nil {
		return nil, fmt.Errorf("error calling resource %s", acc.AccountID)
	}

	// unmarshall
	var accResp acctResponse
	err = json.NewDecoder(*resp).Decode(&accResp)
	if err != nil {
		return nil, err
	}

	return &accResp, nil

}
func (a *Account) Get(ctx context.Context, AccountID any) (*io.ReadCloser, error) {

	accountid := fmt.Sprintf("%v", AccountID)

	fmt.Println(accountid)

	// check authorization
	if a.Resource.AuthToken == "" {
		return nil, fmt.Errorf("authorization token can not be blank")
	}
	// build host url: host address + ":" + host port
	h := a.Resource.Host.Address + ":" + a.Resource.Host.Port
	// build response format
	f := "?" + "$format=" + "json"
	// build url: host url + root resource + resource + key + format
	u := h + odata.RootResource + a.Resource.Name + "('" + accountid + "')" + f

	// GET call
	request, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer request.Body.Close()

	request.Header.Add("Authorization", a.Resource.AuthToken)

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &resp.Body, nil
}

func (a *Account) Insert(ctx context.Context, body any) (*io.ReadCloser, error) {

	// check authorization
	if a.Resource.AuthToken == "" {
		return nil, fmt.Errorf("authorization token can not be blank")
	}

	// check CSRF token
	if a.Resource.CsrfToken == "" {
		return nil, fmt.Errorf("csrf token can not be blank")
	}

	// build host url: host address + ":" + host port
	h := a.Resource.Host.Address + ":" + a.Resource.Host.Port

	// build url: host url + root resource + resource + key + format
	u := h + odata.RootResource + a.Resource.Name

	// Post call
	var data io.Reader
	request, err := http.NewRequest(http.MethodPost, u, data)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer request.Body.Close()

	request.Header.Add("Authorization", a.Resource.AuthToken)
	request.Header.Add("X-Csrf-Token", a.Resource.CsrfToken)

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	r := io.NopCloser(strings.NewReader(resp.Status))
	return &r, nil
}
