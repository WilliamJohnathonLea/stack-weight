package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	sw "github.com/WilliamJohnathonLea/stack-weight/stackweight"
)

var emptyRes = sw.Response{}

// RequestHandler The request handler
func RequestHandler(req sw.Request) (sw.Response, error) {
	if err := req.Validate(); err != nil {
		return emptyRes, err
	}
	out := sw.Response{ApproxWeight: req.CalcApproxOutputWeight()}
	return out, nil
}

func main() {
	lambda.Start(RequestHandler)
}
