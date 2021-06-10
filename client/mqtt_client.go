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
	opts.SetDefaultPublishHandler(messageHandler)
	//opts.OnConnectAttempt = OnConnectAttempt
	opts.OnConnect = onConnect
	opts.OnConnectionLost = connectionLostHandler
	opts.OnReconnecting = reconnectHandler
	client := mqtt.NewClient(opts)
	//订阅下发的指令
	Sub(client)
	return client
}

//var OnConnectAttempt mqtt.ConnectionAttemptHandler = func(broker *url.URL, tlsCfg *tls.Config) *tls.Config {
//	fmt.Println("ConnectionAttemptHandler")
//}

var onConnect mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("OnConnectHandler")
}

var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Println("MessageHandler")
	fmt.Printf("Pub message: %#v from topic: %#v\n", msg.Payload(), msg.Topic())
}

var reconnectHandler mqtt.ReconnectHandler = func(client mqtt.Client, options *mqtt.ClientOptions) {
	fmt.Println("ReconnectHandler")
}

//连接无响应时候调用
var connectionLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("MQTT Connect lost: %v", err)
}

//订阅消息
func Sub(client mqtt.Client) {
	topic := "topic/command"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s", topic)
}
