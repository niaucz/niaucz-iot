package client

import (
	"github.com/goburrow/serial"
	modbus "github.com/things-go/go-modbus"
	"log"
	"time"
)

func ModbusClient(Address string, BaudRate int, DataBits int, StopBits int, Parity string, Timeout int) (modbusClient modbus.Client) {
	//----------------------ModBus----------------------------
	p := modbus.NewRTUClientProvider(modbus.WithEnableLogger(),
		modbus.WithSerialConfig(serial.Config{
			Address:  Address,
			BaudRate: BaudRate,
			DataBits: DataBits,
			StopBits: StopBits,
			Parity:   Parity,
			Timeout:  time.Duration(Timeout * 1000000000),
		}))
	return modbus.NewClient(p)
}

func main() {
	port, err := serial.Open(&serial.Config{Address: "/dev/buzzer"})
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	_, err = port.Write([]byte("1"))
	if err != nil {
		log.Fatal(err)
	}
}
