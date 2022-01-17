package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readLine(s *bufio.Scanner, returnVal *string) {
	s.Scan()
	*returnVal = s.Text()
}

func main() {
	//	var websiteArr []Website
	var appsArr []Application

	f, err := os.Open("files/ksp2.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		// if strings.Contains(scanner.Text(), "Websites") {
		// 	websiteArr = readWebsites(scanner)
		// }

		if strings.Contains(scanner.Text(), "Applications") {
			appsArr = readApps(scanner)
		}
	}

	var tempToPrint string
	for _, w := range appsArr {
		tempToPrint = tempToPrint + w.toCSV(",") + "\n"
	}

	fmt.Println(tempToPrint)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
