# Health
Package for service health checks

# Liveness (GTG) Good to Go - /live
This package has no opinions, but a service with no liveness checks is aprropriate for most services. Liveness meaning that it is able to serve http and is alive/running.

GET /live endpoint executes all the liveness tests, returns no body. Returns 200 status code if all test succeed. 

The "Good To Go" (GTG) returns a successful response in the case that the service is in an operational state and is able to receive traffic. This resource is used by load balancers and monitoring tools to determine if traffic should be routed to this service or not.

Note that GTG is not used to determine if the service is healthy or not, only if it is able to receive traffic. A healthy instance may not be able to accept traffic due to the failure of critical downstream dependencies.

A failed response is a 503. Failure to respond within a predetermined timeout typically 2 seconds should also treated as a failure.

# Readiness (Service Canary) - /ready
Readieness is meant for checking all service dependencies. Dependencies that are considered required for the service to be considerred healthy

The "Service Canary" (ASG) returns a successful response in the case that the service is in a healthy state. If a service returns a failure response or fails to respond within a predefined timeout then the service can expect to be terminated and replaced. (Typically this resouce is used in auto-scaling group healthchecks.)

A successful response is a 200 OK. A failed response is a 503. Failure to respond within a predetermined timeout typically 2 seconds should be treated as a failure.

GET /ready endpoint executes all the readiness AND liveness tests, returns no body. Returns 200 status code if all checks succeed. 

# Report - /health

Used for reporting / human readable & automated tooling. Returns 200 status code if all checks succeed. 

> GET /health

| Property                  | Description                                                                       |     Example               |
| --------------------------|-----------------------------------------------------------------------------------|---------------------------|
| report_as_of_utc          | The time at which this report was generated (this may not be the current time)    | 2015-03-12T19:40:18.877Z  |
| duration_ms               | Sum of all check durations                                                        | 10                        |
| liveness                  | array of liveness healthcheck test reports                                        |                           |
| liveness[].name           | name of the healthcheck test                                                      | sql                       |
| liveness[].duration_ms    | Number of milliseconds taken to run the test                                      | 100                       |
| liveness[].interval_ms    | if background test, how often the test is to be run                               | 100                       |
| liveness[].status         | The state of the test, if not "OK", is in failed/error state                      | OK                        |
| liveness[].tested_at_utc  | The last time the test was run                                                    | 2015-03-12T19:40:18.877Z  |
| readiness                 | array of readiness healthcheck test reports                                       |                           |
| readiness[].name          | name of the healthcheck test                                                      | sql                       |
| readiness[].duration_ms   | Number of milliseconds taken to run the test                                      | 100                       |
| readiness[].interval_ms   | if background test, how often the test is to be run                               | 100                       |
| readiness[].status        | The state of the test, if not "OK", is in failed/error state                      | OK                        |
| readiness[].tested_at_utc | The last time the test was run                                                    | 2015-03-12T19:40:18.877Z  |

# Basic Usage

## Note!
Be warned, you will essentially be ddos'ing things if you choose short staleness times and this is deployed across all microservices (for example: if they all do a http chek every 10 seconds to the same host * how many services are running across all environments) ... the host being used as the health check could be annoyed. 

```go
    checker := health.NewHealthChecker()
    
	checker.AddLiveness("Liveness Test", time.Second*10, func() error {
        //run test
		return nil
    })
    
    checker.AddReadiness("Readiness Test", time.Second*10, func() error {
        // run test	
		return nil
	})
```

# Health Check Timeouts

You should define a timeout for checks

```go
    checker := health.NewHealthChecker()
    
	checker.AddLiveness("Liveness Test",  time.Second*10, func() error {
        ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
        defer cancel()
        //run test
		return nil
    })
```

# Background checks

For the test to run continuosly, register a background check with a duration of how often it should be run. If the test takes longer than the duration, it will be run immedieatly after the previous check.

```go
    checker := health.NewHealthChecker()
    
	checker.AddLivenessBackground("Liveness Test", time.Second * 10, func() error {
        //run test
		return nil
    })
    
    checker.AddReadinessBackground("Readiness Test", time.Second * 1, func() error {
        // run test	
		return nil
	})
```

# Cancel / Gracefully stop background checks

```go
    ctx, cancel := context.Backgound()
    defer cancel()

    checker := health.NewHealthChecker()
    
	checker.AddLivenessBackgroundWithContext(ctx, "Liveness Test", time.Second * 10, func() error {
        //run test
		return nil
    })
    
    checker.AddReadinessBackgroundWithContext(ctx, "Readiness Test", time.Second * 1, func() error {
        // run test	
		return nil
	})
```

# Integration with http handlers

```go
    checker := health.NewHealthChecker()

    h := health.NewHandler(checker)
    
	r.HandleFunc("/live", h.Live)
	r.HandleFunc("/ready", h.Ready)
	r.HandleFunc("/health", h.Report)
```

# Included Checks

## TCPDialCheck

```go
checker := health.NewHealthChecker()

checker.AddReadiness("google-http-connection",  time.Second*10, health.TCPDialCheck("google.com:80", 5*time.Second))
```

## DatabasePingCheck
```go
checker := health.NewHealthChecker()

checker.AddReadiness("sql-db-ping", time.Second*10, health.DatabasePingCheck(db, 1*time.Second))
```

## DNSResolveCheck
```go
checker := health.NewHealthChecker()

checker.AddReadiness("google-dns-resolution", time.Second*10, health.DNSResolveCheck("google.com", 5*time.Second))
```

## GoroutineCountCheck
```go
checker := health.NewHealthChecker()

checker.AddReadiness("max-goroutine", time.Second*10, health.GoroutineCountCheck(1000))
```


# Docker Health Check Sample (Docker-Compose)

```
    healthcheck:
      test: "curl -f http://localhost:8080/health/live || exit 1"
      interval: 30s
      timeout: 2s
      retries: 3
      start_period: 5s
```