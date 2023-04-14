//go:generate mockgen -source=./openai.go -destination=../mock/openai.go -package=mock
package repository

import (
	"bytes"
	"dojogpt/model"
	"encoding/json"
	"io"
	"net/http"
)

type OpenAI interface {
	Completions(r *model.CompletionsRequest) (*model.CompletionsResponse, error)
}

type openai struct {
	url    string
	token  string
	client *http.Client
}

func NewOpenAI(
	url string,
	token string,
	client *http.Client,
) OpenAI {
	return openai{
		url:    url,
		token:  "Bearer " + token,
		client: client,
	}
}

func (o openai) Completions(r *model.CompletionsRequest) (*model.CompletionsResponse, error) {
	jr, err := r.JSON()
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(jr)

	req, err := http.NewRequest(http.MethodPost, o.url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", o.token)

	resp, err := o.client.Do(req)
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	cr := &model.CompletionsResponse{}
	err = json.Unmarshal(b, cr)
	if err != nil {
		return nil, err
	}

	return cr, nil
}
