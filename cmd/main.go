package main

import (
	param "JsonToStruct/params"
	"JsonToStruct/pkg/comparator"
	"JsonToStruct/pkg/fetcher"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
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
	for _, parms := range parameters {
		params := param.GetParams(parms)
		params = map[string]string{}

		differences := make(map[string][]string)

		var firstActions, secondActions map[string]interface{}
		var err error

		if *mode == "endpoints" {
			firstActions, err = fetcher.FetchActions(*firstURL, params)
			if err != nil {
				log.Printf("Failed to fetch actions: %v", err)
				differences["Error"] = []string{err.Error(), ""}
				continue
			}

			secondActions, err = fetcher.FetchActions(*secondURL, params)
			if err != nil {
				log.Printf("Failed to fetch actions: %v", err)
				differences["Error"] = []string{err.Error(), ""}
				continue
			}
		} else if *mode == "files" {
			firstActions, err = fetcher.FetchActionsFromFile("first.json")
			if err != nil {
				log.Printf("Failed to fetch actions: %v", err)
				differences["Error"] = []string{err.Error(), ""}
				continue
			}

			secondActions, err = fetcher.FetchActionsFromFile("second.json")
			if err != nil {
				log.Printf("Failed to fetch actions: %v", err)
				differences["Error"] = []string{err.Error(), ""}
				continue
			}
		} else {
			log.Fatalf("Invalid mode: %s", *mode)
		}

		differences = comparator.CompareActionsResponses(firstActions, secondActions)
		if len(differences) == 0 {
			//no need to add this params to the result
			continue
		}
		differences["ID"] = []string{strconv.Itoa(parms.Id), ""}

		allDifferences = append(allDifferences, differences)
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
