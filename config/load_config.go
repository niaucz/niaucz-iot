package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"niaucz-iot/client"
	"niaucz-iot/common"
	"niaucz-iot/entity"
	"reflect"
	"strconv"
	"strings"
)

func LoadIni() (conf *entity.IniConfig) {
	system := LoadIniFileSystem()
	if system == nil {
		log.Println("无法从文件系统加载配置文件")
		network := LoadIniNetwork()
		if network == nil {
			log.Println("无法从文网络加载配置文件")
			log.Println("无法加载配置文件")

			return nil
		}
		return network
	}
	return system

}

func LoadIniNetwork() (conf *entity.IniConfig) {
	//
	url := "http://config-server:6969/getConfig/" + common.DeviceID
	method := http.MethodGet
	request := client.HttpRequest(url, method)
	if request != nil {
		conf := &entity.IniConfig{}
		err := json.Unmarshal(request, conf)
		if err != nil {
			log.Print("json.Unmarshal err", err)
			return nil
		}
		return conf
	}
	return
}

func LoadIniFileSystem() (conf *entity.IniConfig) {
	//TODO 默认文件位置 ...
	file, err := ioutil.ReadFile("C:/Users/ksd/go/src/niaucz-iot/conf.ini")
	if err != nil {
		log.Print(err)
		return
	}
	split := strings.Split(string(file), "\r\n")
	//var configMap map[string]string
	config := &entity.IniConfig{}
	of := reflect.ValueOf(config).Elem()
	for _, line := range split {
		if len(line) == 0 {
			continue
		}
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") {
			continue
		}
		sp := strings.Split(line, "=")
		if len(sp) != 2 {
			continue
		}
		name := of.FieldByName(sp[0])
		if name.Kind() == 0 {
			continue
		}
		if name.Type().Kind() == reflect.String {
			name.SetString(sp[1])
		}
		if name.Type().Kind() == reflect.Int {
			parseInt, err := strconv.ParseInt(sp[1], 10, 64)
			if err != nil {
				log.Print(err)
				break
			}
			name.SetInt(parseInt)
		}
	}
	return config
}
