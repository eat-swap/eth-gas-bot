package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func HttpRequest(method, url string, data []byte, headers map[string][]string) ([]byte, error) {
	c := &http.Client{}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		for _, e := range v {
			req.Header.Add(k, e)
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
		return nil, err
	}
	return body, nil
}

func HttpPost(url string, data []byte, headers map[string][]string) ([]byte, error) {
	return HttpRequest("POST", url, data, headers)
}
