package entity

type ModbusData struct {
	//采集器ID
	ClientID string `json:"clientId"`
	//从站地址
	SlaveID int `json:"slaveID"`
	//寄存器开始地址
	Address int `json:"address"`
	//寄存器数
	Quantity int `json:"quantity"`
	//modbus 数据
	Data []int `json:"data"`
}
