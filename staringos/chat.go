package staringos

import (
	"github.com/darwinOrg/go-openai"
)

const chatCompletionsUrl = "http://api.ai.staringos.com/"

// ChatCompletionRequest represents a request structure for chat completion API.
type ChatCompletionRequest struct {
	Model        string                         `json:"model,omitempty"`
	Prompt       []openai.ChatCompletionMessage `json:"prompt,omitempty"`
	Temperature  float32                        `json:"temperature,omitempty"`
	TopP         float32                        `json:"top_p,omitempty"`
	Functions    []openai.FunctionDefinition    `json:"functions,omitempty"`
	FunctionCall any                            `json:"function_call,omitempty"`
}

// ChatCompletionResponse represents a response structure for chat completion API.
type ChatCompletionResponse struct {
	Content      string `json:"content,omitempty"`
	IsFinish     bool   `json:"isFinish,omitempty"`
	FunctionCall any    `json:"function_call,omitempty"`
	Model        string `json:"model,omitempty"`
	AppId        int    `json:"appId,omitempty"`
	At           int64  `json:"at,omitempty"`
}
