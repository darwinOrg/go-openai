package dgpt

import (
	dgctx "github.com/darwinOrg/go-common/context"
	dghttp "github.com/darwinOrg/go-httpclient"
	dglogger "github.com/darwinOrg/go-logger"
	ve "github.com/darwinOrg/go-validator-ext"
	"log"
	"os"
)

var (
	headers map[string]string
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

func Execute[T any](ctx *dgctx.DgContext, apiUrl string, request any) (*T, error) {
	dglogger.Infof(ctx, "call chatgpt api url[%s], request: %+v", apiUrl, request)
	err := ve.ValidateDefault(request)
	if err != nil {
		return nil, err
	}

	response, err := dghttp.DoPostJsonToStruct[T](ctx, apiUrl, request, headers)
	if err != nil {
		dglogger.Errorf(ctx, "call chatgpt api url[%s] error: %v", apiUrl, err)
		return nil, err
	}

	return response, nil
}
