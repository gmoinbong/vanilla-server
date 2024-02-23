package server

import (
	"log"
	"net/http"
	"vanilla-server/internal/config"
	"vanilla-server/internal/storage"
	"vanilla-server/router"
)


func RunServer(cfg *config.Config) error {

	db, err := storage.InitDB(cfg)
	if err != nil {
		log.Fatalf("Eror initializing DB: %v", err)
		return err
	}
	defer db.Close()
	router.SetupRoutes()

	if err = http.ListenAndServe(":8082", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
		return err
	}
	return nil
}
