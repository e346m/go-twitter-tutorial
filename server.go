package main

import (
    //"strconv"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"

    "./handler"
)

func main() {
    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    api := e.Group("/api/v1")
    api.GET("/tweets", handler.FetchTweetListHandler)
    api.GET("/tweets/:id", handler.FetchTweetHandler)

    e.Logger.Fatal(e.Start(":4242"))
}

