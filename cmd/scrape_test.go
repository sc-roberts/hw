package cmd

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunscrape(t *testing.T) {

	testserver := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "abcd1234")
	}))

	method := "GET"
	teststring := "abcd1234"
	testc := []string{method + "|" + testserver.URL}

	// Call the Runscrape function and store the returned output
	outputBuffer, err := Runscrape(testc)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("test function output buffer", outputBuffer)

	assert.Contains(
		t,
		outputBuffer.String(),
		teststring,
	)
}
