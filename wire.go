//go:build wireinject

//go:generate wire .
package main

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/max-weis/smarthome/device"
	"github.com/max-weis/smarthome/internal"
)

type AppContext struct {
	echo         *echo.Echo
	deviceServer device.ServerInterface
}

func Initialize() (*AppContext, error) {
	panic(wire.Build(
		internal.Set,
		device.Set,
		wire.Struct(new(AppContext), "*"),
	))
}
