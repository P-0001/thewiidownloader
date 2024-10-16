package modules

import "time"

type DownloadInfo struct {
	Url          string
	totalBytes   int64
	downloaded   int64
	startTime    time.Time
	estimated    time.Duration
	ProgressCh   chan int64
	DoneCh       chan bool
	Name         string
	LastFilePath string
}
