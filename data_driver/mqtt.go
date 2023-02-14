package mqtt

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
)

var client MQTT.Client
var mqttbroker string = os.Getenv("MQTT_BROKER")
var mqttuser string = os.Getenv("mqttuser")
var mqttpsd string = os.Getenv("mqttpsd")

type MessageCallback func(Message []byte, topic string)

func init() {
	setupMqttClient(mqttbroker, mqttuser, mqttpsd)
}

func setupMqttClient(broker string, user string, password string) error {
	opts := MQTT.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetUsername(user)
	opts.SetPassword(password)

	client = MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
		return token.Error()
	}

	log.Println("Setup client to broker ", broker)
	return nil
}

func Subscribe(topic string, qos byte, callback MessageCallback) error {

	NewNotificationHandler := func(client MQTT.Client, msg MQTT.Message) {
		callback(msg.Payload(), msg.Topic())
	}
	return client.Subscribe(topic, qos, NewNotificationHandler).Error()
}

func Publish(topic string, qos byte, retain bool, payload string) error {
	return client.Publish(topic, qos, retain, payload).Error()
}
