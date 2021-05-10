package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/fx"

	"github.com/CassioRoos/poc-lazy-loading/handlers"
	"github.com/CassioRoos/poc-lazy-loading/internal"
)

func main() {
	fx.New(
		handlers.Module,
		internal.Module,
		fx.Invoke(
			StartServer,
		),
	).Run()
}

func StartServer(lc fx.Lifecycle, router http.Handler) {
	server := &http.Server{
		Addr:         ":8080",           // configure the bind address
		Handler:      router,            // set the default handler
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				fmt.Println("Starting the server at port :8080")

				go server.ListenAndServe() // nolint:errcheck if error occurs everything fails

				return nil
			},
			OnStop: func(ctx context.Context) error {
				fmt.Println("Stopping the server...")

				return server.Shutdown(ctx)
			},
		},
	)
}
