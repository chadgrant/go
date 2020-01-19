package health

import (
	"time"
)

func (h *healthChecker) Report() (Report, error) {
	t := time.Now().UTC()

	l, ld, lerr := collect(h.Live)
	r, rd, rerr := collect(h.Ready)

	rpt := Report{
		ReportAsOf: &t,
		Duration:   ld + rd,
		Liveness:   l,
		Readiness:  r,
	}

	if lerr != nil {
		return rpt, lerr
	}

	return rpt, rerr
}

func (h *healthChecker) Liveness() (results []Result, err error) {
	results, _, err = collect(h.Live)
	return
}

func (h *healthChecker) Readiness() (results []Result, err error) {
	results, _, err = collect(h.Live, h.Ready)
	return
}
