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

func TestMakeHttp(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	want := &Response{
		ID:   1,
		Name: "AnuchitO",
		Info: "gopher",
	}

	t.Run("Happy server response", func(t *testing.T) {
		resp, err := MakeHTTPCall(server.URL)

		if !reflect.DeepEqual(resp, want) {
			t.Errorf("expected (%v), got (%v)", want, resp)
		}

		if !errors.Is(err, nil) {
			t.Errorf("expected (%v), got (%v)", nil, err)
		}
	})
}
