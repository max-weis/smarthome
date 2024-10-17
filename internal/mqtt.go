package internal

import mqtt "github.com/eclipse/paho.mqtt.golang"

func NewMqttClient() mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://raspberrypi:1883")
	opts.SetClientID("smarthome-client")

	return mqtt.NewClient(opts)
}
