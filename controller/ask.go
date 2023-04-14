package controller

import (
	"dojogpt/model"
	"dojogpt/repository"
)

type Ask interface {
	MakeQuestion(question string) string
}

type ask struct {
	openai repository.OpenAI
}

func NewAsk(openai repository.OpenAI) Ask {
	return &ask{
		openai: openai,
	}
}

func (a ask) MakeQuestion(question string) string {
	creq := model.NewCompletionsRequest(question)
	cresp, err := a.openai.Completions(creq)
	if err != nil {
		return "I could not anwser the question! Sorry!!!"
	}
	return cresp.Content()
}
