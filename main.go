package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type toChange struct {
	name string
	url  string
}

func readLine(s *bufio.Scanner, returnVal *string) {
	s.Scan()
	*returnVal = s.Text()
}

func main() {
	// might not need
	var websiteArr []Website
	var appsArr []Application
	var needChange []toChange

	delimiter := "," // TODO: to be specified by user

	f, err := os.Open("files/ksp2.txt") // TODO: to be specified by user
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		if strings.Contains(scanner.Text(), "Websites") {
			websiteArr = readWebsites(scanner)
		}

		if strings.Contains(scanner.Text(), "Applications") {
			appsArr = readApps(scanner)
		}
	}

	// write header to file

	os.Remove("files/FORMATTED.csv") // delete the file if it exists
	outputFile, err := os.OpenFile("files/FORMATTED.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}

	defer outputFile.Close()

	outputFile.WriteString("App,name,url,username,password\n")

	addWebsToFile(websiteArr, delimiter, &needChange, outputFile)
	addAppsToFile(appsArr, delimiter, &needChange, outputFile)

	outputFile.Sync()

	// notify the user of what needs to be changed
	fmt.Println("The following account passwords need to be changed because they contain the delimiter:")

	for index, entry := range needChange {
		fmt.Printf("\t%d. %v -> %v\n", index+1, entry.name, entry.url)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
