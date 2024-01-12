/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(scrapeCmd)
	scrapeCmd.Flags().StringSlice("targets", []string{""}, "Define Targets")
}

// scrapeCmd represents the scrape command
var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "scrapes sites",
	Long:  `example usage: --targets="GET|https://google.com","GET|https://facebook.com/test"`,
	Run: func(cmd *cobra.Command, args []string) {
		targets, err := cmd.Flags().GetStringSlice("targets")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		Runscrape(targets)
	},
}
