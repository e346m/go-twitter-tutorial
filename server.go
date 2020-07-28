package main

import (
    "net/http"
    //"strconv"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
)

func main() {
    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    api := e.Group("/api/v1")
    api.GET("/tweets", fetchTweetListHandler)
    api.GET("/tweets/:id", fetchTweetHandler)

    e.Logger.Fatal(e.Start(":4242"))

}

type Tweet struct {
    User string `json:"user"`
    Content string `json:"content"`
}

func fetchTweetHandler (c echo.Context) error {
    id := c.Param("id")
    return c.JSON(http.StatusOK, Tweet{User: id, Content: "Hello world again!"})
}

func fetchTweetListHandler (c echo.Context) error {
    return c.JSON(http.StatusOK, []Tweet{
        { User: "mishiro", Content: "Hello world!" },
        { User: "reiji", Content: "See you tomorrow!" },
    })
}
