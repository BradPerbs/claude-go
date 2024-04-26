package claude

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	defaultBaseURL = "https://api.anthropic.com/v1"
)

// Client is a client for the Claude API.
type Client struct {
	apiKey    string
	baseURL   string
	httpClient *http.Client
}

// NewClient creates a new Client with the given API key.
func NewClient(apiKey string, options ...func(*Client)) *Client {
	c := &Client{
		apiKey:    apiKey,
		baseURL:   defaultBaseURL,
		httpClient: http.DefaultClient,
	}

	for _, option := range options {
		option(c)
	}

	return c
}

// WithBaseURL sets the base URL for the client.
func WithBaseURL(baseURL string) func(*Client) {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// WithHTTPClient sets the HTTP client for the client.
func WithHTTPClient(httpClient *http.Client) func(*Client) {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// SendPrompt sends a prompt to the Claude API and returns the response.
func (c *Client) SendPrompt(prompt string) (string, error) {
	payload := map[string]interface{}{
		"prompt": prompt,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", c.baseURL+"/complete", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", fmt.Errorf("failed to send prompt: %s", resp.Status)
		}
		return "", fmt.Errorf("failed to send prompt: %s (%s)", resp.Status, body)
	}

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	result, ok := response["result"].(string)
	if !ok {
		return "", errors.New("failed to parse Claude's response")
	}

	return result, nil
}