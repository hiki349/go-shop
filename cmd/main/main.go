package main

import (
	"context"
	"go-shop/configuration"
	"go-shop/internal/api/gql"
	"go-shop/internal/api/rest"
	"go-shop/internal/domain/services"
	"go-shop/internal/metrics"
	"go-shop/internal/pkg/logger"
	"go-shop/internal/storage/db"
	"go-shop/internal/storage/repo"
)

func main() {
	config := configuration.MustGetConfig()
	clog := logger.New(config.Mode)

	postgres, err := db.NewPostgres(context.Background(), config.ConnStrPostgres)
	if err != nil {
		clog.Error("%w", err)
		return
	}
	defer postgres.Close(context.Background())

	mongo, err := db.NewMongo(context.Background(), config.ConnStrMongo)
	if err != nil {
		clog.Error("%w", err)
		return
	}
	defer mongo.Disconnect(context.Background())

	productsRepo := repo.NewProductsRepo(postgres)
	cartsRepo := repo.NewCartsRepo(postgres)
	usersRepo := repo.NewUsersRepo(postgres)

	productsService := services.NewProductsService(productsRepo)
	cartsService := services.NewCartsService(cartsRepo)
	usersService := services.NewUsersService(usersRepo)
	authService := services.NewAuthService(usersRepo, config.JwtSecret)

	go metrics.Listen("127.0.0.1:8082")

	go rest.MustStartRestServer(authService, config.RestPort, clog)
	gql.MustStartGqlServer(productsService, cartsService, usersService, clog, config.GqlPort)
}
