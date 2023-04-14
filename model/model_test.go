package model_test

import (
	"dojogpt/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckRequestModel(t *testing.T) {
	cr := model.NewCompletionsRequest("test")
	assert.Equal(t, "gpt-3.5-turbo", cr.Model)
}

func TestCheckRole(t *testing.T) {
	cr := model.NewCompletionsRequest("test")
	assert.Equal(t, "user", cr.Messages[0].Role)
}

func TestCheckContent(t *testing.T) {
	cr := model.NewCompletionsRequest("test")
	assert.Equal(t, "About Golang: test", cr.Messages[0].Content)
}

func TestCheckMessagesLen(t *testing.T) {
	cr := model.NewCompletionsRequest("test")
	assert.Len(t, cr.Messages, 1)
}
