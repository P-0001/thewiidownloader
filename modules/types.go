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
	LinksFilePath   string
	ZippedDirPath   string
	RvzDirPath      string
	IsoDirPath      string
	WBFSDirPath     string
	DolphinToolPath string
	ISOtoWBFSPath   string
	ConcurrentLimit int
}

type GameData []GameDatum

type GameDatum struct {
	ID             int64    `json:"ID"`
	GoodDate       GoodDate `json:"GoodDate"`
	GoodTitle      string   `json:"GoodTitle"`
	Serial         string   `json:"Serial"`
	SortOrder      int64    `json:"SortOrder"`
	Version        string   `json:"Version"`
	Zipped         string   `json:"Zipped"`
	AltZipped      string   `json:"AltZipped"`
	AltZipped2     string   `json:"AltZipped2"`
	GoodHash       string   `json:"GoodHash"`
	GoodMd5        string   `json:"GoodMd5"`
	GoodSha1       string   `json:"GoodSha1"`
	ZippedText     string   `json:"ZippedText"`
	AltZippedText  string   `json:"AltZippedText"`
	AltZipped2Text string   `json:"AltZipped2Text"`
}

type GoodDate struct {
	Date         time.Time `json:"date"`
	TimezoneType int64     `json:"timezone_type"`
	Timezone     string    `json:"timezone"`
}
