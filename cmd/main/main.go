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
	"log"
)

func main() {
	config := configuration.MustGetConfig()
	clog := logger.New(config.Mode)

	db, err := db.New(context.Background(), config.ConnStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Postgres.Close(context.Background())

	productsRepo := repo.NewProductsRepo(db)
	cartsRepo := repo.NewCartsRepo(db)
	userssRepo := repo.NewUsersRepo(db)

	productsService := services.NewProductsService(productsRepo)
	cartsService := services.NewCartsService(cartsRepo)
	usersService := services.NewUsersService(userssRepo)
	authService := services.NewAuthService(userssRepo, config.JwtSecret)

	go rest.MustStartRestServer(authService, config.RestPort, clog)
	gql.MustStartGqlServer(productsService, cartsService, usersService, config.GqlPort)
}
