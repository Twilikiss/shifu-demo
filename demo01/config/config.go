// Package config
// @Author twilikiss 2024/12/13 18:21:21
package config

import (
	"github.com/BurntSushi/toml"
	"os"
	"shifu-demo/log"
)

var Cfg *pollConfig

func init() {
	path := "demo01/etc/config.toml"
	if !fileIsExist(path) {
		log.Error("配置文件不存在")
		panic("[etc\\config.toml]无法找到配置文件")
	}
	Cfg = new(pollConfig)
	_, err := toml.DecodeFile(path, &Cfg)
	if err != nil {
		log.Error("配置文件读取失败")
		panic("[etc\\config.toml]无法找到配置文件")
	}
}

func fileIsExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
