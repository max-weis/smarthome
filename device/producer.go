package device

import (
	"encoding/json"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/rs/zerolog/log"
)

type mqttProducer struct {
	client mqtt.Client
}

func NewMQTTProducer(client mqtt.Client) Producer {
	return &mqttProducer{client: client}
}

func (p *mqttProducer) PublishConfiguration(deviceId string, configId string, configData map[string]any) error {
	data, err := json.Marshal(configData)
	if err != nil {
		log.Error().Err(err).Msg("marshalling configuration data")
		return err
	}

	token := p.client.Publish(fmt.Sprintf("zigbee2mqtt/%s/set", deviceId), 0, false, data)
	if token.Wait() && token.Error() != nil {
		log.Error().Err(token.Error()).Msg("publishing configuration")
		return token.Error()
	}

	log.Info().Msgf("published configuration %s for device %s", configId, deviceId)

	return nil
}
