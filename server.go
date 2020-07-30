package main

import (
    //"strconv"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"

    "go-twitter-tutorial.com/m/handler"
    "go-twitter-tutorial.com/m/model"
)

func main() {
    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    h := &handler.Handler{&model.Tweet{}}

    api := e.Group("/api/v1")
    //api.GET("/tweets", h.FetchTweetListHandler)
    api.GET("/tweets/:id", h.FetchTweetHandler)

    e.Logger.Fatal(e.Start(":4242"))
}

