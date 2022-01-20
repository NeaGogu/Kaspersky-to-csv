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

	Setup() // setup Source, Destination, Delimiter

	f, err := os.Open(SourcePath) // TODO: to be specified by user
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

	os.Remove(DestPath) // delete the file if it exists
	outputFile, err := os.OpenFile(DestPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}

	defer outputFile.Close()

	outputFile.WriteString("App" + Delimiter + "name" + Delimiter + "url" + Delimiter + "username" + Delimiter + "password\n")

	addWebsToFile(websiteArr, Delimiter, &needChange, outputFile)
	addAppsToFile(appsArr, Delimiter, &needChange, outputFile)

	outputFile.Sync()

	// notify the user of what needs to be changed
	fmt.Printf("The following account passwords need to be changed because they contain the Delimiter ( %v ):\n", Delimiter)

	for index, entry := range needChange {
		fmt.Printf("\t%d. %v -> %v\n", index+1, entry.name, entry.url)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
