package main

import (
	"bufio"
	"fmt"
	"strings"
)

type Application struct {
	application string
	loginName   string
	login       string
	password    string
	comment     string
}

func (a Application) toCSV(delimiter string) string {
	return "Applications" + delimiter + a.application + delimiter + delimiter + a.login + delimiter + a.password
}

func readApps(scanner *bufio.Scanner) []Application {
	var appRet []Application // list of websites to be returned

	for scanner.Scan() {
		text := scanner.Text()

		if text == "Notes" {
			return appRet
		}

		if text == "" || text == "---" {
			continue
		}

		app := Application{}
		app.application = text[strings.Index(text, ":")+2:]

		readLine(scanner, &text)
		app.loginName = text[strings.Index(text, ":")+2:]

		readLine(scanner, &text)
		app.login = text[strings.Index(text, ":")+2:]

		readLine(scanner, &text)
		app.password = text[strings.Index(text, ":")+2:]

		readLine(scanner, &text)
		app.comment = text[strings.Index(text, ":")+2:]

		// append the website to the list of websites
		appRet = append(appRet, app)

	}

	return appRet
}

func addAppsToFile(apps []Application, delimiter string, needChange *[]toChange) {

	var tempToPrint string

	for _, a := range apps {
		if strings.Contains(a.password, delimiter) {
			*needChange = append(*needChange, toChange{a.application, ""})
			continue
		}

		// TODO: to be changed with adding to file
		tempToPrint = tempToPrint + a.toCSV(delimiter) + "\n"
	}

	fmt.Println(tempToPrint)
}
