package main

import (
	"context"
	"go-shop/configuration"
	"go-shop/internal/api/gql"
	"go-shop/internal/api/rest"
	"go-shop/internal/domain/services"
	"go-shop/internal/pkg/logger"
	"go-shop/internal/storage/db"
	"go-shop/internal/storage/repo"
)

func main() {
	config := configuration.MustGetConfig()
	clog := logger.New(config.Mode)

	postgres, err := db.NewPostgres(context.Background(), config.ConnStrPostgres)
	defer postgres.Close(context.Background())
	if err != nil {
		clog.Error("%w", err)
		return
	}

	mongo, err := db.NewMongo(context.Background(), config.ConnStrMongo)
	defer mongo.Disconnect(context.Background())
	if err != nil {
		clog.Error("%w", err)
		return
	}

	productsRepo := repo.NewProductsRepo(postgres)
	cartsRepo := repo.NewCartsRepo(postgres)
	usersRepo := repo.NewUsersRepo(postgres)

	productsService := services.NewProductsService(productsRepo)
	cartsService := services.NewCartsService(cartsRepo)
	usersService := services.NewUsersService(usersRepo)
	authService := services.NewAuthService(usersRepo, config.JwtSecret)

	go rest.MustStartRestServer(authService, config.RestPort, clog)
	gql.MustStartGqlServer(productsService, cartsService, usersService, clog, config.GqlPort)
}
