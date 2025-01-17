package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/boyter/scc/v3/processor"
)

type statsProcessor struct{}

func (p *statsProcessor) ProcessLine(job *processor.FileJob, lineContents []byte, currentLine int64, lineType processor.LineType) bool {
	if lineType == processor.LINE_CODE {
		fmt.Print(string(lineContents))
	}
	return true
}

func main() {

	inputFile := flag.String("inputFile", "", "File path to strip comments from")
	language := flag.String("language", "C", "Language of file to strip comments from")

	flag.Parse()

	input, err := os.ReadFile(*inputFile)

	if err != nil {
		log.Fatalf("Failed to ReadFile: %v", err)
	}

	processor.ProcessConstants() // Required to load the language information and need only be done once

	t := &statsProcessor{}
	filejob := &processor.FileJob{
		Language: *language,
		Content:  input,
		Callback: t,
		Bytes:    int64(len(input)),
	}

	processor.CountStats(filejob)
}
