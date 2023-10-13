package dgpt

import (
	dgctx "github.com/darwinOrg/go-common/context"
	dghttp "github.com/darwinOrg/go-httpclient"
	dglogger "github.com/darwinOrg/go-logger"
	ve "github.com/darwinOrg/go-validator-ext"
)

type GptCompletionRequest struct {
	Model       string  `json:"model" binding:"required"`
	Prompt      string  `json:"prompt" binding:"required"`
	MaxTokens   int     `json:"max_tokens" binding:"required"`
	Temperature float64 `json:"temperature"`
}

type GptCompletionResponse struct {
	Id      string                  `json:"id"`
	Object  string                  `json:"object"`
	Created int64                   `json:"created"`
	Model   string                  `json:"model"`
	Choices []*GptCompletionChoices `json:"choices"`
	Usage   *GptUsage               `json:"usage"`
}

type GptCompletionChoices struct {
	Text         string `json:"text"`
	Index        int    `json:"index"`
	Logprobs     string `json:"logprobs"`
	FinishReason string `json:"finish_reason"`
}

func Completion(ctx *dgctx.DgContext, request *GptCompletionRequest) (*GptCompletionResponse, error) {
	dglogger.Infof(ctx, "GptCompletionRequest: %+v", request)
	err := ve.ValidateDefault(request)
	if err != nil {
		return nil, err
	}

	response, err := dghttp.DoPostJsonToStruct[GptCompletionResponse](ctx, COMPLETION_API_URL, request, headers)
	if err != nil {
		dglogger.Errorf(ctx, "gpt completion error: %v", err)
		return nil, err
	}

	return response, nil
}
