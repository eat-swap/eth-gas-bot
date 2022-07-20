package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func HttpRequest(method, url string, data []byte, headers map[string][]string) ([]byte, error) {
	c := &http.Client{}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	if headers != nil {
		for k, v := range headers {
			for _, e := range v {
				req.Header.Add(k, e)
			}
		}
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("non 2xx status: %s", resp.Status)
	}
	return body, nil
}

func HttpPost(url string, data []byte, headers map[string][]string) ([]byte, error) {
	return HttpRequest("POST", url, data, headers)
}

func HttpGet(url string, headers map[string][]string) ([]byte, error) {
	return HttpRequest("GET", url, nil, headers)
}

func HttpGetWithMultiParams(url string, headers, params map[string][]string) ([]byte, error) {
	url = url + "?"
	for k, v := range params {
		for _, e := range v {
			url = url + k + "=" + e + "&"
		}
	}
	return HttpGet(url, headers)
}

func HttpGetWithParams(url string, headers map[string][]string, params map[string]string) ([]byte, error) {
	url = url + "?"
	for k, v := range params {
		url = url + k + "=" + v + "&"
	}
	return HttpGet(url[:len(url)-1], headers)
}
