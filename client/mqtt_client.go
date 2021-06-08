package client

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func MqttClient(clientID string, broker string, port int, username string, password string) (mqttClient mqtt.Client) {
	//-------------------MQTT--------------------------
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID(clientID)
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = onConnectHandler
	opts.OnConnectionLost = connectionLostHandler
	opts.OnReconnecting = reconnectHandler
	return mqtt.NewClient(opts)
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Pub message: %#v from topic: %#v\n", msg.Payload(), msg.Topic())
}

var reconnectHandler mqtt.ReconnectHandler = func(client mqtt.Client, options *mqtt.ClientOptions) {
	fmt.Println("MQTT Reconnect")
}

//连接时调用
var onConnectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("MQTT Connected")
}

//连接无响应时候调用
var connectionLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("MQTT Connect lost: %v", err)
}

//发布消息
func Publish(client mqtt.Client, topic string, qos byte, retained bool, payload []byte) {
	token := client.Publish(topic, qos, retained, payload)
	token.Wait()
}
