package cmd

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestRunscrape(t *testing.T) {

	testserver := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "abcd1234")
	}))

	cmd := &cobra.Command{}
	method := "GET"
	teststring := "abcd1234"
	args := []string{method + "|" + testserver.URL}
	cmd.Flags().StringSlice("targets", args, "test")

	// Call the Runscrape function and store the returned output
	outputBuffer, err := Runscrape(cmd, args)

	if err != nil {
		fmt.Println("Error from Runscrape CMD:", err)
	}

	fmt.Println("test function output buffer", outputBuffer)

	assert.Contains(
		t,
		outputBuffer.String(),
		teststring,
	)
}
