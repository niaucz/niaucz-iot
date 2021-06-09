package main

import (
	"github.com/goburrow/serial"
	"log"
	"time"
)

//蜂鸣器
func main() {
	port, err := serial.Open(&serial.Config{Address: "/dev/buzzer"})
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	//while(1)
	//{
	//	write(fd,"1",1);
	//	usleep(500000);
	//	write(fd,"0",1);
	//	usleep(500000);
	//}
	_, err = port.Write([]byte{'1'})
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Millisecond * 100)
	_, err = port.Write([]byte{'0'})
	if err != nil {
		log.Fatal(err)
	}
}
