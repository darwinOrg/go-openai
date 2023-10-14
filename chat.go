package dgpt

import (
	"context"
	dgctx "github.com/darwinOrg/go-common/context"
	dglogger "github.com/darwinOrg/go-logger"
	"github.com/sashabaranov/go-openai"
)

func Chat(ctx *dgctx.DgContext, request openai.ChatCompletionRequest) (openai.ChatCompletionResponse, error) {
	response, err := client.CreateChatCompletion(context.Background(), request)
	dglogger.Infof(ctx, "create chat completion, request: %+v, response: %+v, error: %v", request, response, err)
	return response, err
}
