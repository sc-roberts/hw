package cmd

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Runscrape(targets []string) (*bytes.Buffer, error) {
	client := &http.Client{}
	var outputBuffer bytes.Buffer

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
			return &outputBuffer, fmt.Errorf("error executing request: %w", err)
		}
		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if err != nil {
			return &outputBuffer, fmt.Errorf("error reading response body: %w", err)
		}

		//fmt.Println("Body Contents:", string(body)) //For unittest testing...
		outputBuffer.Write(body)

		fmt.Println(response.StatusCode)
		fmt.Println("---------------------------")
	}
	return &outputBuffer, nil
}
