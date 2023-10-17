package openai

import (
	"context"
	dgctx "github.com/darwinOrg/go-common/context"
	dglogger "github.com/darwinOrg/go-logger"
	"time"
)

func CreateCompletion(ctx *dgctx.DgContext, request CompletionRequest, apiKey string) (CompletionResponse, error) {
	start := time.Now().UnixMilli()
	client := NewClient(apiKey)
	response, err := client.CreateCompletion(context.Background(), request)
	dglogger.Infof(ctx, "create completion, request: %+v, response: %+v, error: %v, cost: %d ms",
		request, response, err, time.Now().UnixMilli()-start)
	return response, err
}

func CreateChatCompletion(ctx *dgctx.DgContext, request ChatCompletionRequest, apiKey string) (ChatCompletionResponse, error) {
	start := time.Now().UnixMilli()
	client := NewClient(apiKey)
	response, err := client.CreateChatCompletion(context.Background(), request)
	dglogger.Infof(ctx, "create chat completion, request: %+v, response: %+v, error: %v, cost: %d ms",
		request, response, err, time.Now().UnixMilli()-start)
	return response, err
}
