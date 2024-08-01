package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Generate monitoring report",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			generateReport(args[0])
		} else {
			generateAllReports()
		}
	},
}

func generateReport(url string) {
	loadData()
	var reportContent string
	for _, data := range monitorData {
		if data.URL == url {
			reportContent += fmt.Sprintf("URL: %s, Response Time: %s, Timestamp: %s\n", data.URL, data.ResponseTime, data.Timestamp)
		}
	}
	err := os.WriteFile("report.txt", []byte(reportContent), 0644)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
	}
}

func generateAllReports() {
	loadData()
	var reportContent string
	for _, data := range monitorData {
		reportContent += fmt.Sprintf("URL: %s, Response Time: %s, Timestamp: %s\n", data.URL, data.ResponseTime, data.Timestamp)
	}
	err := os.WriteFile("report.txt", []byte(reportContent), 0644)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
	}
}

func loadData() {
	file, err := os.ReadFile("monitor_data.json")
	if err == nil {
		_ = json.Unmarshal(file, &monitorData)
	}
}
