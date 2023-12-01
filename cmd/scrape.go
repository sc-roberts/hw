/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

// scrapeCmd represents the scrape command
var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "scapes sites",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("scrape called")
	},
}

type outp struct {
	url    string
	method string
	status int
}

func init() {
	rootCmd.AddCommand(scrapeCmd)
	method := ""
	target := ""

	scrapeCmd.Flags().StringVarP(&method, "method", "m", "get", "Method (Defaults to GET)")
	scrapeCmd.Flags().StringVarP(&target, "target", "t", "testestest", "URL target (Required)")
	fmt.Println(target)
	request, err := http.NewRequest(method, target, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Client Request Error", err)
		return
	}
	defer response.Body.Close()

	scrapeout := outp{
		url:    target,
		status: response.StatusCode,
	}

	fmt.Println("URL:", scrapeout.url, ", Status:", scrapeout.status)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scrapeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scrapeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
