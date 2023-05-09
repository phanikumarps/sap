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
	Response acctResponse
}

type AcctGetter interface {
	AcctGet(ctx context.Context, AccountID string) (*acctResponse, error)
}

func AcctGet(ctx context.Context, acct Account) (*acctOutput, error) {

	resp, err := acct.Get(ctx, acct.AccountID)
	if err != nil {
		return nil, fmt.Errorf("error calling resource %s", acct.AccountID)
	}

	err = json.NewDecoder(*resp).Decode(&acct.Response)
	if err != nil {
		return nil, err
	}

	acctOutput := convRespToOutput(&acct.Response)
	return acctOutput, nil

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
	request.Header.Add("sap-client", a.Resource.SapClient)

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
	request.Header.Add("sap-client", a.Resource.SapClient)
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

func convRespToOutput(resp *acctResponse) *acctOutput {
	acctop := acctOutput{
		AccountTypeID:             resp.D.AccountTypeID,
		AccountID:                 resp.D.AccountID,
		AccountTitleID:            resp.D.AccountTitleID,
		FirstName:                 resp.D.FirstName,
		LastName:                  resp.D.LastName,
		MiddleName:                resp.D.MiddleName,
		SecondName:                resp.D.SecondName,
		Sex:                       resp.D.Sex,
		Name1:                     resp.D.Name1,
		Name2:                     resp.D.Name2,
		Name3:                     resp.D.Name3,
		Name4:                     resp.D.Name4,
		GroupName1:                resp.D.GroupName1,
		GroupName2:                resp.D.GroupName2,
		FullName:                  resp.D.FullName,
		CorrespondenceLanguage:    resp.D.CorrespondenceLanguage,
		CorrespondenceLanguageISO: resp.D.CorrespondenceLanguageISO,
		Language:                  resp.D.Language,
		LanguageISO:               resp.D.LanguageISO,
	}
	return &acctop
}
