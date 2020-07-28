package handler

import (
    "net/http"
    "github.com/labstack/echo"
)

type Tweet struct {
    User string `json:"user"`
    Content string `json:"content"`
}

func FetchTweetHandler (c echo.Context) error {
    id := c.Param("id")
    return c.JSON(http.StatusOK, Tweet{User: id, Content: "Hello world again!"})
}

func FetchTweetListHandler (c echo.Context) error {
    return c.JSON(http.StatusOK, []Tweet{
        { User: "mishiro", Content: "Hello world!" },
        { User: "reiji", Content: "See you tomorrow!" },
    })
}
