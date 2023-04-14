package model

import "encoding/json"

type CompletionsRequest struct {
	Model    string               `json:"model"`
	Messages []CompletionsMessage `json:"messages"`
}

type CompletionsResponse struct {
	Choices []CompletionsChoice `json:"choices"`
}

type CompletionsChoice struct {
	Message CompletionsMessage `json:"message"`
}

type CompletionsMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func NewCompletionsRequest(prompt string) *CompletionsRequest {
	return &CompletionsRequest{
		Model: "gpt-3.5-turbo",
		Messages: []CompletionsMessage{
			{
				Role:    "user",
				Content: "About Golang: " + prompt,
			},
		},
	}
}

func (cr CompletionsRequest) JSON() ([]byte, error) {
	return JSON(cr)
}

func JSON(a any) ([]byte, error) {
	return json.MarshalIndent(a, "", " ")
}

func (cr CompletionsResponse) Content() string {
	return cr.Choices[0].Message.Content
}
