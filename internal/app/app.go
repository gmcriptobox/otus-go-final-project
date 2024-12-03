package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gmcriptobox/otus-go-final-project/internal/config"
	"github.com/gmcriptobox/otus-go-final-project/internal/controller/httpapi"
	"github.com/gmcriptobox/otus-go-final-project/internal/controller/httpapi/handler"
	"github.com/gmcriptobox/otus-go-final-project/internal/repository"
	"github.com/gmcriptobox/otus-go-final-project/internal/repository/client"
	"github.com/gmcriptobox/otus-go-final-project/internal/service"
	_ "github.com/lib/pq"
)

func main() {
	projectConfig, err := config.Read("./configs/config.yaml")
	if err != nil {
		fmt.Println("error while reading config file: ", err)
		exitAfterDefer(1)
	}

	psql := client.NewPostgresSQL(projectConfig)
	err = psql.Connect(context.Background())
	if err != nil {
		fmt.Println("error while connecting to database: ", err)
		exitAfterDefer(1)
	}
	defer func() {
		psql.Close()
	}()

	blackListRepo := repository.NewListRepo(psql, repository.BlackListTable)
	blackListService := service.NewListService(blackListRepo)

	whiteListRepo := repository.NewListRepo(psql, repository.WhiteListTable)
	whiteListService := service.NewListService(whiteListRepo)

	listHandler := handler.NewListHandler(whiteListService, blackListService)

	authService := service.NewAuthorization(projectConfig, blackListService, whiteListService)
	authHandler := handler.NewAuthHandler(authService)

	bucketHandler := handler.NewBucketHandler(authService)

	router := httpapi.NewAPIRouter(authHandler, bucketHandler, listHandler)
	router.RegisterRoutes()

	httpServer := httpapi.NewServer(router.GetRouter(), &projectConfig)
	notifyContext, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	go httpServer.ShutdownService(notifyContext, cancel)
	err = httpServer.Start()
	if err != nil {
		fmt.Println(err)
		exitAfterDefer(1)
	}
}

func exitAfterDefer(code int) {
	os.Exit(code)
}
