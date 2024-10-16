package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"
	"sync"
	"thewiidownloader/modules"
	"time"
)

const (
	linksFilePath = "new-links.txt"
)

func main() {
	links, err := readLinks(linksFilePath)
	if err != nil {
		fmt.Println("Error reading links:", err)
		return
	}

	concurrentDownloads := make(chan struct{}, 5) // Limit to 5 concurrent downloads
	stopNewDownloads := make(chan struct{})
	allDownloadsComplete := make(chan struct{})

	var wg sync.WaitGroup
	reader := bufio.NewReader(os.Stdin)

	go func() {
		for {
			fmt.Print("Type 'exit' to quit: ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if input == "exit" {
				close(stopNewDownloads)
				fmt.Println("Stopping new downloads and waiting for ongoing downloads to complete...")
				return
			}
			time.Sleep(15 * time.Second)
		}
	}()

	go func() {
		for _, link := range links {
			select {
			case <-stopNewDownloads:
				fmt.Println("No more new downloads will be started.")
				goto WaitForDownloads
			case concurrentDownloads <- struct{}{}:
				wg.Add(1)
				go func(link string) {
					defer wg.Done()
					defer func() { <-concurrentDownloads }()

					download := &modules.DownloadInfo{
						Url:        link,
						ProgressCh: make(chan int64),
						DoneCh:     make(chan bool),
						Name:       fixName(path.Base(link)),
					}

					// taskId := strconv.Itoa(i) + download.Name

					go modules.MonitorProgress(download)

					modules.DownloadFile(download)
					modules.UnzipTask(download)
					modules.CovertRVZToISO(download)
					modules.ConvertToWBFS(download)
				}(link)
			}
		}

	WaitForDownloads:
		wg.Wait()
		close(allDownloadsComplete)
	}()

	<-allDownloadsComplete
	fmt.Println("All downloads completed. Exiting...")
}

func fixName(name string) string {
	decoded, err := url.QueryUnescape(name)
	if err != nil {
		panic(err)
	}

	invalidChars := regexp.MustCompile(`[<>:"/\\|?*%]`)
	safeFileName := invalidChars.ReplaceAllString(decoded, "_")
	safeFileName = strings.TrimSpace(safeFileName)

	return safeFileName
}

func readLinks(fileName string) ([]string, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var links []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		link := strings.TrimSpace(scanner.Text())

		if link != "" {
			links = append(links, link)
		}
	}

	return links, scanner.Err()
}
