package fetcher

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

//const (
//	midgardUrl = "https://m-stage.thorswap.net/mu/ohlcv/ETH.ETH?interval=day&count=10"
//	otherUrl   = "http://95.217.108.62:5054/ohlcv/ETH.ETH?interval=day&count=10"
//)

func FetchActionsFromFile(filename string) (map[string]interface{}, error) {
	file, err := os.Open("JsonsToCompare/" + filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var actions map[string]interface{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&actions)
	if err != nil {
		return nil, err
	}

	return actions, nil
}

func FetchActions(url string, params map[string]string) (map[string]interface{}, error) {
	if len(params) != 0 {
		url += "?"
		for key, value := range params {
			url += key + "=" + value + "&"
		}
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "abc")
	req.Header.Set("Referer", "thor")

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Failed to close response body: %v", err)
		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var actions map[string]interface{}
	err = json.Unmarshal(body, &actions)
	if err != nil {
		return nil, err
	}

	return actions, nil
}

//
//func FetchFirstActions(params map[string]string) (map[string]interface{}, error) {
//	return fetchActions(midgardUrl, params)
//}
//
//func FetchSecondActions(params map[string]string) (map[string]interface{}, error) {
//	return fetchActions(otherUrl, params)
//}
