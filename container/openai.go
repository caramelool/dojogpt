package container

import (
	"dojogpt/repository"
	"net/http"
	"os"
)

const openaiUrl = "https://api.openai.com/v1/chat/completions"

var openaiToken = os.Getenv("OPENAI_TOKEN")

var openaiRepository repository.OpenAI

func OpenAI() repository.OpenAI {
	if openaiRepository == nil {
		openaiRepository = repository.NewOpenAI(
			openaiUrl,
			openaiToken,
			http.DefaultClient,
		)
	}
	return openaiRepository
}
