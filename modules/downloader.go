package modules

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"sync/atomic"
	"time"
)

const (
	linksFilePath = "new-links.txt"
)

func DownloadFile(info *DownloadInfo) {
	defer close(info.DoneCh) // Ensure doneCh is closed when function exits

	fileName := info.Name

	filePath := filepath.Join("C:\\Users\\chrsh\\OneDrive\\Desktop\\code\\go\\downloader\\downloads", fileName)

	info.LastFilePath = filePath

	// Check if file already exists
	if _, err := os.Stat(filePath); err == nil {
		fmt.Printf("File %s already exists. Skipping download.\n", fileName)
		removeLinkFromFile(info.Url)
		return
	}

	// Start download
	resp, err := http.Get(info.Url)
	if err != nil {
		fmt.Println("Error downloading:", info.Url, err)
		return
	}

	defer resp.Body.Close()

	// Get the file size
	info.totalBytes = resp.ContentLength
	info.startTime = time.Now()

	// Create the output file
	out, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer out.Close()

	// Track the progress of the download
	progressReader := io.TeeReader(resp.Body, progressWriter(info))

	// Copy the response body to the file
	_, err = io.Copy(out, progressReader)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	err = removeLinkFromFile(info.Url)

	if err != nil {
		fmt.Println("Error removing link from file:", err)
	}

	// fmt.Printf("Download complete: %s\n", info.url)
}

func removeLinkFromFile(linkToRemove string) error {
	// Open the file for reading
	file, err := os.Open(linksFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a slice to store the links (lines)
	var links []string

	// Scan through the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		link := scanner.Text()

		// Add the link to the slice if it's not the one to remove
		if link != linkToRemove {
			links = append(links, link)
		}
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		return err
	}

	// Open the file for writing (truncate to clear existing content)
	file, err = os.OpenFile(linksFilePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the filtered links back to the file
	writer := bufio.NewWriter(file)
	for _, link := range links {
		_, err := writer.WriteString(link + "\n")
		if err != nil {
			return err
		}
	}
	// Flush the buffered writer
	err = writer.Flush()
	if err != nil {
		return err
	}

	fmt.Printf("Removed links from %s.\n", linksFilePath)

	return nil
}

func MonitorProgress(info *DownloadInfo) {
	randomNum := rand.Intn(5) + 1
	ticker := time.NewTicker(time.Duration(randomNum) * time.Second)
	defer ticker.Stop()

	var lastBytes int64
	var lastTime time.Time
	ewma := 0.0 // Exponentially Weighted Moving Average for speed

	for {
		select {
		case <-ticker.C:
			now := time.Now()
			downloaded := atomic.LoadInt64(&info.downloaded)

			if !lastTime.IsZero() {
				deltaBytes := float64(downloaded - lastBytes)
				deltaTime := now.Sub(lastTime).Seconds()
				instantSpeed := deltaBytes / deltaTime

				// Update EWMA speed
				if ewma == 0 {
					ewma = instantSpeed
				} else {
					ewma = 0.2*instantSpeed + 0.8*ewma
				}

				percent := float64(downloaded) / float64(info.totalBytes) * 100
				remainingBytes := float64(info.totalBytes - downloaded)
				var remainingTime time.Duration
				if ewma > 0 {
					remainingTime = time.Duration(remainingBytes/ewma) * time.Second
				} else {
					remainingTime = time.Duration(0)
				}

				ts := "[" + timeStamp() + "]"
				fmt.Printf("%s %s: %.2f%% complete. Speed: %.2f MB/s. ETA: %s\n",
					ts, info.Name, percent, ewma/(1024*1024), remainingTime.Round(time.Second))
			}

			lastBytes = downloaded
			lastTime = now

		case <-info.DoneCh:
			fmt.Printf("Download complete: %s\n", info.Name)
			return
		}
	}
}

func timeStamp() string {
	now := time.Now()
	return now.Format("03:04:05")
}

func progressWriter(info *DownloadInfo) io.Writer {
	return &progressTracker{info: info}
}

type progressTracker struct {
	info *DownloadInfo
}

func (pt *progressTracker) Write(p []byte) (int, error) {
	n := len(p)
	atomic.AddInt64(&pt.info.downloaded, int64(n))
	return n, nil
}
