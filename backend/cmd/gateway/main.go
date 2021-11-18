package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/itohio/collective/backend/pkg/api"
	"github.com/itohio/collective/backend/pkg/auth"
	"github.com/itohio/collective/backend/pkg/db"
	"github.com/itohio/collective/backend/pkg/rpc"
	"github.com/itohio/collective/backend/pkg/spa"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

const defaultPort = "8080"

type arrayFlags []string

func (i *arrayFlags) String() string {
	return strings.Join(*i, ", ")
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

//
func main() {
	migrate := flag.Bool("migrate", true, "Migrate DB")
	flag.Parse()

	frontendPath := os.Getenv("FRONTEND_PATH")
	origins := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ";")

	if len(origins) == 0 {
		log.Fatal("At least one origin must be provided")
		return
	}

	db, err := db.New(*migrate)
	if err != nil {
		log.Fatal("Initializing database failed: ", err)
	}

	// if *migrate {
	// 	log.Println("Migrations done")
	// 	return
	// }

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	cors := cors.New(cors.Options{
		AllowedOrigins:   origins,
		AllowedMethods:   []string{"POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	router.Use(cors.Handler)

	router.Route("/", func(r chi.Router) {
		if frontendPath != "" {
			router.Mount("/", spa.New(frontendPath))
		}
		router.Mount("/api", api.New(db))
		router.Mount("/rpc", rpc.New(db))
		router.Mount("/auth", auth.New(db))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%s", port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println("Failed to listen: ", err)
			cancel()
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	log.Println("Running")
	select {
	case <-c:
		log.Println("Signal")
	case <-ctx.Done():
		log.Println("Failure")
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
