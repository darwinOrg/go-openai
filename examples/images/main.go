package main

import (
	"context"
	"fmt"
	"github.com/darwinOrg/go-openai"
)

func main() {
	respUrl, err := openai.DefaultClient.CreateImage(
		context.Background(),
		openai.ImageRequest{
			Prompt:         "Parrot on a skateboard performs a trick, cartoon style, natural light, high detail",
			Size:           openai.CreateImageSize256x256,
			ResponseFormat: openai.CreateImageResponseFormatURL,
			N:              1,
		},
	)
	if err != nil {
		fmt.Printf("Image creation error: %v\n", err)
		return
	}
	fmt.Println(respUrl.Data[0].URL)
}
