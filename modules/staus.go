package modules

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

// @todo make uniform log system

var tasks []*DownloadInfo

func TaskMonitorProgress() {

	go func() {
		for _, task := range tasks {
			for msg := range task.TaskStatus {
				fmt.Println(msg)
			}
		}
	}()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {

		case <-ticker.C:
			if len(tasks) == 0 {
				return
			}
			var messages []string
			for index, task := range tasks {
				if <-task.TaskEnd {
					removeTask(index)
				}
				message := <-task.TaskStatus
				messages = append(messages, message)
			}

			result := strings.Join(messages, "\r\n")
			//	clearConsole()
			println(result)

		}
	}

}

func AddTask(info *DownloadInfo) {
	tasks = append(tasks, info)
}

func removeTask(index int) {
	tasks = append(tasks[:index], tasks[index+1:]...)
}

func clearConsole() {
	var clearCommand *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		clearCommand = exec.Command("cmd", "/c", "cls")
	default:
		clearCommand = exec.Command("clear")
	}

	clearCommand.Stdout = os.Stdout
	clearCommand.Run()
}
