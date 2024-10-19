package device

import "github.com/google/wire"

var Set = wire.NewSet(
	NewRepository,
	NewHandler,
	NewConsumer,
	NewProducer,
)
