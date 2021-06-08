package entity

type IniConfig struct {
	//MQTT
	Broker   string `ini:"Broker"`
	Port     int    `ini:"Port"`
	Username string `ini:"Username"`
	Password string `ini:"Password"`
	Qos      int    `ini:"Qos"`
	Topic    string `ini:"Topic"`
	Retained int    `ini:"Retained"`
	Cron     string `ini:"Cron"`
	//Modbus
	// Device path (/dev/ttyS0)
	SerialAddress string `ini:"SerialAddress"`
	// Baud rate (default 19200)
	BaudRate int `ini:"BaudRate"`
	// Data bits: 5, 6, 7 or 8 (default 8)
	DataBits int `ini:"DataBits"`
	// Stop bits: 1 or 2 (default 1)
	StopBits int `ini:"StopBits"`
	// Parity: N - None, E - Even, O - Odd (default E)
	// (The use of no parity requires 2 stop bits.)
	Parity string `ini:"Parity"`
	// Read (Write) timeout unit Second
	Timeout int `ini:"Timeout"`
	//寄存器开始地址
	StartAddress int `ini:"StartAddress"`
	//寄存器数
	Quantity int `ini:"Quantity"`
	//从站地址
	SlaveID int `ini:"SlaveID"`
}
