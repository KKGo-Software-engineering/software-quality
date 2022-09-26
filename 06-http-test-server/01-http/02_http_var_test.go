package http

import (
	"bytes"
	"io"
	"net/http"
	"reflect"
	"testing"
)

func TestBadVar(t *testing.T) {
	get = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(`{"id": 1, "name": "AnuchitO", "info": "gopher"}`)),
		}, nil
	}

	resp, err := MakeHTTPCall("http://localhost:8080")
	if err != nil {
		t.Fatal(err)
	}

	want := &Response{
		ID:   1,
		Name: "AnuchitO",
		Info: "gopher",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("expected (%v), got (%v)", want, resp)
	}
}
