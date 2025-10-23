package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/t-ogawa/hokkaido-nandoku-api/internal/handler"
	"github.com/t-ogawa/hokkaido-nandoku-api/internal/repository"
	"github.com/t-ogawa/hokkaido-nandoku-api/pkg/csvloader"
)

var httpAdapter *httpadapter.HandlerAdapter
var mux *http.ServeMux

func init() {
	log.Println("Initializing handler")
	csvPath := "data/nandoku_chimei.csv"

	// If running on Lambda, construct the absolute path to the data file
	if taskRoot := os.Getenv("LAMBDA_TASK_ROOT"); taskRoot != "" {
		csvPath = filepath.Join(taskRoot, "data/nandoku_chimei.csv")
	}

	placeNames, err := csvloader.LoadPlaceNames(csvPath)
	if err != nil {
		log.Fatalf("Failed to load place names from %s: %v", csvPath, err)
	}

	repo := repository.NewInMemoryPlaceNameRepository(placeNames)
	newHandler := handler.NewHandler(repo)

	mux = http.NewServeMux()
	mux.Handle("/", newHandler)

	httpAdapter = httpadapter.New(mux)
}

func main() {
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != "" {
		log.Println("Running on AWS Lambda")
		lambda.Start(Handler)
	} else {
		log.Println("Running on local")
		startLocalServer()
	}
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return httpAdapter.ProxyWithContext(ctx, req)
}

func startLocalServer() {
	// The handler is already initialized in init()
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
