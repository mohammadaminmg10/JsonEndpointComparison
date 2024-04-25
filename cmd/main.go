package main

import (
	param "JsonToStruct/params"
	"JsonToStruct/pkg/comparator"
	_ "JsonToStruct/pkg/fetcher"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	mode := flag.String("mode", "endpoints", "Specify 'endpoints' or 'files'")
	firstURL := flag.String("firstURL", "", "First endpoint URL")
	secondURL := flag.String("secondURL", "", "Second endpoint URL")
	flag.Parse()

	if *mode == "endpoints" && (*firstURL == "" || *secondURL == "") {
		log.Fatalf("Please specify both endpoint URLs")
	}

	parameters := param.LoadParams()

	var allDifferences []map[string][]string

	if len(parameters) == 0 {
		comparator.HandleComparison(*mode, *firstURL, *secondURL, nil, &allDifferences)
	} else {
		for _, parms := range parameters {
			params := param.GetParams(parms)
			comparator.HandleComparison(*mode, *firstURL, *secondURL, params, &allDifferences)
		}
	}

	// Write differences to a text file
	file, err := os.Create("result.txt")
	if err != nil {
		log.Fatalf("Failed to create result file: %v", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, diffMap := range allDifferences {
		for key, vals := range diffMap {
			line := fmt.Sprintf("%s: [%s] vs [%s]\n", key, vals[0], vals[1])
			_, err = writer.WriteString(line)
			if err != nil {
				log.Printf("Failed to write to file: %v", err)
				continue
			}
		}
		_, err = writer.WriteString("------\n") // Separator line
		if err != nil {
			log.Printf("Failed to write separator to file: %v", err)
		}
	}
	writer.Flush()
}
