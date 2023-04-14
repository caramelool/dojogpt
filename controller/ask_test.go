package controller_test

import (
	"dojogpt/controller"
	"dojogpt/mock"
	"dojogpt/model"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAskAGoodQuestion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockOpenAI(ctrl)

	// req := model.NewCompletionsRequest("good question")
	m.EXPECT().Completions(
		gomock.Any(),
	).Return(
		&model.CompletionsResponse{
			Choices: []model.CompletionsChoice{
				{
					Message: model.CompletionsMessage{
						Content: "yes, this is a good question",
					},
				},
			},
		}, nil,
	)

	c := controller.NewAsk(m)
	a := c.MakeQuestion("good question")
	assert.Equal(t, "yes, this is a good question", a)
}

func TestAskABadQuestion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockOpenAI(ctrl)

	// req := model.NewCompletionsRequest("bad question")
	m.EXPECT().Completions(
		gomock.Any(),
	).Return(
		nil, assert.AnError,
	)

	c := controller.NewAsk(m)
	a := c.MakeQuestion("bad question")
	assert.Equal(t, "I could not anwser the question! Sorry!!!", a)
}
