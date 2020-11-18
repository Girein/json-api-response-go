# JSON API Response for Go

## Installation
`go get github.com/Girein/json-api-response-go`

## Usage
In the example below, I'm using the [Echo](https://github.com/labstack/echo) web framework.
```
import pretty "github.com/Girein/json-api-response-go"

func Hello(c echo.Context) error {
    response := new(pretty.JSONResponse)
    response.Simple("200", "OK", "Hello World!")

    return c.JSON(200, response)
}
```