package main

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

func endpointLoggingMiddleware(logger log.Logger) func(endpoint.Endpoint) endpoint.Endpoint {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			logger.Log("msg", "calling endpoint")
			defer func(begin time.Time) {
				logger.Log("msg", "called endpoint", "took", time.Since(begin))
			}(time.Now())
			return next(ctx, request)
		}
	}
}
