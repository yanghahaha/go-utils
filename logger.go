package wooUtils

import (
	"os"
	"time"

	"strings"

	"errors"

	"fmt"

	logger "github.com/apsdehal/go-logger"
)

const (
	LEVEL_ERROR   = 1
	LEVEL_WARNING = 2
	LEVEL_INFO    = 3
	LEVEL_DEBUG   = 4
)

//ELogger 日志类
type ELogger struct {
	inConfig   bool
	level      int
	logger     logger.Logger
	loggerList []logger.Logger
}

func (log *ELogger) Debug(msg string) bool {
	if log.inConfig {
		time.Sleep(10000000)
		log.Debug(msg)
	}
	if log.level >= LEVEL_DEBUG {
		len := len(log.loggerList)
		for i := 0; i < len; i++ {
			log.loggerList[i].Debug(msg)
		}
		return true
	}
	return false
}
func (log *ELogger) Info(msg string) bool {
	if log.inConfig {
		time.Sleep(10000000)
		log.Info(msg)
	}
	if log.level >= LEVEL_INFO {
		len := len(log.loggerList)
		for i := 0; i < len; i++ {
			log.loggerList[i].Info(msg)
		}
		return true
	}
	return false
}
func (log *ELogger) Error(msg string) bool {
	if log.inConfig {
		time.Sleep(10000000)
		log.Error(msg)
	}
	if log.level >= LEVEL_ERROR {
		len := len(log.loggerList)
		for i := 0; i < len; i++ {
			log.loggerList[i].Error(msg)
		}
		return true
	}
	return false
}
func (log *ELogger) Warning(msg string) bool {
	if log.inConfig {
		time.Sleep(10000000)
		log.Warning(msg)
	}
	if log.level >= LEVEL_WARNING {
		len := len(log.loggerList)
		for i := 0; i < len; i++ {
			log.loggerList[i].Warning(msg)
		}
		return true
	}
	return false
}

func (log *ELogger) initOutput(outConfigStr string) (err error) {
	//解析outconfigstr
	configList := strings.Split(outConfigStr, "|")
	configLen := len(configList)
	if configLen == 0 {
		return errors.New("没有日志配置")
	}
	for i := 0; i < configLen; i++ {
		outStr := configList[i]
		//判断输出类型
		if outStr == "stdOut" { //控制台输出
			logwriter, err := logger.New("AMS", 1, os.Stdout)
			if err != nil {
				println("日志初始化出错: " + err.Error())
			}
			log.loggerList = append(log.loggerList, *logwriter)
		} else if strings.Contains(outStr, "file://") { //日志文件输出
			logfilePath := StrFormatTime(strings.Replace(outStr, "file://", "", 1))
			logfile, err := os.OpenFile(logfilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
			if err != nil {
				println("日志初始化出错: " + err.Error())
			}
			logwriter, err := logger.New("AMS", 1, logfile)
			if err != nil {
				println("日志初始化出错: " + err.Error())
			}
			log.loggerList = append(log.loggerList, *logwriter)
		} else { //不支持的类型
			println("不支持的日志配置:" + outStr)
		}
	}
	return nil
}

//Config 配置logger
// LEVEL_ERROR   = 1
// LEVEL_WARNING = 2
// LEVEL_INFO    = 3
// LEVEL_DEBUG   = 4
// file:///xxx/xxx/xxx/%y%m%d|stdOut
func (log *ELogger) Config(level int, outstr string) bool {
	if level != LEVEL_DEBUG && level != LEVEL_ERROR && level != LEVEL_WARNING && level != LEVEL_INFO {
		level = LEVEL_ERROR
	}
	log.level = level
	log.loggerList = []logger.Logger{}
	log.inConfig = true
	err := log.initOutput(outstr)
	log.inConfig = false
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

var _loggerInstance *ELogger

//GetLogger 获取logger单例
func GetLogger() *ELogger {
	if _loggerInstance == nil {
		_loggerInstance = new(ELogger)
	}
	return _loggerInstance
}
