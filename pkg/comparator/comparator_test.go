package comparator

import (
	"reflect"
	"testing"
)

func TestCompareActionsResponses(t *testing.T) {
	testCases := []struct {
		name      string
		response1 map[string]interface{}
		response2 map[string]interface{}
		want      map[string][]string
	}{
		{
			name:      "Test Case 1",
			response1: map[string]interface{}{"key1": "value1", "key2": "value2"},
			response2: map[string]interface{}{"key1": "value1", "key2": "value3"},
			want:      map[string][]string{"key2": {"value2", "value3"}},
		},
		{
			name:      "Test Case 2",
			response1: map[string]interface{}{"key1": "value1", "key2": "value2"},
			response2: map[string]interface{}{"key1": "value1", "key2": "value2"},
			want:      map[string][]string{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			var totalFieldsCompared int
			got, _ := CompareActionsResponses(tc.response1, tc.response2, &totalFieldsCompared)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("CompareActionsResponses(%v, %v) = %v; want %v", tc.response1, tc.response2, got, tc.want)
				t.Logf("Got: %v, Want: %v", got, tc.want) // Print the got and want values for debugging
			}
		})
	}
}
