package config

import (
	"io/ioutil"
	"log"
	"niaucz-iot/entity"
	"reflect"
	"strconv"
	"strings"
)

func LodIni(filePath string) (conf *entity.IniConfig) {
	file, err := ioutil.ReadFile(filePath)
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
