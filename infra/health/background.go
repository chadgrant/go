package health

import (
	"context"
	"errors"
	"time"
)

const (
	notRun  = "Not Run"
	success = "OK"
)

var errNotRun = errors.New(notRun)

// backgroundState a struct for currying parameters onto channel
type backgroundState struct {
	Duration time.Duration
	TestedAt time.Time
	Error    error
}

// background creates a ticker and loops forever checking the health check pre-emptive
// can be gracefully stopped by canceling the context
func background(ctx context.Context, name string, interval time.Duration, check executor) executor {
	// create a chan that will buffer the most recent check result
	resultsC := make(chan backgroundState, 1)

	// don't want to be ready/live until we've actually executed the check
	resultsC <- backgroundState{Duration: 0, Error: errNotRun}

	// func that runs the check and swaps out the current head of
	// the channel with the latest result
	executeCheck := func(check executor) {
		d, t, err := check()
		<-resultsC
		resultsC <- backgroundState{d, t, err}
	}

	go func(check executor) {
		// call once right away (time.Tick() doesn't always tick immediately
		// but we want an initial result as soon as possible)
		executeCheck(check)

		ticker := time.NewTicker(interval)
		for {
			select {
			case <-ticker.C:
				executeCheck(check)
			case <-ctx.Done():
				return
			}
		}
	}(check)

	return func() (duration time.Duration, testedAt time.Time, err error) {
		// peek at the head of the channel, then put it back
		res := <-resultsC
		resultsC <- res
		return res.Duration, res.TestedAt, res.Error
	}
}
