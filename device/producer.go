package device

import (
	"encoding/json"
	"fmt"
	"log/slog"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Producer struct {
	client mqtt.Client
	logger *slog.Logger
}

func NewProducer(logger *slog.Logger, client mqtt.Client) *Producer {
	return &Producer{logger: logger, client: client}
}

func (p *Producer) PublishConfiguration(deviceId string, configId string, configData map[string]any) error {
	data, err := json.Marshal(configData)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", p)

	token := p.client.Publish("zigbee2mqtt/%s/set", 0, false, data)
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}

	p.logger.Info("published configuration %s for device %s", configId, deviceId)

	return nil
}
