package gpt_client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

const (
	BaseUrl             = "https://api.openai.com/v1"
	ChatCompletionRoute = "/chat/completions"
	ModelsListRoute     = "/models"
)

func (c *Client) makePostRequest(body *[]byte, route string) ([]byte, error) {
	req, err := http.NewRequest("POST", BaseUrl+route, bytes.NewReader(*body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, fmt.Errorf("HTTP Error: %d", res.StatusCode)
	}

	return io.ReadAll(res.Body)
}

func (c *Client) makeGetRequest(route string) ([]byte, error) {
	req, err := http.NewRequest("GET", BaseUrl+route, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, fmt.Errorf("HTTP Error: %d", res.StatusCode)
	}

	return io.ReadAll(res.Body)
}
