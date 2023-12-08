/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	var targets []string
	rootCmd.AddCommand(scrapeCmd)
	//	scrapeCmd.Flags().StringP("method", "m", "get", "Method (Defaults to GET)")
	//	scrapeCmd.Flags().StringP("target", "t", "testestest", "URL target (Required)")
	scrapeCmd.Flags().StringSliceVarP(&targets, "targets", "f", []string{""}, "Define Targets")
}

// scrapeCmd represents the scrape command
var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "scapes sites",
	Long:  `A longer description`,
	Run:   runscrape,
}

func runscrape(cmd *cobra.Command, args []string) {
	fmt.Println("scrape called")
	targets, _ := cmd.Flags().GetStringSlice("targets")
	fmt.Println("Target Slice:", targets)
	// method, _ := cmd.Flags().GetString("method")
	// target, _ := cmd.Flags().GetString("target")
	// fmt.Println("Method:", method)
	// fmt.Println("Target:", target)

	// request, err := http.NewRequest(method, target, nil)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// client := &http.Client{}
	// response, err := client.Do(request)
	// if err != nil {
	// 	fmt.Println("Client Request Error", err)
	// 	return
	// }
	// defer response.Body.Close()

	// scrapeout := outp{
	// 	url:    target,
	// 	status: response.StatusCode,
	// }

	// fmt.Println("URL:", scrapeout.url, ", Status:", scrapeout.status)

}

type outp struct {
	url    string
	method string
	status int
}
