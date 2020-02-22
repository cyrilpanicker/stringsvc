package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httpTransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)
	var stringService StringService = stringServiceImplementation{}
	stringService = stringServiceWithLogging{logger, stringService}
	var countEndpoint = makeCountEndpoint(stringService)
	countEndpoint = addLogging(log.With(logger, "method", "count"))(countEndpoint)
	var countHandler = httpTransport.NewServer(countEndpoint, decodeCountRequest, encodeCountResponse)
	http.Handle("/count", countHandler)
	http.ListenAndServe(":8080", nil)
}
