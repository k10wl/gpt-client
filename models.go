package gpt_client

import "encoding/json"

type Model struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	OwnedBy string `json:"owned_by"`
}

type ModelsList struct {
	Object string  `json:"object"`
	Data   []Model `json:"data"`
}

func (c *Client) GetModelsList() (*ModelsList, error) {
	body, err := c.makeGetRequest(ModelsListRoute)
	if err != nil {
		return nil, err
	}

	var modelsList ModelsList
	err = json.Unmarshal(body, &modelsList)
	if err != nil {
		return nil, err
	}

	return &modelsList, nil
}
