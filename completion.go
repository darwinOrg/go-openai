package dgpt

import (
	"context"
	dgctx "github.com/darwinOrg/go-common/context"
	dglogger "github.com/darwinOrg/go-logger"
	"github.com/sashabaranov/go-openai"
)

func Completion(ctx *dgctx.DgContext, request openai.CompletionRequest) (openai.CompletionResponse, error) {
	response, err := client.CreateCompletion(context.Background(), request)
	dglogger.Infof(ctx, "create completion, request: %+v, response: %+v, error: %v", request, response, err)
	return response, err
}
