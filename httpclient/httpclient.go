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
	request, err := http.NewRequest("GET", c.baseURL, nil)
	if err != nil {
		return nil, err
	}

	// add headers to the request
	for key, value := range headers {
		request.Header.Add(key, value)
	}

	response, err := c.sendRequestWithRetry(request)
	if err != nil {
		return nil, err
	}

	// if response is an error (not a 200)
	if response.StatusCode > 299 {
		return nil, errors.New(response.Status)
	}
	// read the body as an array of bytes
	responseBody, err := ioutil.ReadAll(response.Body)
	return responseBody, err
}

func (c *Client) Post(headers map[string]string, body []byte) ([]byte, error) {

	uri, err := url.Parse(c.baseURL)
	if err != nil {
		return nil, err
	}
	c.baseURL = uri.String()

	// create a new post request
	request, err := http.NewRequest("POST", c.baseURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// add headers to the request
	for key, value := range headers {
		request.Header.Add(key, value)
	}

	response, err := c.sendRequestWithRetry(request)
	if err != nil {
		return nil, err
	}

	// if response is an error (not a 200)
	if response.StatusCode > 299 {
		return nil, errors.New(response.Status)
	}

	// read the body as an array of bytes
	responseBody, err := ioutil.ReadAll(response.Body)
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
	request, err := http.NewRequest("DELETE", c.baseURL, nil)
	if err != nil {
		return err
	}

	// add headers to the request
	for key, value := range headers {
		request.Header.Add(key, value)
	}

	response, err := c.sendRequestWithRetry(request)
	if err != nil {
		return err
	}

	// if response is an error (not a 204)
	if response.StatusCode != 204 {
		return errors.New(response.Status)
	}

	return nil
}

func (c *Client) sendRequestWithRetry(request *http.Request) (*http.Response, error) {

	retryCount := 0
	retry := true
	var response *http.Response
	var err error
	var data []byte

	// check if body is not null. We need to store it in order to be used in case of retry
	if request.Body != nil {
		data, err = ioutil.ReadAll(request.Body)
		if err != nil {
			return nil, err
		}
		err = request.Body.Close()
		if err != nil {
			return nil, err
		}
	}

	// if we need to retry and we have not exceed the retry limit
	for retry && retryCount <= maxRetries {

		// if we need to retry then wait
		if retryCount > 0 {
			// sleeping time is calculated has 1.5^number of retry multiplied 500
			sleepingTime := math.Pow(1.5, float64(retryCount)) * float64(500)
			time.Sleep(time.Duration(sleepingTime))
		}

		// populate the body
		request.Body = ioutil.NopCloser(bytes.NewReader(data))
		// send the request
		response, err = c.HTTPClient.Do(request)
		if err != nil {
			return nil, err
		}

		// based on status code, do we need a retry?
		if response.StatusCode == http.StatusTooManyRequests || response.StatusCode == http.StatusInternalServerError || response.StatusCode == http.StatusBadGateway || response.StatusCode == http.StatusServiceUnavailable || response.StatusCode == http.StatusGatewayTimeout {
			retryCount = retryCount + 1
			retry = true
		} else {
			retry = false
		}
	}

	if request.Body != nil {
		// close the original body, we don't need it anymore
		if err := request.Body.Close(); err != nil {
			return nil, err
		}
	}

	return response, nil

}
