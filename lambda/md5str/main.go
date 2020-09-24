package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dgravesa/hello-lambda/pkg/md5hex"
)

// Request is a request
type Request struct {
	Str string `json:"str"`
}

// HandleRequest handles the request
func HandleRequest(ctx context.Context, r Request) (string, error) {
	return md5hex.Sum(r.Str), nil
}

func main() {
	lambda.Start(HandleRequest)
}
