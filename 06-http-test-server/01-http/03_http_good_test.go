package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"id": 1, "name": "AnuchitO", "info": "gopher"}`))
}

func setup(t *testing.T) (*httptest.Server, func()) {
	t.Helper()

	server := httptest.NewServer(http.HandlerFunc(handler))

	teardown := func() {
		server.Close()
	}

	return server, teardown
}

func TestMakeHTTP(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id": 1, "name": "AnuchitO", "info": "gopher"}`))
	}))
	want := &Response{
		ID:   1,
		Name: "AnuchitO",
		Info: "gopher",
	}

	t.Run("Happy server response", func(t *testing.T) {
		defer server.Close()

		resp, err := MakeHTTPCall(server.URL)

		if !reflect.DeepEqual(resp, want) {
			t.Errorf("expected (%v), got (%v)", want, resp)
		}

		if !errors.Is(err, nil) {
			t.Errorf("expected (%v), got (%v)", nil, err)
		}
	})
}
