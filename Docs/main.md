# main

## Table of Contents

- [Package](#package)
- [Variables](#variables)
- [Functions](#functions)

## Package

Package main is the entrypoint to the RESTful api. It creates the server and
hosts it with a goroutine. The server can be stopped with a system
interrupt.

## VARIABLES

```go
var routes = []struct {
        path    string
        method  string
        handler http.HandlerFunc
}{
        {
                path:    "/advantage",
                method:  http.MethodGet,
                handler: handlers.Advantage,
        },
        {
                path:    "/commonMoves",
                method:  http.MethodGet,
                handler: handlers.CommonMoves,
        },
}
```

routes is a structure containing all the routes used with this RESTful api.

## FUNCTIONS

```go
func main()
```

main is in charge of setting the port and executing the server's goroutine.
