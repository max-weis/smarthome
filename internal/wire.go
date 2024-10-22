package internal

import "github.com/google/wire"

var Set = wire.NewSet(
	NewEchoServer,
	NewDatabase,
	NewMqttClient,
)
