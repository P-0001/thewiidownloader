package modules

import (
	"bufio"
	"os"
	"strings"
)

func ReadLinks(filePath string) ([]string, error) {
	file, err := os.Open(filePath)

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

// @todo get links from website

const baseUrl = "https://vimm.net/vault/"

func TestGetLinks() {
	//	urlPaths := {"?p=list&system=Wii&section=number"}
}
