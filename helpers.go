package main

import (
	"bufio"
	"log"
	"strings"
)

func readLine(s *bufio.Scanner, returnVal *string) {
	s.Scan()
	*returnVal = s.Text()
}

func grabEntry(text string) string {

	index := strings.Index(text, ":")

	if index == -1 {
		return ""
	}

	temp := strings.Trim(text[index:], ":")

	log.Printf("Adding: %s -> %s", text[:index], temp)

	return temp
}
