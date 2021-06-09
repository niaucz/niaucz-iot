package client

import (
	"fmt"
	"net"
	"niaucz-iot/entity"
	"time"
)

//获取设备4G模块信息 详情参见设备说明书
const (
	CMD_GET_ATI  = 51
	CMD_GET_IMEI = 52
	CMD_GET_IMSI = 101
	CMD_GET_CSQ  = 102
	CMD_GET_COPS = 103
	CMD_GET_CREG = 111
)

func GetDeviceInfo() (deviceInfo entity.DeviceInfo) {
	n := conn()
	if n == nil {
		fmt.Println("获取连接失败")
		return
	}
	defer n.Close()
	time.Sleep(time.Millisecond * 20)
	imei := sendCmd(CMD_GET_IMEI, n)
	time.Sleep(time.Millisecond * 20)
	imsi := sendCmd(CMD_GET_IMSI, n)
	time.Sleep(time.Millisecond * 20)
	cops := sendCmd(CMD_GET_COPS, n)
	time.Sleep(time.Millisecond * 20)
	creg := sendCmd(CMD_GET_CREG, n)
	time.Sleep(time.Millisecond * 20)
	csq := sendCmd(CMD_GET_CSQ, n)
	return entity.DeviceInfo{
		IMEI: imei,
		IMSI: imsi,
		COPS: cops,
		CREG: creg,
		CSQ:  csq,
	}
}

//获取Socket连接
func conn() net.Conn {
	conn, err1 := net.Dial("unix", "@phone_server")
	if err1 != nil {
		fmt.Printf("conn error:%v\n", err1)
		return nil
	}
	return conn
}

func sendCmd(cmd byte, conn net.Conn) (data string) {
	cmdb := cmd >> 8
	bytes := []byte{'H', 'T', 'S', cmdb, cmd}
	_, err2 := conn.Write(bytes)
	if err2 != nil {
		fmt.Printf("Write error:%v\n", err2)
		return
	}
	buf := make([]byte, 20)
	_, err3 := conn.Read(buf)
	if err3 != nil {
		fmt.Printf("err2:%v\n", err3)
		return
	}
	return string(buf[7:])
}
