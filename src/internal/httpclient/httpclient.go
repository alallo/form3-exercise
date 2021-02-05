package httpclient

import (
	"bytes"
	"errors"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"time"
)

const requestTimeout = 30
const maxRetries = 10

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	baseURL    string
	HTTPClient HttpClient
}

//CreateHTTPClient creates an HTTPClient to perform an Http request
func CreateHTTPClient(requestURL string) (*Client, error) {
	_, err := url.ParseRequestURI(requestURL)
	if err != nil {
		return nil, err
	}
	return &Client{
		HTTPClient: &http.Client{
			Timeout: time.Duration(requestTimeout) * time.Second,
		},
		baseURL: requestURL,
	}, nil
}

//Get send an http get request using the url passed through
//it also accept a list of headers option to add to the request
func (c *Client) Get(headers map[string]string, queryParams map[string]string) ([]byte, error) {

	// add parameters to the url
	v := url.Values{}
	for key, value := range queryParams {
		v.Add(key, value)
	}
	uri, err := url.Parse(c.baseURL)
	if err != nil {
		return nil, err
	}
	uri.RawQuery = v.Encode()
	c.baseURL = uri.String()

	// create a new get request
	req, err := http.NewRequest("GET", c.baseURL, nil)
	if err != nil {
		return nil, err
	}

	// add headers to the request
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	retryCount := 0
	retry := true
	var resp *http.Response

	for retry && retryCount <= maxRetries {
		resp, err = c.HTTPClient.Do(req)
		if err != nil {
			return nil, err
		}
		retry = retryRequest(resp.StatusCode, retryCount)
		retryCount = retryCount + 1
	}

	// defer and close the body stream
	defer resp.Body.Close()
	// if response is an error (not a 200)
	if resp.StatusCode > 299 {
		return nil, errors.New(resp.Status)
	}
	// read the body as an array of bytes
	responseBody, err := ioutil.ReadAll(resp.Body)
	return responseBody, err
}

func (c *Client) Post(headers map[string]string, body []byte) ([]byte, error) {

	uri, err := url.Parse(c.baseURL)
	if err != nil {
		return nil, err
	}
	c.baseURL = uri.String()

	// create a new post request
	req, err := http.NewRequest("POST", c.baseURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// add headers to the request
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	retryCount := 0
	retry := true
	var resp *http.Response

	for retry && retryCount <= maxRetries {
		resp, err = c.HTTPClient.Do(req)
		if err != nil {
			return nil, err
		}
		retry = retryRequest(resp.StatusCode, retryCount)
		retryCount = retryCount + 1
	}

	// defer and close the body stream
	defer resp.Body.Close()

	// if response is an error (not a 200)
	if resp.StatusCode > 299 {
		return nil, errors.New(resp.Status)
	}

	// read the body as an array of bytes
	responseBody, err := ioutil.ReadAll(resp.Body)
	return responseBody, err
}

func (c *Client) Delete(headers map[string]string, queryParams map[string]string) error {

	// add parameters to the url
	v := url.Values{}
	for key, value := range queryParams {
		v.Add(key, value)
	}
	uri, err := url.Parse(c.baseURL)
	if err != nil {
		return err
	}
	uri.RawQuery = v.Encode()
	c.baseURL = uri.String()

	// create a new delete request
	req, err := http.NewRequest("DELETE", c.baseURL, nil)
	if err != nil {
		return err
	}

	// add headers to the request
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	retryCount := 0
	retry := true
	var resp *http.Response

	for retry && retryCount <= maxRetries {
		resp, err = c.HTTPClient.Do(req)
		if err != nil {
			return err
		}
		retry = retryRequest(resp.StatusCode, retryCount)
		retryCount = retryCount + 1
	}

	// defer and close the body stream
	defer resp.Body.Close()

	// if response is an error (not a 204)
	if resp.StatusCode != 204 {
		return errors.New(resp.Status)
	}

	return nil
}

func retryRequest(statusCode int, retryCount int) bool {
	if retryCount > 0 {
		sleepingTime := math.Pow(1.5, float64(retryCount)) * float64(500)
		time.Sleep(time.Duration(sleepingTime))
	}

	if statusCode == http.StatusOK {
		return false
	} else if statusCode == http.StatusTooManyRequests {
		return true
	} else if statusCode == http.StatusInternalServerError {
		return true
	} else {
		return false
	}
}
