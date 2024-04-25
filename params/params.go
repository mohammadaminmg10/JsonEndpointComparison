package params

import (
	"encoding/json"
	"os"
)

type Param struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func LoadParams() []Param {
	file, err := os.Open("params/params.json")
	if err != nil {
		return []Param{}
	}
	defer file.Close()

	var params []Param
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&params)
	if err != nil {
		return []Param{}
	}

	return params
}

func GetParams(p Param) map[string]string {
	paramsMap := make(map[string]string)
	paramsMap[p.Key] = p.Value

	return paramsMap
}
