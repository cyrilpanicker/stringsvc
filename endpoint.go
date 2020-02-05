package main

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V int `json:"v"`
}

func makeCountEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return countResponse{svc.Count(request.(countRequest).S)}, nil
	}
}
