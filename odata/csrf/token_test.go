package csrf_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/phanikumarps/sap/odata/csrf"
)

func TestGetTest(t *testing.T) {
	exp := "200 OK"
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer svr.Close()

	got, err := csrf.GetTest(svr.URL)
	if err != nil {
		t.Error("error here")
		t.Errorf(err.Error())
	}

	if *got != exp {
		t.Errorf("expected %s, got %s", exp, *got)
	}
}
