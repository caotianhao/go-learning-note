package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [url]",
	Short: "Delete a monitoring target",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		deleteTarget(url)
	},
}

func deleteTarget(url string) {
	loadConfig()
	for i, target := range targets {
		if target.URL == url {
			targets = append(targets[:i], targets[i+1:]...)
			saveConfig()
			fmt.Println("Deleted target:", url)
			return
		}
	}
	fmt.Println("Target not found:", url)
}
