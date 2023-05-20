package odata

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"runtime"

	"github.com/phanikumarps/sap/odata/httpclient"
	"github.com/rs/zerolog/log"
)

// CheckErrors checks if there was a err or the response returned with a non-viable status and returns an error if so.
func CheckErrors(response *http.Response, err error) error {
	if err != nil {
		pc, _, _, _ := runtime.Caller(2)
		details := runtime.FuncForPC(pc)
		if response != nil {
			log.Error().
				Err(err).
				Int("status code", response.StatusCode).
				Str("url", response.Request.URL.Path).
				Str("handler", details.Name()).
				Msg("received non-200 respons")
			return err
		}
		log.Error().
			Err(err).
			Str("handler", details.Name()).
			Str("error message", err.Error()).
			Msg("request failed")
		return err
	}
	return nil
}

// HandleResponse checks errors and unmarshals the response data, returning the deserializationTarget if all is well.
func HandleResponse[T any](response *http.Response, err error, deserializationTarget *T) (*T, error) {
	if checkErr := CheckErrors(response, err); checkErr != nil {
		return nil, checkErr
	}
	defer func() {
		_ = response.Body.Close()
	}()

	data, parseErr := io.ReadAll(response.Body)
	if parseErr != nil {
		log.Error().Err(parseErr).Msg("unable to read response.Body")
		return nil, parseErr
	}

	if unmarshalErr := json.Unmarshal(data, deserializationTarget); unmarshalErr != nil {
		log.Error().Err(unmarshalErr).Msg("unable to unmarshal response data")
		return nil, unmarshalErr
	}

	return deserializationTarget, nil
}

func SetDefaults() *httpclient.Options {
	o := httpclient.Options{
		HostUrl:   os.Getenv("SAP_HOST"),
		RootPath:  os.Getenv("UMC_SERVICE"),
		SapClient: os.Getenv("SAP_CLIENT"),
		AuthToken: os.Getenv("SAP_AUTH"),
	}
	return &o
}
