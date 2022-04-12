package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

type Client struct {
	BaseUrl    string
	HTTPClient *http.Client
}

func NewClient(baseUrl string) *Client {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return &Client{}
	}

	return &Client{
		BaseUrl: baseUrl,
		HTTPClient: &http.Client{
			Jar: jar,
		},
	}
}

func (c *Client) sendRequest(req *http.Request) (io.ReadCloser, error) {
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		message, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("error while sending request, status code: %d, error: %s", res.StatusCode, message)
	}

	return res.Body, nil
}

func (c *Client) FormPost(path string, values url.Values) (io.ReadCloser, error) {
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", c.BaseUrl, path),
		strings.NewReader(values.Encode()),
	)

	if err != nil {
		return nil, fmt.Errorf("error while creating FormPost request, error: %s", err.Error())
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return c.sendRequest(req)
}

func (c *Client) Post(path string, data map[string]string, headers map[string]string) (io.ReadCloser, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", c.BaseUrl, path),
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return nil, fmt.Errorf("error while creating Post request, error: %s", err.Error())
	}
	req.Header.Add("Content-Type", "application/json")

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	return c.sendRequest(req)
}

func (c *Client) Get(path string, headers map[string]string) (io.ReadCloser, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/%s", c.BaseUrl, path),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("error while creating Get request, error: %s", err.Error())
	}
	req.Header.Add("Content-Type", "application/json")
	for key, value := range headers {
		req.Header.Add(key, value)
	}
	return c.sendRequest(req)
}
