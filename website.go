package main

import (
	"bufio"
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
	return "Websites" + delimiter + w.websiteName + delimiter + w.url + delimiter + w.login + delimiter + w.password
}

func readWebsites(scanner *bufio.Scanner) []Website {
	var websiteRet []Website // list of websites to be returned

	for scanner.Scan() {
		text := scanner.Text()

		if text == "Applications" || text == "Notes" {
			return websiteRet
		}

		if text == "" || text == "---" {
			continue
		}

		website := Website{}
		website.websiteName = text[strings.Index(text, ":")+2:]

		readLine(scanner, &text)
		website.url = text[strings.Index(text, ":")+2:]

		readLine(scanner, &text)
		website.loginName = text[strings.Index(text, ":")+2:]

		readLine(scanner, &text)
		website.login = text[strings.Index(text, ":")+2:]

		readLine(scanner, &text)
		website.password = text[strings.Index(text, ":")+2:]

		readLine(scanner, &text)
		website.comment = text[strings.Index(text, ":")+2:]

		// append the website to the list of websites
		websiteRet = append(websiteRet, website)

	}

	return websiteRet
}
