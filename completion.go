package dgpt

import (
	dgctx "github.com/darwinOrg/go-common/context"
)

const completionApiUrl = "https://api.openai.com/v1/completions"

type GptCompletionRequest struct {
	Model       string  `json:"model" binding:"required"`
	Prompt      string  `json:"prompt" binding:"required"`
	MaxTokens   int     `json:"max_tokens" binding:"required"`
	Temperature float64 `json:"temperature"`
}

type GptCompletionResponse struct {
	Id      string                 `json:"id"`
	Object  string                 `json:"object"`
	Created int64                  `json:"created"`
	Model   string                 `json:"model"`
	Choices []*GptCompletionChoice `json:"choices"`
	Usage   *GptUsage              `json:"usage"`
}

type GptCompletionChoice struct {
	Text         string `json:"text"`
	Index        int    `json:"index"`
	Logprobs     string `json:"logprobs"`
	FinishReason string `json:"finish_reason"`
}

func Completion(ctx *dgctx.DgContext, request *GptCompletionRequest) (*GptCompletionResponse, error) {
	return Execute[GptCompletionResponse](ctx, completionApiUrl, request)
}
