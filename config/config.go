package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

func LogInfo() error {
	file := "./" + time.Now().Format("2006-01-01") + ".log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		return fmt.Errorf("open file failed")
	}

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(logFile)

	return nil
}

func Init() error {
	// 初始化配置文件.
	if err := Config(); err != nil {
		return err
	}

	//初始化
	err := LogInfo()
	if err != nil {
		fmt.Println("log 日志管理 初始化失败.")
	}
	return nil
}

// config viper解析配置文件
func Config() error {
	viper.AddConfigPath("conf")
	viper.SetConfigFile("config")
	if err := viper.ReadConfig(); err != nil {
		return err
	}

	return nil
}
