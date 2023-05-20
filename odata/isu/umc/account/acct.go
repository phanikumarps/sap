package account

/*
type Srv struct {
	client *httpclient.Client
}

type Account struct {
	Id string
}

func (s *Srv) GetAcct(ctx context.Context, id string) (*Account, error) {
	resp, err := s.client.Get(ctx, "")
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	var acc Account
	dec := json.NewDecoder(io.TeeReader(resp.Body, &buf))
	err = dec.Decode(&acc)
	if err != nil {
		return nil, err
	}

	return &acc, nil
}
func New(authToken string) *Srv {
	t := transport{
		authToken: authToken,
	}
	fmt.Println(t)

	return &Srv{
		client: httpclient.NewClnt(&http.Client{Transport: &t},
			httpclient.Options{
				HostUrl:   "",
				RootPath:  "",
				SapClient: "",
				AuthToken: "",
				Verbose:   os.Getenv("VERBOSE") != "",
			},
		),
	}

}

type transport struct {
	authToken string
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	r := req.Clone(req.Context())
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Accept", "application/json")
	r.Header.Add("Accept-Charset", "UTF-8")

	r.Header.Add("X-AUTH-API-KEY", t.authToken)

	return http.DefaultTransport.RoundTrip(r)
}
*/
