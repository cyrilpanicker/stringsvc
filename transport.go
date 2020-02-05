package main

import (
	"context"
	"encoding/json"
	"net/http"
)

func decodeCountRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	var countRequest countRequest
	if err := json.NewDecoder(request.Body).Decode(&countRequest); err != nil {
		return nil, err
	}
	return countRequest, nil
}

func encodeCountResponse(ctx context.Context, writer http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(writer).Encode(response)
}
