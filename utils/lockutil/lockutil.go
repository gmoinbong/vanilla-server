package lockutil

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

const lockFileName = "server.lock"

func Check() bool {
	_, err := os.Stat(lockFileName)
	return err == nil
}

func Create() error {
	f, err := os.Create(lockFileName)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}

func Remove() error {
	return os.Remove(lockFileName)
}

func RunWithLock(f func()) {
	if Check() {
		fmt.Println("Server is already running")
		return
	}
	err := Create()
	if err != nil {
		fmt.Println("Error creating lock file:", err)
		return
	}
	defer Remove()

	sig := make(chan os.Signal, 1)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sig
		Remove()
		os.Exit(1)
	}()

	f()
}
