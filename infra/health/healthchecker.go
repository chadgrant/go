package health

import (
	"context"
	"sync"
	"time"
)

type (
	// executor is a wrapper function type that wraps the provided Check function
	// with metadata such as when the test was executed, how long it took
	executor func() (duration time.Duration, testedAt time.Time, err error)

	// healthCheck is an internal state / wrapper mechanism
	healthCheck struct {
		Name     string
		Interval time.Duration
		Check    executor
	}

	// healthChecker is the main type that stores the health checks
	healthChecker struct {
		liveness  []*healthCheck
		readiness []*healthCheck
	}

	// wrapper to curry the exact error instance to the caller so that sentinel error checks work
	resultWrapper struct {
		Result Result
		Err    error
	}
)

// NewHealthChecker created a health check runner that satisfies the HealthChecker interface
func NewHealthChecker() HealthChecker {
	return &healthChecker{
		liveness:  make([]*healthCheck, 0),
		readiness: make([]*healthCheck, 0),
	}
}

// AddLiveness implementation of HealthChecker interface. See interface docs
func (h *healthChecker) AddLiveness(name string, check Check) {
	hc := &healthCheck{name, 0, wrapChecker(check)}
	h.liveness = append(h.liveness, hc)
}

func (h *healthChecker) AddLivenessBackground(name string, interval time.Duration, check Check) {
	h.AddLivenessBackgroundWithContext(context.Background(), name, interval, check)
}

func (h *healthChecker) AddLivenessBackgroundWithContext(ctx context.Context, name string, interval time.Duration, check Check) {
	hc := &healthCheck{name, interval, background(ctx, name, interval, wrapChecker(check))}
	h.liveness = append(h.liveness, hc)
}

func (h *healthChecker) AddReadiness(name string, check Check) {
	hc := &healthCheck{name, 0, wrapChecker(check)}
	h.readiness = append(h.readiness, hc)
}

func (h *healthChecker) AddReadinessBackground(name string, interval time.Duration, check Check) {
	h.AddReadinessBackgroundWithContext(context.Background(), name, interval, check)
}

func (h *healthChecker) AddReadinessBackgroundWithContext(ctx context.Context, name string, interval time.Duration, check Check) {
	hc := &healthCheck{name, interval, background(ctx, name, interval, wrapChecker(check))}
	h.readiness = append(h.readiness, hc)
}

func wrapChecker(check Check) executor {
	return func() (duration time.Duration, testedAt time.Time, err error) {
		testedAt = time.Now().UTC()
		start := time.Now()
		err = check()
		duration = time.Since(start)
		return
	}
}

// collect is a fan in pattern to collect the results of runChecks
func collect(checks ...[]*healthCheck) ([]Result, uint32, error) {
	var dur uint32
	var err error
	results := make([]Result, 0)

	for res := range runChecks(checks...) {
		dur += res.Result.Duration
		results = append(results, res.Result)
		if res.Err != nil {
			err = res.Err
		}
	}

	return results, dur, err
}

// runChecks is a fan out pattern to run tests concurrently
func runChecks(checks ...[]*healthCheck) <-chan resultWrapper {
	resultsC := make(chan resultWrapper, 10)

	var wg sync.WaitGroup

	for _, g := range checks {
		wg.Add(len(g))

		for _, c := range g {
			go func(c *healthCheck) {

				d, t, err := c.Check()

				r := resultWrapper{
					Result: Result{
						Name:     c.Name,
						Interval: uint32(c.Interval.Milliseconds()),
						Duration: uint32(d.Milliseconds()),
						TestedAt: &t,
					},
					Err: err,
				}

				if err != nil {
					r.Result.Status = err.Error()
				} else {
					r.Result.Status = success
				}

				resultsC <- r
				wg.Done()
			}(c)
		}
	}

	wg.Wait()
	close(resultsC)

	return resultsC
}
