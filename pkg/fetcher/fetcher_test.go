package fetcher

import (
	"testing"
)

func TestFetchActionJSONsFromFile(t *testing.T) {
	actions, err := FetchJSONsFromFile("test.json")
	if err != nil {
		t.Errorf("FetchJSONsFromFile was incorrect, got: %v, want: nil.", err)
	}
	if len(actions["actions"].([]interface{})) != 2 {
		t.Errorf("FetchJSONsFromFile was incorrect, got: %d, want: %d.", len(actions["actions"].([]interface{})), 2)
	}
	if actions["actions"].([]interface{})[0].(map[string]interface{})["id"] != 1 {
		t.Errorf("FetchJSONsFromFile was incorrect, got: %d, want: %d.", actions["actions"].([]interface{})[0].(map[string]interface{})["id"], 1)
	}
	if actions["actions"].([]interface{})[1].(map[string]interface{})["id"] != 2 {
		t.Errorf("FetchJSONsFromFile was incorrect, got: %d, want: %d.", actions["actions"].([]interface{})[1].(map[string]interface{})["id"], 2)
	}
}

func TestFetchEndPoints(t *testing.T) {
	actions, err := FetchEndPoints("testURL", nil)
	if err != nil {
		t.Errorf("FetchEndPoints was incorrect, got: %v, want: nil.", err)
	}
	if len(actions) == 0 {
		t.Errorf("FetchEndPoints was incorrect, got: %d, want: more than %d.", len(actions), 0)
	}
}
