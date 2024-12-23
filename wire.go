//go:build wireinject
// +build wireinject

package main

import (
    "github.com/google/wire"
    "github.com/AntonioTWize/goledger/handlers"
    "github.com/AntonioTWize/goledger/routes"
    
    "github.com/labstack/echo/v4"
)

func InitializeServer() (*echo.Echo, error) {
    wire.Build(
        handlers.NewChargeHandler,
        routes.NewRouter,
    )
    return &echo.Echo{}, nil
}
