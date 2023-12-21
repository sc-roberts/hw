package cmd

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spf13/cobra"
)

func TestRunscrape(t *testing.T) {

	testserver := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

	cmd := &cobra.Command{}
	args := []string{"GET|" + testserver.URL}

	// Call the Runscrape function and store the returned status codes
	outputBuffer, err := Runscrape(cmd, args)

	if err != nil {
		fmt.Println("Error from Runscrape CMD:", err)
	}

	// Print or assert the status codes as needed
	fmt.Println("testing GOGOGO", outputBuffer)
}
