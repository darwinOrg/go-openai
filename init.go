package dgpt

import (
	"github.com/sashabaranov/go-openai"
	"log"
	"os"
)

var client *openai.Client

func init() {
	apiKey, ok := os.LookupEnv("API_KEY")
	if !ok {
		log.Fatal("lookup API_KEY env fail")
	}

	client = openai.NewClient(apiKey)
}
