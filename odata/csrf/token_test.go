package csrf_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/phanikumarps/sap/odata/csrf"
)

func TestNewTk(t *testing.T) {

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Csrf-Token", "token")
		w.Header().Add("Authorization", "auth")
		w.Header().Add("sap-client", "100")
	}))
	defer svr.Close()

	exp := "token"
	got, err := csrf.NewTk(svr.URL, "auth")
	if err != nil {
		t.Error("error here")
		t.Errorf(err.Error())
	}

	if *got != exp {
		t.Errorf("expected %s, got %s", exp, *got)
	}
}
