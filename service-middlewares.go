package main

import (
	"time"

	"github.com/go-kit/kit/log"
)

type stringServiceWithLogging struct {
	logger log.Logger
	next   StringService
}

func (mw stringServiceWithLogging) Count(s string) (n int) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "count",
			"input", s,
			"output", n,
			"took", time.Since(begin),
		)
	}(time.Now())
	n = mw.next.Count(s)
	return
}
