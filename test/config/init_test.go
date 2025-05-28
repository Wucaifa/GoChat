package config

import (
	"GoChat/internal/config"
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	conf := config.GetConfig()
	fmt.Println(conf.MainConfig)
	fmt.Println(conf.MysqlConfig)
	fmt.Println(conf.RedisConfig)
	fmt.Println(conf.AuthCodeConfig)
}
