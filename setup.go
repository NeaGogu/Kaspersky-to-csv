package main

import (
	"flag"
	"log"
	"strings"
)

var (
	SourcePath string
	DestPath   string
	Delimiter  string
)

func Setup() {
	flag.StringVar(&SourcePath, "src", "", "REQUIRED: Kaspersky Source file")
	flag.StringVar(&DestPath, "des", "./KSP_parsed.csv", "Destination file")
	flag.StringVar(&Delimiter, "delimiter", ",", "Delimiter in the CSV file")

	flag.Parse()

	if SourcePath == "" {
		log.Fatal("Source file not specified")
	}

	if !strings.HasSuffix(DestPath, ".csv") {
		DestPath += ".csv"
	}
}
