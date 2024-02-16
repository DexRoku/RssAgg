package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/rs/cors" // Import cors package
)

func main() {
	fmt.Println("Hi Rss feed")

	godotenv.Load(".env")
	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT is not found in env")
	}

	router := chi.NewRouter()

	// Here we have cors handler have a definiton of what the server accepts
	// Like https http, GET, POST, PUT, DELETE, OPTIONS etc
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	})

	router.Use(corsHandler.Handler)

	v1Router := chi.NewRouter()
	v1Router.HandleFunc("/healthz", handlerReadiness)

	router.Mount("/v1", v1Router)

	httpServer := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("server starting on port %v", portString)
	err := httpServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
