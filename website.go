package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Website struct {
	websiteName string
	url         string
	loginName   string
	login       string
	password    string
	comment     string
}

func (w Website) toCSV(delimiter string) string {
	return "Websites" + delimiter + w.websiteName + delimiter + w.url + delimiter + w.login + delimiter + w.password + "\n"
}

func readWebsites(scanner *bufio.Scanner) []Website {
	var websiteRet []Website // list of websites to be returned

	for scanner.Scan() {
		text := scanner.Text()

		if text == "Applications" || text == "Notes" {
			return websiteRet
		}

		// if text == "" || text == "---" {
		// 	continue
		// }

		if !strings.HasPrefix(text, "Website name") {
			continue
		}

		website := Website{}
		website.websiteName = grabEntry(text)

		readLine(scanner, &text)
		website.url = grabEntry(text)

		readLine(scanner, &text)
		website.loginName = grabEntry(text)

		readLine(scanner, &text)
		website.login = grabEntry(text)

		readLine(scanner, &text)
		website.password = grabEntry(text)

		readLine(scanner, &text)
		website.comment = grabEntry(text)

		// append the website to the list of websites
		// log.Printf("Appending the following website -> \n%v", website)
		fmt.Println()

		websiteRet = append(websiteRet, website)

	}

	return websiteRet
}

func addWebsToFile(websites []Website, delimiter string, needChange *[]toChange, outputFile *os.File) {

	for _, w := range websites {
		if strings.Contains(w.password, delimiter) {
			*needChange = append(*needChange, toChange{w.websiteName, w.url})
			continue
		}

		outputFile.WriteString(w.toCSV(delimiter))

	}
	fmt.Println("Websites added to", outputFile.Name())

}
