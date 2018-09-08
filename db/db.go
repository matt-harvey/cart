package db

import (
	"os"
	signal "os/signal"
	"sync"
	"syscall"

	"github.com/gobuffalo/pop"
)

var (
	db               *pop.Connection
	initializeDbOnce sync.Once
)

// TODO Ensure database is closed no matter where it's first initialized from.
func initializeDb() {
	var err error
	db, err = pop.Connect(os.Getenv("ENV"))
	if err != nil {
		panic(err)
	}

	// cleanup
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)
	signal.Notify(done, syscall.SIGTERM)
	signal.Notify(done, syscall.SIGKILL)
	go func() {
		sig := <-done
		db.Close()
		panic(sig)
	}()
}

func Conn() *pop.Connection {
	initializeDbOnce.Do(initializeDb)
	return db
}
