//go:build wireinject
// +build wireinject

package main

import (
    "github.com/google/wire"
    "github.com/AntonioTWize/goledger/db"
    "github.com/AntonioTWize/goledger/handlers"
    "github.com/AntonioTWize/goledger/repositories"
    "github.com/AntonioTWize/goledger/routes"
    "github.com/labstack/echo/v4"
)

func InitializeServer() (*echo.Echo, error) {
    wire.Build(
        db.Connect,                    // Use the existing database connection function
        repositories.NewChargeRepository, // Inject the repository
        handlers.NewChargeHandler,        // Inject the handler
        routes.NewRouter,                 // Inject the router
    )
    return &echo.Echo{}, nil
}
