package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	ctx, err := Initialize()
	if err != nil {
		panic(err)
	}

	// handle ui
	ctx.echo.Static("/", "ui/dist")

	var wg sync.WaitGroup

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalChan
		ctx.echo.Close()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := ctx.echo.Start(":8080"); err != nil {
			ctx.echo.Logger.Fatal(err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		ctx.deviceConsumer.Start()
	}()

	wg.Wait()
	ctx.echo.Logger.Info("Server shutting down gracefully")
}
