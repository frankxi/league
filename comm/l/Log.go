package logSetting

import (
	"os"
	log "github.com/sirupsen/logrus"
)

//打印日志
func Setup() {
	// 设置日志格式为json格式
	//log.SetFormatter(&log.JSONFormatter{})
	// 设置将日志输出到标准输出（默认的输出为stderr，标准错误）
	// 日志消息输出可以是任意的io.writer类型
	log.SetOutput(os.Stdout)
	log.StandardLogger()
	// 设置日志级别为warn以上
	log.SetLevel(log.DebugLevel)
}
