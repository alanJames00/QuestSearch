package main

import (
	"log"
	"net"
	"questsearch/configs"
	"questsearch/internal/formatters"
	"questsearch/internal/handlers"
	"questsearch/internal/repositories"
	"questsearch/internal/services"
	"questsearch/pkg/database"
	"questsearch/proto/pb/question"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// load the config
	cfg := configs.LoadConfig()

	// connect to mongodb
	client, err := database.ConnectToMongo(cfg.MongoURI)
	if err != nil {
		log.Fatal("unable to connect to mongodb", err)
	}

	// select the db with dbname
	db := client.Database(cfg.DBName)

	// setup the repositories with collection name
	repo := repositories.NewMongoQuestionRepository(db, "questions")

	// init the Generic question service
	svc := services.NewQuestionService(repo)

	// init the grpc formatter service
	grpcFormatterService := formatters.NewGrpcFormatterService(svc)

	// init the grpc handler
	grpcQuestionHandler := handlers.NewGrpcHandler(grpcFormatterService)

	// create a grpc server and register the grpc handler for questionService
	grpcServer := grpc.NewServer()
	question.RegisterQuestionServiceServer(grpcServer, grpcQuestionHandler)

	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", cfg.Addr)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("grpc server starting on Addr: %s\n", cfg.Addr)
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatal("failed to serve grpc server", err)
	}

}
