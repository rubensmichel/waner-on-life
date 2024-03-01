package main

import (
	"context"
	"fmt"

	"github.com/rubensmichel/waner-on-life/internal"
	"github.com/rubensmichel/waner-on-life/internal/infra/gracefully"
	httpserver "github.com/rubensmichel/waner-on-life/internal/infra/transport/http/server"
)

func main() {
	fmt.Print(context.Background(), "Starting Fraud Prevention SPD application")

	ft, err := internal.NewFactory()
	if err != nil {
		fmt.Print(err)
	}

	http := httpserver.New(ft)

	go http.Listen()

	gracefully.New().
		Add(http).
		Add(ft).
		Wait()
}
