package device

import (
	"encoding/json"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/rs/zerolog/log"
)

type Consumer struct {
	repository Repository
	client     mqtt.Client
}

func NewConsumer(repository Repository, client mqtt.Client) *Consumer {
	return &Consumer{repository: repository, client: client}
}

func (c *Consumer) Start() {
	if token := c.client.Connect(); token.Wait() && token.Error() != nil {
		log.Error().Err(token.Error()).Msg("connecting to MQTT broker")
		return
	}

	if token := c.client.Subscribe("zigbee2mqtt/bridge/devices", 0, c.messageHandler); token.Wait() && token.Error() != nil {
		log.Error().Err(token.Error()).Msg("subscribing to MQTT topic")
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
		log.Error().Err(err).Msg("unmarshal devices")
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
			log.Error().Err(err).Msg("creating device")
			return
		}
	}
}
