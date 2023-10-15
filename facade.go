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
	response, err := client.CreateCompletion(buildContextWithTraceId(ctx), request)
	dglogger.Infof(ctx, "create completion, request: %+v, response: %+v, error: %v", request, response, err)
	return response, err
}

func Chat(ctx *dgctx.DgContext, request ChatCompletionRequest) (ChatCompletionResponse, error) {
	response, err := client.CreateChatCompletion(buildContextWithTraceId(ctx), request)
	dglogger.Infof(ctx, "create chat, request: %+v, response: %+v, error: %v", request, response, err)
	return response, err
}

func buildContextWithTraceId(ctx *dgctx.DgContext) context.Context {
	return context.WithValue(context.Background(), constants.TraceId, ctx.TraceId)
}
