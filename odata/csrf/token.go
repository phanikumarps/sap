package csrf

import (
	"context"
	"log"
	"net/http"

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
	r := config.DefaultRootPath("ZERP_ISU_UMC")
	resp, err := s.Call(ctx, http.MethodHead, string(*r)+"/", resource.Path, nil, "")
	if err != nil {
		return nil, err
	}
	t := resp.Header.Get("X-Csrf-Token")
	return &t, nil
}

func GetTest(Url string) (*string, error) {
	request, err := http.NewRequest(http.MethodGet, Url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer func() {
		_ = request.Body.Close()
	}()

	//calling the URL
	transport := &http.Transport{
		Proxy: http.ProxyURL(nil),
	}

	//adding the Transport object to the http Client
	client := &http.Client{
		Transport: transport,
	}

	resp, err := client.Do(request)
	if err != nil {
		//log.Println(err)
		return nil, err
	}

	return &resp.Status, nil
}
