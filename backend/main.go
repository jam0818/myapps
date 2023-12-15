package main

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"db/db"
	"db/handler/router"
)

func main() {
	const (
		defaultPort   = ":8080"
		defaultDBPath = "/mnt/c/Users/nkawa/Desktop/DB/sentence.sqlite"
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = defaultDBPath
	}
	// set time zone
	var err error
	time.Local, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
	}

	// set up sqlite3
	sentimentTextDB, err := db.NewDB(dbPath)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connect DB successfully")
	}
	defer func(todoDB *sql.DB) {
		err := todoDB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(sentimentTextDB)
	// NOTE: 新しいエンドポイントの登録はrouter.NewRouterの内部で行うようにする
	mux := router.NewRouter(sentimentTextDB)

	// TODO: サーバーをlistenする
	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}
	log.Println("Start listening server")
	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			// Error starting or closing listener:
			log.Fatalln("Server closed with error:", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt, os.Kill)
	log.Printf("SIGNAL %d received, then shutting down...\n", <-quit)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		// Error from closing listeners, or context timeout:
		log.Println("Failed to gracefully shutdown:", err)
	}
	log.Println("Server shutdown")
}
