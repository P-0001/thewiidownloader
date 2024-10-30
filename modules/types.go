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
	TaskEnd      chan bool
	TaskStatus   chan string
	Name         string
	LastFilePath string
}

type ConfigS struct {
	LinksFilePath string
	ZippedDirPath string
	RvzDirPath    string
	IsoDirPath    string
	WBFSDirPath   string
}
