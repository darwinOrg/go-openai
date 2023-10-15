package openai

import (
	"context"
	"github.com/darwinOrg/go-common/constants"
	dgctx "github.com/darwinOrg/go-common/context"
	dglogger "github.com/darwinOrg/go-logger"
	"os"
)

var client *Client

func init() {
	apiKey := os.Getenv("OPENAI_API_KEY")

	client = NewClient(apiKey)
}

func Completion(ctx *dgctx.DgContext, request CompletionRequest) (CompletionResponse, error) {
	reqCtx := context.WithValue(context.Background(), constants.TraceId, ctx.TraceId)
	response, err := client.CreateCompletion(reqCtx, request)
	dglogger.Infof(ctx, "create completion, request: %+v, response: %+v, error: %v", request, response, err)
	return response, err
}

func Chat(ctx *dgctx.DgContext, request ChatCompletionRequest) (ChatCompletionResponse, error) {
	reqCtx := context.WithValue(context.Background(), constants.TraceId, ctx.TraceId)
	response, err := client.CreateChatCompletion(reqCtx, request)
	dglogger.Infof(ctx, "create chat completion, request: %+v, response: %+v, error: %v", request, response, err)
	return response, err
}
