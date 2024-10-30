package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"thewiidownloader/modules"
	"time"
)

func main() {
	modules.LoadConfig()

	if !modules.ConfigLoaded {
		panic("Config not loaded")
	}

	if modules.GetEnv("debug") == "1" {
		modules.TestGetLinks()
		return
	}

	config := modules.Config

	links, err := modules.ReadLinks(config.LinksFilePath)

	if err != nil {
		fmt.Println("Error reading links:", err)
		return
	}

	concurrentDownloads := make(chan struct{}, 8) // Limit the downloads
	stopNewDownloads := make(chan struct{})
	allDownloadsComplete := make(chan struct{})
	var running = 0

	var wg sync.WaitGroup
	reader := bufio.NewReader(os.Stdin)

	// exit task
	go func() {
		for {
			fmt.Print("Type 'exit' to quit: ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if input == "exit" {
				close(stopNewDownloads)
				fmt.Printf("Stopping new downloads and waiting for ongoing downloads to complete [%d]\n", running)
				return
			}
			time.Sleep(15 * time.Second)
		}
	}()

	// main task system
	go func() {
		for _, link := range links {
			select {
			case <-stopNewDownloads:
				fmt.Println("No more new downloads will be started.")
				goto WaitForDownloads
			case concurrentDownloads <- struct{}{}:
				wg.Add(1)
				running += 1
				go func(link string) {
					// close wg and remove from concurrentDownloads list
					defer wg.Done()
					defer func() { <-concurrentDownloads }()

					download := &modules.DownloadInfo{
						Url:        link,
						ProgressCh: make(chan int64),
						DoneCh:     make(chan bool),
						TaskStatus: make(chan string),
						TaskEnd:    make(chan bool),
						Name:       fixName(path.Base(link)),
					}

					defer close(download.TaskEnd)
					defer close(download.TaskStatus)

					// taskId := strconv.Itoa(i) + download.Name

					go modules.MonitorProgress(download)
					// go modules.AddTask(download)

					success := modules.DownloadFile(download)
					if success == 0 {
						running -= 1
						return
					}

					time.Sleep(1 * time.Millisecond)
					modules.UnzipTask(download)
					time.Sleep(1 * time.Millisecond)
					modules.CovertRVZToISO(download)
					time.Sleep(1 * time.Millisecond)
					modules.ConvertToWBFS(download)

					// fmt.Printf("[from main] Done [%d]\n", download.Name)

					running -= 1
				}(link)
			}
		}

	WaitForDownloads:
		wg.Wait()
		close(allDownloadsComplete)
	}()

	// go modules.TaskMonitorProgress()

	<-allDownloadsComplete

	fmt.Println("All downloads completed. Exiting...")
}

func fixName(name string) string {
	decoded, err := url.QueryUnescape(name)
	if err != nil {
		// if error then the name will be added when converted to wbfs anyway, so it doesn't matter
		return strconv.FormatInt(time.Now().UnixNano(), 10)
	}

	invalidChars := regexp.MustCompile(`[<>:"/\\|?*%]`)
	safeFileName := invalidChars.ReplaceAllString(decoded, "_")
	safeFileName = strings.TrimSpace(safeFileName)

	return safeFileName
}
