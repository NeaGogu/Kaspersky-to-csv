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

	addWebsToFile(websiteArr, delimiter, &needChange)
	addAppsToFile(appsArr, delimiter, &needChange)

	fmt.Println("The following account passwords need to be changed because they contain the delimiter:")

	for index, entry := range needChange {
		fmt.Printf("\t%d. %v -> %v\n", index+1, entry.name, entry.url)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
