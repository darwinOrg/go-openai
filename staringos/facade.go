package staringos

import (
	dgctx "github.com/darwinOrg/go-common/context"
	dghttp "github.com/darwinOrg/go-httpclient"
)

var httpClient *dghttp.DgHttpClient

func init() {
	httpClient = dghttp.NewHttpClient(true)
}

// CreateChatCompletion — API call to Create a completion for the chat message.
func CreateChatCompletion(
	ctx *dgctx.DgContext,
	request *ChatCompletionRequest,
	apiKey string,
) (string, error) {
	headers := map[string]string{}
	headers["Authorization"] = "Bearer " + apiKey
	//headers["Accept-Encoding"] = "identity"

	body, err := httpClient.DoPostJson(ctx, chatCompletionsUrl, request, headers)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
