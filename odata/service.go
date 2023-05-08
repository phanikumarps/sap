package odata

import (
	"context"
	"log"
	"net/http"
)

const (
	RootResource = "/sap/opu/odata/sap/ERP_ISU_UMC/"
)

type Host struct {
	Address string
	Port    string
}

func NewHost(h Host) *Host {

	return &Host{
		Address: h.Address,
		Port:    h.Port,
	}
}

type Resource struct {
	Name      string
	Host      Host
	AuthToken string
	CsrfToken string
}

func NewResource(r Resource) *Resource {
	return &Resource{
		Name:      r.Name,
		Host:      r.Host,
		AuthToken: r.AuthToken,
		CsrfToken: r.CsrfToken,
	}
}

func NewCSRFToken(ctx context.Context, hostUrl string, auth string) (*string, error) {

	Url := hostUrl + RootResource + "$metadata"

	request, err := http.NewRequest(http.MethodHead, Url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer request.Body.Close()

	request.Header.Add("Authorization", auth)
	request.Header.Add("X-Csrf-Token", "fetch")

	//calling the URL
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	token := resp.Header.Get("X-Csrf-Token")
	return &token, nil
}
