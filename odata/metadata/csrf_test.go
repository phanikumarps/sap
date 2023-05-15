package metadata

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewCSRFToken(t *testing.T) {
	want := "token"
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, want)
	}))
	defer svr.Close()

	ctx := context.TODO()
	got, err := NewCSRFToken(ctx, "localhost", "443", "100", "auth")
	if err != nil {
		t.Errorf(err.Error())
	}

	if *got != want {
		t.Errorf("expected %s, got %s", want, *got)
	}
}
