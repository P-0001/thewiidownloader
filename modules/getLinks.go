package modules

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	http "github.com/bogdanfinn/fhttp"
	"io"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"thewiidownloader/tls"
	"time"
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

//

// @todo get links from website

const baseUrl = "https://vimm.net/vault/"

func TestGetLinks() {

	client := tls.GetClient()

	data := make(map[string]string)

	//                  this is others for some reason
	paths := [27]string{"?p=list&system=Wii&section=number", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	vaultRegex := regexp.MustCompile(`/vault/\d+`)

	for i, path := range paths {
		url := ""
		if i == 0 {
			url = baseUrl + path
		} else {
			url = baseUrl + "Wii/" + path
			time.Sleep(time.Second * 5)
		}

		fmt.Println(url)

		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}

		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36")
		req.Header.Set("Referer", "https://www.google.com/")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error making request:", err)
			return
		}

		doc, err := goquery.NewDocumentFromReader(resp.Body)

		if err != nil {
			fmt.Println("Error parsing response:", err)
			break
		}

		doc.Find("a").Each(func(index int, item *goquery.Selection) {
			// Get the href attribute
			href, exists := item.Attr("href")

			// item.Find()

			if !exists {
				return
			}

			if !vaultRegex.MatchString(href) {
				return
			}

			// Get the name of the link
			name := item.Text()

			id := strings.Split(href, "vault/")[1]

			if id == "" {
				fmt.Printf(" ERROR: Link: %s, Name: %s\n", href, name)

				return
			}

			data[id] = name
		})

		//

		// no leak in for loop
		resp.Body.Close()

		break
	}

	getMediaID(data)

}

func getMediaID(data map[string]string) {

	client := tls.GetClient()

	for id, _ := range data {
		x := rand.Intn(1000)
		url := baseUrl + id + "?q=" + strconv.Itoa(x)

		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}

		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36")
		req.Header.Set("Referer", "https://www.google.com/")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error making request:", err)
			return
		}

		bodyBytes, _ := io.ReadAll(resp.Body)

		body := string(bodyBytes)

		// fmt.Println(body)

		startIndex := strings.Index(body, `var allMedia = `)

		if startIndex == -1 {
			fmt.Println("allMedia not found " + body)
			break
		}

		jsonStart := body[startIndex+len("var allMedia ="):]

		// Find the end of the JSON object (until the semicolon)
		endIndex := strings.Index(jsonStart, ";")
		if endIndex == -1 {
			fmt.Println("JSON object not properly terminated")
			return
		}

		// Extract the JSON object
		jsonStr := strings.TrimSpace(jsonStart[:endIndex])

		var gameData GameData

		err = json.Unmarshal([]byte(jsonStr), &gameData)

		fmt.Println(gameData[0].ID)

		resp.Body.Close()

		break
	}
}

type linkGetter struct {
	// info *DownloadInfo
}

func (info *linkGetter) Write() {
	return
}
