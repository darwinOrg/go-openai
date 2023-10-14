package dgpt

import (
	dgctx "github.com/darwinOrg/go-common/context"
	dghttp "github.com/darwinOrg/go-httpclient"
	dglogger "github.com/darwinOrg/go-logger"
	ve "github.com/darwinOrg/go-validator-ext"
)

const chatApiUrl = "https://api.openai.com/v1/chat/completions"

type GptChatRequest struct {
	Model       string        `json:"model" binding:"required"`
	Messages    []*GptMessage `json:"messages" binding:"required,minLength=1"`
	Temperature float64       `json:"temperature"`
}

type GptChatResponse struct {
	Id      string           `json:"id"`
	Object  string           `json:"object"`
	Created int64            `json:"created"`
	Model   string           `json:"model"`
	Choices []*GptChatChoice `json:"choices"`
	Usage   *GptUsage        `json:"usage"`
}

type GptChatChoice struct {
	Index        int         `json:"index"`
	Message      *GptMessage `json:"message"`
	FinishReason string      `json:"finish_reason"`
}

func Chat(ctx *dgctx.DgContext, request *GptChatRequest) (*GptChatResponse, error) {
	dglogger.Infof(ctx, "GptChatRequest: %+v", request)
	err := ve.ValidateDefault(request)
	if err != nil {
		return nil, err
	}

	response, err := dghttp.DoPostJsonToStruct[GptChatResponse](ctx, chatApiUrl, request, headers)
	if err != nil {
		dglogger.Errorf(ctx, "gpt chat error: %v", err)
		return nil, err
	}

	return response, nil
}
