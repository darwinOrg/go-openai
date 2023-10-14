package dgpt

import (
	"github.com/sashabaranov/go-openai"
	"os"
)

var client *openai.Client

func init() {
	apiKey := os.Getenv("OPENAI_API_KEY")

	client = openai.NewClient(apiKey)
}
