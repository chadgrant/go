package health

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

var errIntentional = errors.New("intentional error")

type HealthCheckAdd func(string, time.Duration, Check)
type HealthCheckReporter func() ([]Result, error)

func TestHealthChecker(t *testing.T) {

	t.Run("TestLivenessCheckSucceeds", func(t *testing.T) {
		checker := NewHealthChecker()
		checkSucceeds(t, "liveness", checker.AddLiveness, checker.Liveness)
	})

	t.Run("TestReadinessCheckSucceeds", func(t *testing.T) {
		checker := NewHealthChecker()
		checkSucceeds(t, "readiness", checker.AddReadiness, checker.Readiness)
	})

	t.Run("TestLivenessCheckReportsFailure", func(t *testing.T) {
		checker := NewHealthChecker()
		testCheckReportsFailure(t, "readiness", checker.AddReadiness, checker.Readiness)
	})

	t.Run("TestReadinessCheckReportsFailure", func(t *testing.T) {
		checker := NewHealthChecker()
		testCheckReportsFailure(t, "readiness", checker.AddReadiness, checker.Readiness)
	})
}

func checkSucceeds(t *testing.T, name string, add HealthCheckAdd, report HealthCheckReporter) {

	checkRun := false

	add(fmt.Sprintf("%s-Test", name), 0, func() error {
		checkRun = true
		return nil
	})

	res, err := report()
	if err != nil {
		t.Errorf("could not get %s report %v", name, err)
	}

	if len(res) != 1 {
		t.Errorf("unexpected %s results count %d", name, len(res))
	}

	if !checkRun {
		t.Errorf("health check %s was not run", name)
	}
}

func testCheckReportsFailure(t *testing.T, name string, add HealthCheckAdd, report HealthCheckReporter) {

	checkRun := false

	add(fmt.Sprintf("%s-Test", name), 0, func() error {
		checkRun = true
		return errIntentional
	})

	res, err := report()
	if err == nil {
		t.Errorf("%s check did not return error", name)
	}

	if !errors.Is(err, errIntentional) {
		t.Errorf("%s unexpected error: %v", name, errIntentional)
	}

	if len(res) != 1 {
		t.Errorf("unexpected %s results count %d", name, len(res))
	}

	if !checkRun {
		t.Errorf("health check was not run for %s", name)
	}
}

func TestHealthReadinessAlsoRunsLivenessChecks(t *testing.T) {
	checker := NewHealthChecker()
	checker.AddReadiness("Test-Readiness", 0, func() error {
		return nil
	})
	checker.AddLiveness("Test-Liveness", 0, func() error {
		return nil
	})

	rpt, err := checker.Readiness()
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	if len(rpt) != 2 {
		t.Errorf("expected 2 tests run, got %d", len(rpt))
	}
}

func TestHealthOnlyRunsStaleChecks(t *testing.T) {
	testWasRun := 0

	checker := NewHealthChecker()
	checker.AddReadiness("Test-Readiness", time.Millisecond*50, func() error {
		testWasRun++
		return nil
	})

	_, err := checker.Readiness()
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	_, err = checker.Readiness()
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	if testWasRun > 1 {
		t.Errorf("test should not have been run twice")
	}

	time.Sleep(time.Millisecond * 100)
	_, err = checker.Readiness()
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	if testWasRun != 2 {
		t.Errorf("test should not have been run since it is considered stale")
	}
}

func TestHealthLivenessDoesNotRunReadinessChecks(t *testing.T) {
	livenessRun := false
	readinessRun := false

	checker := NewHealthChecker()
	checker.AddReadiness("Test-Readiness", 0, func() error {
		readinessRun = true
		return nil
	})
	checker.AddLiveness("Test-Liveness", 0, func() error {
		livenessRun = true
		return nil
	})

	_, err := checker.Liveness()
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	if readinessRun {
		t.Error("readiness check was run, it should not have been")
	}

	if !livenessRun {
		t.Error("liveness check not run")
	}
}

func TestHealthReport(t *testing.T) {
	livenessRun := false
	readinessRun := false

	checker := NewHealthChecker()
	checker.AddReadiness("Test-Readiness", 0, func() error {
		readinessRun = true
		time.Sleep(time.Millisecond * 10)
		return nil
	})
	checker.AddLiveness("Test-Liveness", 0, func() error {
		livenessRun = true
		time.Sleep(time.Millisecond * 10)
		return nil
	})

	rpt, err := checker.Report()
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	if len(rpt.Liveness) != 1 {
		t.Errorf("expected 1 liveness tests run, got %d", len(rpt.Liveness))
	}

	if len(rpt.Readiness) != 1 {
		t.Errorf("expected 1 readiness tests run, got %d", len(rpt.Readiness))
	}

	if rpt.ReportAsOf == nil {
		t.Error("report didnt return time it was executed")
	}

	if rpt.Duration <= 0 {
		t.Errorf("report did not return time taken to run tests %v", rpt.Duration)
	}

	if !readinessRun {
		t.Error("readiness check not run")
	}

	if !livenessRun {
		t.Error("liveness check not run")
	}
}

func TestHealthCheckWrapperReportsDuration(t *testing.T) {
	f := wrapChecker(func() error {
		time.Sleep(time.Millisecond * 10)
		return nil
	})

	d, _, _ := f()

	// give the test a little wiggle room

	x := d.Milliseconds()
	if x > 15 || x < 10 {
		t.Errorf("expected test to take %d ms, but got %d ms", 10, x)
	}
}

func TestHealthCheckWrapperReportsTimeTakenAndInUTC(t *testing.T) {
	f := wrapChecker(func() error {
		time.Sleep(time.Millisecond * 10)
		return nil
	})

	_, d, _ := f()

	x := time.Now().UTC().Sub(d).Milliseconds()
	if x > 15 || x < 10 {
		t.Errorf("expected test taken at %d ms, but got %d ms", 10, x)
	}
}

func BenchmarkHealthChecker(b *testing.B) {
	checker := NewHealthChecker()

	for i := 0; i < 1; i++ {
		checker.AddLiveness(fmt.Sprintf("live-%d", i), 0, func() error {
			return nil
		})

		checker.AddReadiness(fmt.Sprintf("readiness-%d", i), 0, func() error {
			return nil
		})
	}

	for i := 0; i < b.N; i++ {
		_, err := checker.Report()
		if err != nil {
			b.Error(err)
		}
	}
}
