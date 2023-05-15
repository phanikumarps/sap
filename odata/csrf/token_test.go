package csrf

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewToken(t *testing.T) {
	want := "token"
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, want)
	}))
	defer svr.Close()

	ctx := context.TODO()
	got, err := NewToken(ctx, "localhost", "443", "100", "auth")
	if err != nil {
		t.Errorf(err.Error())
	}

	if *got != want {
		t.Errorf("expected %s, got %s", want, *got)
	}
}
