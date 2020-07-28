package handler

import (
    "fmt"
    "reflect"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"

    "github.com/labstack/echo"
)

var (
    userJSON = `{"user":"1","content": "Hello world"}` + "\n"
)

func TestFetchTweetHandler (t *testing.T) {
    e := echo.New()
    req := httptest.NewRequest(http.MethodGet, "/api/v1", strings.NewReader(userJSON))
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)
    c.SetPath("/tweets/:id")
    c.SetParamNames("id")
    c.SetParamValues("1")

    FetchTweetHandler(c)

    if http.StatusOK != rec.Code {
        t.Errorf(
            "response status code: %d, expected: %d", rec.Code, http.StatusOK,
        )
    }

    fmt.Println(reflect.DeepEqual(userJSON, rec.Body.String()))

    if userJSON != rec.Body.String() {
        t.Errorf(
            "response body: %v, expected: %v", rec.Body.String(), userJSON,
        )
    }
}
