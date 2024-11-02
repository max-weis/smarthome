// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/labstack/echo/v4"
	"github.com/max-weis/smarthome/device"
	"github.com/max-weis/smarthome/internal"
)

// Injectors from wire.go:

func Initialize() (*AppContext, error) {
	echo := internal.NewEchoServer()
	db, err := internal.NewDatabase()
	if err != nil {
		return nil, err
	}
	repository := device.NewRepository(db)
	client := internal.NewMqttClient()
	consumer := device.NewConsumer(repository, client)
	producer := device.NewMQTTProducer(client)
	serverInterface := device.NewHandler(echo, repository, producer)
	appContext := &AppContext{
		echo:           echo,
		deviceConsumer: consumer,
		deviceServer:   serverInterface,
	}
	return appContext, nil
}

// wire.go:

type AppContext struct {
	echo           *echo.Echo
	deviceConsumer *device.Consumer
	deviceServer   device.ServerInterface
}
