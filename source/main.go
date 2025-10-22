package main

import (
	"log"
	"net/http"
	"os"

	"github.com/t-ogawa/hokkaido-nandoku-api/internal/handler"
	"github.com/t-ogawa/hokkaido-nandoku-api/internal/repository"
	"github.com/t-ogawa/hokkaido-nandoku-api/pkg/csvloader"
)

func main() {
	// For AWS Lambda, the handler is executed directly.
	// For local execution, we start a simple HTTP server.
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != "" {
		// Lambda execution is handled by the runtime, so we just need to register the handler.
		// This part will be completed in a later task.
		log.Println("Running on AWS Lambda")
	} else {
		log.Println("Running on local")
		startLocalServer()
	}
}

func startLocalServer() {
	placeNames, err := csvloader.LoadPlaceNames("data/nandoku_chimei.csv")
	if err != nil {
		log.Fatalf("Failed to load place names: %v", err)
	}

	repo := repository.NewInMemoryPlaceNameRepository(placeNames)
	h := handler.NewRandomPlaceNameHandler(repo)

	http.Handle("/random", h)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
