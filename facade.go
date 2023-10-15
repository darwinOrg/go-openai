package openai

import (
	"context"
	dgctx "github.com/darwinOrg/go-common/context"
	dglogger "github.com/darwinOrg/go-logger"
	"os"
	"time"
)

var DefaultClient *Client

func init() {
	apiKey := os.Getenv("OPENAI_API_KEY")

	DefaultClient = NewClient(apiKey)
}

func Completion(ctx *dgctx.DgContext, request CompletionRequest) (CompletionResponse, error) {
	start := time.Now().UnixMilli()
	response, err := DefaultClient.CreateCompletion(context.Background(), request)
	dglogger.Infof(ctx, "create completion, request: %+v, response: %+v, error: %v, cost: %d ms",
		request, response, err, time.Now().UnixMilli()-start)
	return response, err
}

func Chat(ctx *dgctx.DgContext, request ChatCompletionRequest) (ChatCompletionResponse, error) {
	start := time.Now().UnixMilli()
	response, err := DefaultClient.CreateChatCompletion(context.Background(), request)
	dglogger.Infof(ctx, "create chat, request: %+v, response: %+v, error: %v, cost: %d ms",
		request, response, err, time.Now().UnixMilli()-start)
	return response, err
}
