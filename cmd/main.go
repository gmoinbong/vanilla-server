package main

import (
	_ "github.com/lib/pq"
	"vanilla-server/cmd/server"
	"vanilla-server/utils/lockutil"
)

func main() {
	lockutil.RunWithLock(server.RunServer)
}
