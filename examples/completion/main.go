package main

import (
	"context"
	"fmt"
	"github.com/darwinOrg/go-openai"
)

func main() {
	resp, err := openai.DefaultClient.CreateCompletion(
		context.Background(),
		openai.CompletionRequest{
			Model:     openai.GPT3Ada,
			MaxTokens: 5,
			Prompt:    "Lorem ipsum",
		},
	)
	if err != nil {
		fmt.Printf("Completion error: %v\n", err)
		return
	}
	fmt.Println(resp.Choices[0].Text)
}
