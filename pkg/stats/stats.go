package stats

import (
	"fmt"
	"log"
	"os"
)

func CalculateDifferencePercentage(diffs []map[string][]string, totalFieldsCompared int) {
	totalDifferences := 0
	for _, diffMap := range diffs {
		totalDifferences += len(diffMap)
	}

	var percentage float64
	if totalFieldsCompared > 0 {
		percentage = (float64(totalDifferences) / float64(totalFieldsCompared)) * 100
	}

	file, err := os.Create("stats.txt")
	if err != nil {
		log.Fatalf("Failed to create stats file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("Total fields compared: %d\nTotal differences: %d\n---------------------------- \nDifference percentage: %.2f%%\n", totalFieldsCompared, totalDifferences, percentage))
	if err != nil {
		log.Fatalf("Failed to write to stats file: %v", err)
	}
}
