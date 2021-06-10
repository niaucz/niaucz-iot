package corn

import (
	"github.com/robfig/cron"
	"log"
)

// 返回一个支持至 秒 级别的 cron
func newWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}

func CronFunc(spec string, fun func()) {
	//c := cron.New()
	c := newWithSeconds()
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
