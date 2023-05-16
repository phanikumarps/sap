package csrf_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/phanikumarps/sap/odata/csrf"
)

func TestNewToken(t *testing.T) {
	exp := "test"
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Test request parameters
		// equals(t, r.URL.String(), "/some/path")
		w.Header().Add("X-Csrf-Token", "test")
		w.Header().Add("Authorization", "auth")
		w.Header().Add("sap-client", "100")
	}))
	defer svr.Close()

	ctx := context.TODO()

	got, err := csrf.NewToken(ctx, "https://localhost:8000", "100", "auth")
	if err != nil {
		t.Error("error here")
		t.Errorf(err.Error())
	}

	if *got != exp {
		t.Errorf("expected %s, got %s", exp, *got)
	}
}
