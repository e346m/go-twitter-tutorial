package handler

import (
    "net/http"
    "github.com/labstack/echo"
)


func (h *Handler) FetchTweetHandler (c echo.Context) error {
    id := c.Param("id")
    h.ItemGetter.GetItem(id)
    return c.JSON(http.StatusOK, h.ItemGetter)
}

/*func (h *Handler) FetchTweetListHandler (c echo.Context) error {
    h.ItemGetter.GetItem(id)
    return c.JSON(http.StatusOK, )
}
*/
