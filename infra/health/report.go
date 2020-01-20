package health

import (
	"os"
	"runtime"
	"time"
)

var (
	hostname, _ = os.Hostname()
	startTime   = time.Now().UTC()
)

func (h *healthChecker) Report() (Report, error) {
	t := time.Now().UTC()

	l, ld, lerr := collect(h.Live)
	r, rd, rerr := collect(h.Ready)

	rpt := Report{
		ReportAsOf:      &t,
		Duration:        ld + rd,
		Liveness:        l,
		Readiness:       r,
		HostName:        hostname,
		UpSince:         startTime,
		OsArch:          runtime.GOARCH,
		OsName:          runtime.GOOS,
		OsGoRoutines:    runtime.NumGoroutine(),
		OsNumProcessors: runtime.NumCPU(),
		GoVersion:       runtime.Version(),
		Version:         1,
		Memory:          memory(),
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

func memory() Memory {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return Memory{
		Alloc:         m.Alloc,
		TotalAlloc:    m.TotalAlloc,
		Sys:           m.Sys,
		Lookups:       m.Lookups,
		Mallocs:       m.Mallocs,
		Frees:         m.Frees,
		HeapAlloc:     m.HeapAlloc,
		HeapSys:       m.HeapSys,
		HeapIdle:      m.HeapIdle,
		HeapInuse:     m.HeapInuse,
		HeapReleased:  m.HeapReleased,
		HeapObjects:   m.HeapObjects,
		StackInuse:    m.StackInuse,
		StackSys:      m.StackSys,
		MSpanInuse:    m.MSpanInuse,
		MSpanSys:      m.MSpanSys,
		MCacheInuse:   m.MCacheInuse,
		MCacheSys:     m.MCacheSys,
		BuckHashSys:   m.BuckHashSys,
		GCSys:         m.GCSys,
		OtherSys:      m.OtherSys,
		NextGC:        m.NextGC,
		LastGC:        m.LastGC,
		PauseTotalNs:  m.PauseTotalNs,
		NumGC:         m.NumGC,
		NumForcedGC:   m.NumForcedGC,
		GCCPUFraction: m.GCCPUFraction,
		EnableGC:      m.EnableGC,
		DebugGC:       m.DebugGC,
	}
}
