# Go Wrapper for Claude API

This is a Go wrapper for the [Claude API](https://www.anthropic.com/product) by Anthropic. Claude is an AI assistant that can help with a wide range of tasks, including writing, analysis, coding, and more.

## Installation

To use this wrapper, you need to have Go installed on your system. You can install it by following the instructions on the [official Go website](https://golang.org/doc/install).

Once you have Go installed, you can install the wrapper by running the following command:

go get github.com/BradPerbs/claude-go


## Usage

To use the wrapper, you'll need an API key from Anthropic. You can sign up for an API key on the [Anthropic website](https://www.anthropic.com/product).

Here's an example of how to use the wrapper:

```go
package main

import (
	"fmt"
	"github.com/BradPerbs/claude-go"
)

func main() {
	// Replace "your-api-key" with your actual API key
	client := claude.NewClient("your-api-key")

	// Send a prompt to Claude
	prompt := "What is the capital of France?"
	response, err := client.SendPrompt(prompt)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(response)
}


This example creates a new Client instance with your API key and sends a prompt to Claude. The SendPrompt method returns the response from Claude as a string.
Configuration
You can configure the behavior of the client by passing optional configuration functions to the NewClient function:


client := claude.NewClient("your-api-key",
    claude.WithBaseURL("https://api.example.com/v1"), // Set a custom base URL
    claude.WithHTTPClient(&http.Client{Timeout: 10 * time.Second}), // Set a custom HTTP client
)


The available configuration functions are:

- WithBaseURL(baseURL string): Sets the base URL for the Claude API.
- WithHTTPClient(httpClient *http.Client): Sets the HTTP client used for making requests to the Claude API.

Contributing
If you find any issues or want to contribute to this project, feel free to open an issue or submit a pull request on the GitHub repository.
License
This project is licensed under the MIT License.