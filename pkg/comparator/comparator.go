package comparator

import (
	"fmt"
	"reflect"
)

func compareAny(a, b interface{}, path string, differences *map[string][]string) {
	switch aValue := a.(type) {
	case map[string]interface{}:
		bValue, ok := b.(map[string]interface{})
		if !ok {
			(*differences)[path] = []string{fmt.Sprintf("%v", a), fmt.Sprintf("%v", b)}
			return
		}
		for key, valA := range aValue {
			valB, exists := bValue[key]
			if !exists {
				(*differences)[path+"."+key] = []string{fmt.Sprintf("%v", valA), "nil"}
			} else {
				compareAny(valA, valB, path+"."+key, differences)
			}
		}
		// Check for keys in b not in a
		for key := range bValue {
			if _, exists := aValue[key]; !exists {
				(*differences)[path+"."+key] = []string{"nil", fmt.Sprintf("%v", bValue[key])}
			}
		}
	case []interface{}:
		bValue, ok := b.([]interface{})
		if !ok || len(aValue) != len(bValue) {
			(*differences)[path] = []string{fmt.Sprintf("%v", a), fmt.Sprintf("%v", b)}
			return
		}
		for i, valA := range aValue {
			compareAny(valA, bValue[i], fmt.Sprintf("%s[%d]", path, i), differences)
		}
	default:
		if !reflect.DeepEqual(a, b) {
			(*differences)[path] = []string{fmt.Sprintf("%v", a), fmt.Sprintf("%v", b)}
		}
	}
}

func CompareActionsResponses(response1, response2 map[string]interface{}) map[string][]string {
	result := make(map[string][]string)
	compareAny(response1, response2, "", &result)
	return result
}
