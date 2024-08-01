package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all monitoring targets",
	Run: func(cmd *cobra.Command, args []string) {
		listTargets()
	},
}

func listTargets() {
	loadConfig()
	for _, target := range targets {
		fmt.Println("Monitoring target:", target.URL)
	}
}
