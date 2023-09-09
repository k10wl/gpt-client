package gpt_client

import (
	"encoding/json"
	"fmt"
)

type Usage struct {
	PromptTokens     uint `json:"prompt_tokens"`
	CompletionTokens uint `json:"completion_tokens"`
	TotalTokens      uint `json:"total_tokens"`
}

type Choices struct {
	Index        uint    `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type Completion struct {
	ID      string    `json:"id"`
	Object  string    `json:"object"`
	Created int64     `json:"created"`
	Model   string    `json:"model"`
	Choices []Choices `json:"choices"`
	Usage   `          json:"usage"`
}

type Message struct {
	Role    string `json:"role,omitempty"    binding:"required"`
	Content string `json:"content,omitempty" binding:"required"`
}

type Request struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float32   `json:"temperature"`
}

func (c *Client) TextCompletion(message *[]Message) (*Completion, error) {
	payloadData, err := json.Marshal(Request{
		Model:       DefaultModel,
		Messages:    *message,
		Temperature: 1,
	})
	if err != nil {
		return nil, err
	}

	responseData, err := c.makePostRequest(&payloadData, ChatCompletionRoute)
	if err != nil {
		return nil, err
	}

	var resBody Completion
	if err := json.Unmarshal(responseData, &resBody); err != nil {
		return nil, err
	}

	return &resBody, nil
}

func (c *Client) HasTokensOverflow(message *Message) bool {
	return c.CountMessageTokens(message) > MaxRequestTokens
}

func (c *Client) CountMessageTokens(message *Message) int {
	return len(c.tokenizer.client.Encode(message.Content))
}

func (c *Client) BuildHistory(prevMsgs *[]Message) (*[]Message, error) {
	messages := []Message{}
	tokensUsage := 0

	if len(*prevMsgs) > 0 && (*prevMsgs)[0].Role == "system" {
		messages = append(messages, (*prevMsgs)[0])
		tokensUsage = c.CountMessageTokens(&messages[0])
	}

	for i := len(*prevMsgs) - 1; i >= 0; i-- {
		prevMsg := (*prevMsgs)[i]
		sum := tokensUsage + c.CountMessageTokens(&prevMsg)
		if sum > MaxRequestTokens {
			break
		}

		tokensUsage = sum

		messages = append([]Message{prevMsg}, messages...)
	}

	if len(messages) == 0 {
		return nil, fmt.Errorf("%s", TokensOverflowError)
	}

	return &messages, nil
}
