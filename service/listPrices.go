package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const baseURL = "https://api.btcmarkets.net"

var apiPath = "/v3/markets/BTC-AUD/ticker"

func GetPrices() bool {
	fmt.Printf("Get price is running \n")

	responseBody, err := makeHTTPCall("GET", apiPath, "", "")
	if err != nil {
		log.Fatalln(err)
	}
	//read the response body

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(responseBody)
	return true
}

func makeHTTPCall(method string, path string, query string, body string) (string, error) {

	url := baseURL + apiPath
	if query != "" {
		url += "?" + query
	}

	var request *http.Request
	if body != "" {
		request, _ = http.NewRequest(method, url, strings.NewReader(body))
	} else {
		request, _ = http.NewRequest(method, url, nil)
	}

	response, err := (&http.Client{}).Do(request)
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", nil
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return "", fmt.Errorf("HTTP %d %v", response.StatusCode, response)
	}

	responseBody := string(data)
	return responseBody, nil

}
