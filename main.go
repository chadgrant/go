package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/chadgrant/go-http-infra/infra"
	"github.com/chadgrant/go-http-infra/infra/health"
	"github.com/chadgrant/go-http-infra/infra/schema"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	host := *flag.String("host", infra.GetEnvVar("SVC_HOST", "localhost"), "default binding 0.0.0.0")
	port := *flag.Int("port", infra.GetEnvVarInt("SVC_PORT", 8080), "default port 8080")
	r := mux.NewRouter()

	gorillaW := func(s string, w http.HandlerFunc) {
		r.HandleFunc(s, w)
	}

	checker := health.NewHealthChecker()
	schemas := schema.NewRegistry()
	if err := infra.RegisterInfraHandlers(gorillaW, checker, schemas); err != nil {
		panic(err)
	}
	if err := schemas.AddDirectory("./infra/schema/test-schemas"); err != nil {
		panic(err)
	}

	checker.AddReadiness("max-goroutine", time.Millisecond*500, health.GoroutineCountCheck(1000))
	checker.AddReadiness("google-http-connection", time.Second*10, health.TCPDialCheck("google.com:80", 5*time.Second))
	checker.AddReadiness("google-dns-resolution", time.Second*10, health.DNSResolveCheck("google.com", 5*time.Second))

	checker.AddLivenessBackground("Liveness-Test-That-Sleeps", time.Second*3, func() error {
		log.Println("running liveness background test")
		time.Sleep(time.Millisecond * 50)
		return nil
	})

	checker.AddLiveness("Liveness-Lazy-Test-That-Sleeps-A-Little", time.Millisecond*150, func() error {
		//log.Println("running lazy liveness test")
		time.Sleep(time.Millisecond * 50)
		return nil
	})

	checker.AddReadinessBackground("Readiness-Test-That-Sleeps-1-Second", time.Second*3, func() error {
		log.Println("running readiness background test")
		time.Sleep(time.Second * 1)
		return nil
	})

	checker.AddReadiness("Readiness-Lazy-Test-That-Sleeps-A-Little", time.Millisecond*150, func() error {
		//log.Println("running lazy readiness test")
		time.Sleep(time.Millisecond * 50)
		return nil
	})

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		ExposedHeaders:   []string{"Location"},
		MaxAge:           86400,
	})

	c.Handler(r)

	log.Printf("Started, serving at %s:%d\n", host, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), r))
}
