package cmd

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ScrapeT(t *testing.T) {
	testserver := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

	outBuffer := bytes.Buffer{}

	method := "GET"

}
