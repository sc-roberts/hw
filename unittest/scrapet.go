package cmd

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

func Runscrape(cmd *cobra.Command, args []string) (*bytes.Buffer, error) {
	targets, _ := cmd.Flags().GetStringSlice("targets")
	var outputBuffer bytes.Buffer
	client := &http.Client{}

	for _, targ := range targets {
		splitter := strings.SplitN(targ, "|", 2)
		method := splitter[0]
		target := splitter[1]

		fmt.Println(target, method)
		request, err := http.NewRequest(method, target, nil)
		if err != nil {
			fmt.Println("Error:", err)
			return &outputBuffer, err
		}

		response, err := client.Do(request)
		if err != nil {
			fmt.Println("Error making request:", err)
			return &outputBuffer, err
		}

		fmt.Println(response.StatusCode)
		fmt.Println("---------------------------")
	}
	return &outputBuffer, nil
}
