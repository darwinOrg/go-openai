package dgpt

import (
	ve "github.com/darwinOrg/go-validator-ext"
	"log"
	"os"
)

var (
	COMPLETION_API_URL = "https://api.openai.com/v1/completions"
	CHAT_API_URL       = "https://api.openai.com/v1/chat/completions"
	headers            map[string]string
)

type GptMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GptUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

func init() {
	apiKey, ok := os.LookupEnv("API_KEY")
	if !ok {
		log.Fatal("lookup API_KEY env fail")
	}

	headers = map[string]string{}
	headers["Authorization"] = "Bearer " + apiKey

	ve.NewCustomValidator()
}
