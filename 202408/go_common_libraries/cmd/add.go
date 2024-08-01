package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [url]",
	Short: "Add a new monitoring target",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		addTarget(url)
	},
}

type MonitorTarget struct {
	URL string `json:"url"`
}

var configFile = "monitor_config.json"
var targets []MonitorTarget

func addTarget(url string) {
	loadConfig()
	targets = append(targets, MonitorTarget{URL: url})
	saveConfig()
	fmt.Println("Added target:", url)
}

func loadConfig() {
	file, err := os.ReadFile(configFile)
	if err == nil {
		_ = json.Unmarshal(file, &targets)
	}
}

func saveConfig() {
	data, _ := json.Marshal(targets)
	_ = os.WriteFile(configFile, data, 0644)
}
