package device

import (
	"encoding/json"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/rs/zerolog/log"
)

type Producer struct {
	client mqtt.Client
}

func NewProducer(client mqtt.Client) *Producer {
	return &Producer{client: client}
}

func (p *Producer) PublishConfiguration(deviceId string, configId string, configData map[string]any) error {
	data, err := json.Marshal(configData)
	if err != nil {
		log.Error().Err(err).Msg("marshalling configuration data")
		return err
	}

	token := p.client.Publish("zigbee2mqtt/%s/set", 0, false, data)
	if token.Wait() && token.Error() != nil {
		log.Error().Err(token.Error()).Msg("publishing configuration")
		return token.Error()
	}

	log.Info().Msgf("published configuration %s for device %s", configId, deviceId)

	return nil
}
