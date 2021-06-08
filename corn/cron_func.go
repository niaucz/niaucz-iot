package corn

import (
	"github.com/robfig/cron"
	"log"
)

func CronFunc(spec string, fun func()) {
	c := cron.New()
	//spec := "* * * * *"
	//slaveID := 1
	//address := 0
	//quantity := 7
	_, err := c.AddFunc(spec, fun)
	if err != nil {
		log.Println("AddFunc failed", err)
		return
	}
	c.Start()
}
