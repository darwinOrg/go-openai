package staringos

import (
	dgctx "github.com/darwinOrg/go-common/context"
	dgerr "github.com/darwinOrg/go-common/enums/error"
	dghttp "github.com/darwinOrg/go-httpclient"
	"time"
)

var httpClient *dghttp.DgHttpClient

func init() {
	httpClient = dghttp.NewHttpClient(true)
}

// CreateChatCompletion — API call to Create a completion for the chat message.
func CreateChatCompletion(
	ctx *dgctx.DgContext,
	request *ChatByCorpusRequest,
	apiKey string,
) (*MessageResponse, error) {
	headers := map[string]string{}
	headers["Authorization"] = "Bearer " + apiKey

	messageIdBody, err := httpClient.DoPostJson(ctx, chatByCorpusUrl, request, headers)
	if err != nil {
		return nil, err
	}

	messageId, err := dghttp.ConvertJsonToStruct[MessageIdResponse](messageIdBody)
	if err != nil {
		return nil, err
	}

	messageUrl := messageUrlPrefix + messageId.Id
	time.Sleep(3 * time.Second)

	for i := 0; i < 100; i++ {
		time.Sleep(2 * time.Second)

		messageBody, err := httpClient.DoGet(ctx, messageUrl, map[string]string{}, headers)
		if err != nil {
			return nil, err
		}

		message, err := dghttp.ConvertJsonToStruct[MessageResponse](messageBody)
		if err != nil {
			return nil, err
		}

		if message.IsFinish {
			return message, nil
		}
	}

	return nil, dgerr.TIME_OUT
}
