package device

import (
	"encoding/json"
	"log/slog"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Consumer struct {
	logger     *slog.Logger
	repository Repository
	client     mqtt.Client
}

func NewConsumer(logger *slog.Logger, repository Repository, client mqtt.Client) *Consumer {
	return &Consumer{logger: logger, repository: repository, client: client}
}

func (c *Consumer) Start() {
	if token := c.client.Connect(); token.Wait() && token.Error() != nil {
		c.logger.Error("onnecting to MQTT broker %w", token.Error())
		return
	}

	if token := c.client.Subscribe("zigbee2mqtt/bridge/devices", 0, c.messageHandler); token.Wait() && token.Error() != nil {
		c.logger.Error("subscribing to topic %w", token.Error())
		return
	}
}

type deviceDTO struct {
	FriendlyName string `json:"friendly_name"`
	I3E          string `json:"ieee_address"`
}

func (c *Consumer) messageHandler(client mqtt.Client, msg mqtt.Message) {
	var deviceDTOs []deviceDTO

	if err := json.Unmarshal(msg.Payload(), &deviceDTOs); err != nil {
		c.logger.Error("unmarshalling message payload %w", err)
		return
	}

	for _, deviceDTO := range deviceDTOs {
		_, err := c.repository.CreateDevice(DeviceEntity{
			ID:     deviceDTO.I3E,
			Name:   deviceDTO.FriendlyName,
			Type:   "light",
			Status: "idle",
		})
		if err != nil {
			c.logger.Error("creating device %w", err)
			return
		}
	}
}
