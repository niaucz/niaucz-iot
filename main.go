package main

import (
	"encoding/json"
	"fmt"
	"github.com/things-go/go-modbus"
	"log"
	"niaucz-iot/client"
	"niaucz-iot/common"
	"niaucz-iot/config"
	"niaucz-iot/corn"
	"niaucz-iot/entity"
	"time"
)

//交叉编译
//SET CGO_ENABLED=0
//SET GOOS=linux
//SET GOARCH=arm
//go build

//订阅消息
//func sub(client mqtt.Client) {
//	topic := "topic/test"
//	token := client.Subscribe(topic, 1, nil)
//	token.Wait()
//	fmt.Printf("Subscribed to topic: %s", topic)
//}
// 返回一个支持至 秒 级别的 cron
//func newWithSeconds() *cron.Cron {
//	secondParser := cron.NewParser(cron.Second | cron.Minute |
//		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
//	return cron.New(cron.WithParser(secondParser), cron.WithChain())
//}

func main() {
	//读取配置文件
	ini := config.LoadIni()
	if ini == nil {
		log.Println("无法从网络和本地加载配置文件 请检查")
		return
	}
	//mqtt和modbus客户端调用
	mqttClient := client.MqttClient(common.DeviceID, ini.Broker, ini.Port, ini.Username, ini.Password)
	modbusClient := client.ModbusClient(ini.SerialAddress, ini.BaudRate, ini.DataBits, ini.StopBits, ini.Parity, ini.Timeout)
	//释放资源
	defer func(modbusClient modbus.Client) {
		err := modbusClient.Close()
		if err != nil {

		}
	}(modbusClient)
	defer mqttClient.Disconnect(250)

	//连接mqttClient
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	//连接modbusClient
	err := modbusClient.Connect()
	if err != nil {
		log.Println("connect failed, ", err)
		return
	}

	log.Println("Start monitoring data...")

	corn.CronFunc(ini.Cron, func() {
		registers := modbusClientReadHoldingRegisters(modbusClient, common.DeviceID, ini.SlaveID, ini.StartAddress, ini.Quantity)
		mqttClient.Publish(ini.Topic, byte(ini.Qos), ini.Retained == 1, registers)
	})
	for {
		time.Sleep(time.Hour)
	}
}

//01 03 0e 05 02 00 86 00 35 00 00 5d 7e 72 df 00 00 3e 8c
func modbusClientReadHoldingRegisters(modbusClient modbus.Client, clientID string, slaveID, address, quantity int) (marshal []byte) {
	results, err := modbusClient.ReadHoldingRegistersBytes(byte(slaveID), uint16(address), uint16(quantity))
	if err != nil {
		log.Println("ReadHoldingRegistersBytes failed, ", err)
		return nil
	}
	fmt.Printf("ReadDiscreteInputs %v\n", results)
	result := make([]int, len(results))
	for i := 0; i < len(results); i++ {
		result[i] = int(results[i])
	}

	var modbusD = entity.ModbusData{
		ClientID: clientID,
		SlaveID:  slaveID,
		Address:  address,
		Quantity: quantity,
		Data:     result,
	}
	marshal, jsonErr := json.Marshal(modbusD)
	if jsonErr != nil {
		log.Println("json Marshal failed", err)
		return nil
	}
	return marshal
}
