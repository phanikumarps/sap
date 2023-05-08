package umc

import (
	"context"
	"log"
	"net/http"
)

func NewCSRFToken(ctx context.Context, host string, auth string) (*string, error) {

	Url := host + RootResource + "$metadata"

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
