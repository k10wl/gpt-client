package gpt_client

import (
	"net/http"
	"time"
)

type Client struct {
	apiKey           string
	httpClient       *http.Client
	tokenizer        *Tokenizer
	maxRequestTokens int
}

const (
	DefaultModel = "gpt-3.5-turbo"
	HTTPTimeout  = 2 * time.Minute
	// TODO This should be dynamic
	MaxRequestTokens = 4096

	System    = "system"
	Assistant = "assistant"
	User      = "user"
)

func NewClient(apiKey string) *Client {
	tokenizerClient, err := GetTokenizerForModel(DefaultModel)
	if err != nil {
		panic("Could not initialize tokenizer, cannot proceed")
	}

	return &Client{
		apiKey:           apiKey,
		httpClient:       &http.Client{Timeout: HTTPTimeout},
		tokenizer:        tokenizerClient,
		maxRequestTokens: MaxRequestTokens,
	}
}
