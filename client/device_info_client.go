package client

import (
	"fmt"
	"net"
	"niaucz-iot/entity"
	"time"
)

//获取设备4G模块信息 详情请参见设备说明书
const (
	CMD_GET_ATI  = 51
	CMD_GET_IMEI = 52
	CMD_GET_IMSI = 101
	CMD_GET_CSQ  = 102
	CMD_GET_COPS = 103
	CMD_GET_CREG = 111
)

//func main() {
//	info := GetDeviceInfo()
//	fmt.Println(info)
//}
func GetDeviceInfo() (deviceInfo entity.DeviceInfo) {
	n := conn()
	if n == nil {
		fmt.Println("获取连接失败")
		return
	}
	defer n.Close()
	time.Sleep(time.Millisecond * 20)
	imei := sendCmd(CMD_GET_IMEI, n)
	fmt.Printf("imei:%s\n", imei)

	time.Sleep(time.Millisecond * 20)
	imsi := sendCmd(CMD_GET_IMSI, n)
	fmt.Printf("imsi:%s\n", imsi)

	time.Sleep(time.Millisecond * 20)
	cops := sendCmd(CMD_GET_COPS, n)
	fmt.Printf("cops:%s\n", cops)

	time.Sleep(time.Millisecond * 20)
	creg := sendCmd(CMD_GET_CREG, n)
	fmt.Printf("creg:%s\n", creg)

	time.Sleep(time.Millisecond * 20)
	csq := sendCmd(CMD_GET_CSQ, n)
	fmt.Printf("csq:%s\n", csq)

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
	buf := make([]byte, 50)
	_, err3 := conn.Read(buf)
	if err3 != nil {
		fmt.Printf("err2:%v\n", err3)
		return
	}
	return string(buf[7:])
}

//fmt.Printf("[THR s]: %s\n", string(buf[:3]))
//fmt.Printf("[->]: %v\n", buf[3:5])
//fmt.Printf("[->]: %v\n", buf[5:7])
//fmt.Printf("[THR s]: %s\n", string(buf[7:]))
//fmt.Printf("[3->]: %v\n", (buf[3]<<8)+buf[4])
//fmt.Printf("[5->]: %v\n", (buf[5]<<8)+buf[6])
//fmt.Println("----------------------")
