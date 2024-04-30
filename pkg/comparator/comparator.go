package comparator

import (
	"JsonToStruct/pkg/fetcher"
	"fmt"
	"log"
	"reflect"
)

func compareAny(a, b interface{}, path string, differences *map[string][]string, totalFieldsCompared *int) {
	//a variable to store the total number of fields that were compared whether they were equal
	switch aValue := a.(type) {
	case map[string]interface{}:
		bValue, ok := b.(map[string]interface{})
		if !ok {
			(*differences)[path] = []string{fmt.Sprintf("%v", a), fmt.Sprintf("%v", b)}
			return
		}
		for key, valA := range aValue {
			valB, exists := bValue[key]
			prefix := ""
			if path != "" {
				prefix = "."
			}
			if !exists {
				(*differences)[path+prefix+key] = []string{fmt.Sprintf("%v", valA), "nil"}
			} else {
				compareAny(valA, valB, path+prefix+key, differences, totalFieldsCompared)
			}
		}

		for key := range bValue {
			if _, exists := aValue[key]; !exists {
				prefix := ""
				if path != "" {
					prefix = "."
				}
				(*differences)[path+prefix+key] = []string{"nil", fmt.Sprintf("%v", bValue[key])}
			}
		}
	case []interface{}:
		bValue, ok := b.([]interface{})
		if !ok || len(aValue) != len(bValue) {
			(*differences)[path] = []string{fmt.Sprintf("%v", a), fmt.Sprintf("%v", b)}
			return
		}
		for i, valA := range aValue {
			compareAny(valA, bValue[i], fmt.Sprintf("%s[%d]", path, i), differences, totalFieldsCompared)
		}
	default:
		*totalFieldsCompared++
		if !reflect.DeepEqual(a, b) {
			(*differences)[path] = []string{fmt.Sprintf("%v", a), fmt.Sprintf("%v", b)}
		}
	}
}

func HandleComparison(mode, firstURL, secondURL string, params map[string]string, allDifferences *[]map[string][]string, totalFieldsCompared *int) {
	var firstResponse, secondResponse map[string]interface{}
	var err error

	if mode == "endpoints" {
		firstResponse, err = fetcher.FetchEndPoints(firstURL, params)
		if err != nil {
			log.Printf("Failed to fetch from %s: %v", firstURL, err)
			return
		}

		secondResponse, err = fetcher.FetchEndPoints(secondURL, params)
		if err != nil {
			log.Printf("Failed to fetch from %s: %v", secondURL, err)
			return
		}
	} else if mode == "files" {
		firstResponse, err = fetcher.FetchJSONsFromFile("first.json")
		if err != nil {
			log.Printf("Failed to fetch actions: %v", err)
			return
		}

		secondResponse, err = fetcher.FetchJSONsFromFile("second.json")
		if err != nil {
			log.Printf("Failed to fetch actions: %v", err)
			return
		}
	} else {
		log.Fatalf("Invalid mode: %s", mode)
	}

	differences, totalFieldsCompared := CompareActionsResponses(firstResponse, secondResponse, totalFieldsCompared)
	if len(differences) > 0 {
		*allDifferences = append(*allDifferences, differences)
	}
}

func CompareActionsResponses(response1 map[string]interface{}, response2 map[string]interface{}, totalComp *int) (map[string][]string, *int) {
	result := make(map[string][]string)
	compareAny(response1, response2, "", &result, totalComp)
	return result, totalComp
}
