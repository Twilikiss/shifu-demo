// Package main
// @Author twilikiss 2024/12/13 17:17:17
package main

import (
	"os"
	"os/signal"
	"shifu-demo/config"
	"shifu-demo/log"
	"shifu-demo/task"
	"strconv"
	"syscall"
)

func main() {
	url := config.Cfg.ServiceConfig.Url
	time, _ := strconv.Atoi(config.Cfg.ServiceConfig.Time)
	log.Infof("任务中心启动成功，url: %s, time: %d m", url, time)

	// 创建并启动任务
	t := task.NewTask()
	t.Run(url, time)

	//优雅退出
	go func() {
		exit := make(chan os.Signal)
		signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-exit:
			log.Info("任务中心中断执行，开始clear资源")
			t.Stop()
		}
	}()

	t.StartBlocking()
}
