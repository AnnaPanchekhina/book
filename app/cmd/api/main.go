package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

type application struct {
	config config
	logger *log.Logger
	db     *sql.DB
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://user:password@db:5432/itbookworm?sslmode=disable", "PostgreSQL DSN")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Подключение к PostgreSQL с ретраями
	var db *sql.DB
	var err error

	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres", cfg.db.dsn)
		if err == nil {
			err = db.Ping()
			if err == nil {
				break
			}
		}
		logger.Printf("Waiting for database... (attempt %d/10)", i+1)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		logger.Fatal("Failed to connect to database after retries:", err)
	}
	defer db.Close()

	logger.Println("Database connection pool established")

	app := &application{
		config: cfg,
		logger: logger,
		db:     db,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		logger.Println("Shutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
		}

		logger.Println("Completing background tasks...")
		shutdownError <- nil
	}()

	logger.Printf("Starting %s server on %s", cfg.env, srv.Addr)

	err = srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Fatal(err)
	}

	err = <-shutdownError
	if err != nil {
		logger.Fatal(err)
	}

	logger.Println("Server stopped")
}