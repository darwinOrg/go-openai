package dgpt

import (
	dgctx "github.com/darwinOrg/go-common/context"
	dghttp "github.com/darwinOrg/go-httpclient"
	dglogger "github.com/darwinOrg/go-logger"
	ve "github.com/darwinOrg/go-validator-ext"
)

const editApiUrl = "https://api.openai.com/v1/edits"

type GptEditRequest struct {
	Model       string `json:"model" binding:"required"`
	Input       string `json:"input" binding:"required"`
	Instruction string `json:"instruction" binding:"required"`
}

type GptEditResponse struct {
	Object  string           `json:"object"`
	Created int64            `json:"created"`
	Choices []*GptEditChoice `json:"choices"`
	Usage   *GptUsage        `json:"usage"`
}

type GptEditChoice struct {
	Text  string `json:"text"`
	Index int    `json:"index"`
}

func Edit(ctx *dgctx.DgContext, request *GptEditRequest) (*GptEditResponse, error) {
	dglogger.Infof(ctx, "GptEditRequest: %+v", request)
	err := ve.ValidateDefault(request)
	if err != nil {
		return nil, err
	}

	response, err := dghttp.DoPostJsonToStruct[GptEditResponse](ctx, editApiUrl, request, headers)
	if err != nil {
		dglogger.Errorf(ctx, "gpt chat error: %v", err)
		return nil, err
	}

	return response, nil
}
