package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)
	stringService := stringService{}
	http.Handle("/count", httptransport.NewServer(
		endpointLoggingMiddleware(log.With(logger, "method", "count"))(
			makeCountEndpoint(
				stringServiceWithLogging{
					logger, stringService,
				},
			),
		),
		decodeCountRequest,
		encodeCountResponse,
	))
	http.ListenAndServe(":8080", nil)
}
