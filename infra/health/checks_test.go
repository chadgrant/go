package health

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type mockDB struct {
	ReturnError error
}

func (db *mockDB) Ping() error {
	return db.ReturnError
}

func (db *mockDB) PingContext(ctx context.Context) error {
	return db.ReturnError
}

func TestTCPDialCheck(t *testing.T) {
	if err := TCPDialCheck("google.com:80", 5*time.Second)(); err != nil {
		t.Errorf("could not dial port 80 to google: %v", err)
	}
	if err := TCPDialCheck("google.com:666", 5*time.Second)(); err == nil {
		t.Error("should not be able to connect to port 666 at google.com")
	}
}

func TestHTTPGetCheck(t *testing.T) {
	if err := HTTPGetCheck("https://golang.org/", 5*time.Second)(); err != nil {
		t.Errorf("should be able to do a http get request to golang.org: %v ", err)
	}

	if err := HTTPGetCheck("http://google.com", 5*time.Second)(); err == nil {
		t.Errorf("redirect should fail : %v", err)
	}

	if err := HTTPGetCheck("https://google.com/lol/wut", 5*time.Second)(); err == nil {
		t.Errorf("404 should fail: %v", err)
	}
}

func TestDatabasePingCheck(t *testing.T) {
	if err := DatabasePingCheck(nil, 1*time.Second)(); err == nil {
		t.Error("nil DB should fail")
	}

	failer := &mockDB{ReturnError: errors.New("not implemented")}
	if err := DatabasePingCheck(failer, 1*time.Second)(); err == nil {
		t.Error("ping should fail")
	}

	db := &mockDB{}
	if err := DatabasePingCheck(db, 1*time.Second)(); err != nil {
		t.Errorf("ping should succeed: %v", err)
	}

	rdb, err := sql.Open("sqlite3", ":memory:")

	if err != nil {
		t.Fatalf("could not open sql lite db: %v", err)
	}
	if err := DatabasePingCheck(rdb, 1*time.Second)(); err != nil {
		t.Errorf("ping should succeed: %v", err)
	}
}

func TestDNSResolveCheck(t *testing.T) {
	if err := DNSResolveCheck("google.com", 5*time.Second)(); err != nil {
		t.Errorf("could not resolve dns for google.com: %v", err)
	}
	if err := DNSResolveCheck("lol.watdidthefoxsay.com", 5*time.Second)(); err == nil {
		t.Error("dns for non existant domain should fail")
	}
}

func TestGoroutineCountCheck(t *testing.T) {
	if err := GoroutineCountCheck(1000)(); err != nil {
		t.Errorf("go routine check should not fail: %v", err)
	}
	if err := GoroutineCountCheck(0)(); err == nil {
		t.Error("threshold of zero goroutines should fail")
	}
}
