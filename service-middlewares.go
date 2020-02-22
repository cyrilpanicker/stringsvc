package main

import (
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
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

type stringServiceWithInstrumenting struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	next           StringService
}

func (mw stringServiceWithInstrumenting) Count(s string) (n int) {
	defer func(begin time.Time) {
		values := []string{"method", "count"}
		mw.requestCount.With(values...).Add(1)
		mw.requestLatency.With(values...).Observe(time.Since(begin).Seconds())
		mw.countResult.Observe(float64(n))
	}(time.Now())
	n = mw.next.Count(s)
	return
}
