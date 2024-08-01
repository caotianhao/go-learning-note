package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var (
	//monitorCtx    context.Context
	//monitorCancel context.CancelFunc
	//mu            sync.Mutex
	stopProb = 0.05
)

type MonitorData struct {
	URL          string        `json:"url"`
	ResponseTime time.Duration `json:"response_time"`
	Timestamp    time.Time     `json:"timestamp"`
}

var monitorData []MonitorData

var startCmd = &cobra.Command{
	Use:   "start [url]",
	Short: "Start monitoring targets",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			startMonitoring(args[0])
		} else {
			startAllMonitoring()
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

// startMonitoring 启动单个监控任务
func startMonitoring(url string) {
	for {
		checkTarget(url)
		if rand.Float64() < stopProb {
			fmt.Println("Stopping monitoring due to probability check")
			break
		}
	}
	fmt.Println("Started monitoring:", url)
}

// startAllMonitoring 启动所有监控任务
func startAllMonitoring() {
	loadConfig()
	if len(targets) == 0 {
		fmt.Println("No targets to monitor")
		return
	}
	for _, target := range targets {
		startMonitoring(target.URL)
	}
}

// checkTarget 检查指定目标的响应时间
func checkTarget(url string) {
	start := time.Now()
	_, _ = checkPingResponseTime(url)
	elapsed := time.Since(start)

	data := MonitorData{
		URL:          url,
		ResponseTime: elapsed,
		Timestamp:    time.Now(),
	}
	monitorData = append(monitorData, data)
	saveData()
	fmt.Println("Checked target:", url, "Response time:", elapsed)
}

func checkPingResponseTime(address string) (time.Duration, error) {
	// 执行ping命令
	cmd := exec.Command("ping", "-c", "1", address)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return 0, err
	}

	// 解析ping输出，提取时间
	output := out.String()
	re := regexp.MustCompile(`time=(\d+\.\d+) ms`)
	matches := re.FindStringSubmatch(output)
	if len(matches) < 2 {
		return 0, fmt.Errorf("failed to parse ping output: %s", output)
	}

	// 将时间字符串转换为浮点数
	timeMs, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		return 0, err
	}

	// 转换为Duration类型并返回
	return time.Duration(timeMs * float64(time.Millisecond)), nil
}

// saveData 保存监控数据到文件
func saveData() {
	data, err := json.Marshal(monitorData)
	if err != nil {
		fmt.Printf("Error marshaling data: %v\n", err)
		return
	}
	err = os.WriteFile("monitor_data.json", data, 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
	}
}
