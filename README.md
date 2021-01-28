[![Go Report Card](https://goreportcard.com/badge/github.com/TomWright/gracehttpserverrunner)](https://goreportcard.com/report/github.com/TomWright/gracehttpserverrunner)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/tomwright/gracehttpserverrunner)](https://pkg.go.dev/github.com/tomwright/gracehttpserverrunner)
![GitHub License](https://img.shields.io/github/license/TomWright/gracehttpserverrunner)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/TomWright/gracehttpserverrunner?label=latest%20release)

# Grace HTTP Server Runner

A GRPC Server Runner for use with [grace](https://github.com/TomWright/grace).

## Usage

```go
package main

import (
	"context"
	"github.com/tomwright/grace"
	"github.com/tomwright/gracehttpserverrunner"
	"net/http"
	"time"
)

func main() {
	g := grace.Init(context.Background())

	// Create and configure your HTTP server.
	server := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {}),
	}

	// Create and configure the HTTP server runner.
	runner := &gracehttpserverrunner.HTTPServerRunner{
		Server:  server,
		ShutdownTimeout: time.Second * 5,
	}

	// Run the runner.
	g.Run(runner)

	g.Wait()
}
```
