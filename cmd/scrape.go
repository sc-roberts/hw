/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(scrapeCmd)
	scrapeCmd.Flags().StringSlice("targets", []string{""}, "Define Targets")
}

// scrapeCmd represents the scrape command
var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "scapes sites",
	Long:  `A longer description`,
	Run:   runscrape,
}

func runscrape(cmd *cobra.Command, args []string) {
	targets, _ := cmd.Flags().GetStringSlice("targets")

	client := &http.Client{}
	for _, targ := range targets {

		splitter := strings.SplitN(targ, "|", 2)
		method := splitter[0]
		target := splitter[1]
		fmt.Println(target, method)
		request, err := http.NewRequest(method, target, nil)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		response, err := client.Do(request)
		fmt.Println(response.StatusCode)
		fmt.Println("---------------------------")
	}
}
