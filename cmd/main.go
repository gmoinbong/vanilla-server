package main

import (
	"log"
	"vanilla-server/cmd/server"
	"vanilla-server/internal/config"
	"vanilla-server/utils/lockutil"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.MustLoadConfig()

	lockutil.RunWithLock(func() {
		if err := server.RunServer(cfg); err != nil {
			log.Fatalf("Error running server: %v", err)
		}
	})
}
