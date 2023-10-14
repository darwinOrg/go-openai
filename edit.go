package dgpt

import (
	dgctx "github.com/darwinOrg/go-common/context"
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
	return Execute[GptEditResponse](ctx, editApiUrl, request)
}
