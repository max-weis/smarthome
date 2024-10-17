package internal

import "github.com/google/wire"

var Set = wire.NewSet(
	NewLogger,
	NewEchoServer,
	NewDatabase,
	NewMqttClient,
)
