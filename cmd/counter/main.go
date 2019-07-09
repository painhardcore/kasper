package main

import (
	"github.com/painhardcore/kasper/pkg/server"
	"time"
)

func main() {
	server.Start("counter.storage", time.Minute)
}
