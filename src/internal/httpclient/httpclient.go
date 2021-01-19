package httpclient

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

var httpClient *http.Client

const requestTimeout = 30

func init() {
	httpClient = createHTTPClient()
}

// createHTTPClient for connection re-use
func createHTTPClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			MaxConnsPerHost: 1, //stop creation of unlimited number of connection
		},
		Timeout: time.Duration(requestTimeout) * time.Second,
	}
	return client
}

//Get send an http get request using the url passed through
//it also accept a list of headers option to add to the request
func Get(url string, headers map[string]string, queryParams map[string]string) ([]byte, error) {

	// add parameters to the url
	url = url + "?"
	for key, value := range queryParams {
		url = url + key + "=" + value + "&"
	}

	// create a new get request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// add headers to the request
	for key, value := range headers {
		req.Header.Add(key, value)
	}
	println(url)
	// send the http request
	resp, err := httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	// defer and close the body stream
	defer resp.Body.Close()

	// if response is an error (not a 200)
	if resp.StatusCode > 299 {
		return nil, errors.New(resp.Status)
	}

	// read the body as an array of bytes
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}
