package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hi Rss feed")

	godotenv.Load(".env")
	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT is not found in env")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:  []string{"https://*", "http://*"},
		AllowedMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:  []string{"*"},
		ExposedHeaders:  []string{"Link"},
		AllowCredetials: false,
		MaxAge:          300,
	}))

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
