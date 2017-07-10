package wooUtils

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

//StrFormatTime 字符串格式化为当前时间
//%Y %m %d %H %i %s : year month day hour minutes second
func StrFormatTime(instr string) (outstr string) {
	//2006-01-02 15:04:05
	nowtime := time.Now()
	year := nowtime.Format("2006")
	month := nowtime.Format("01")
	day := nowtime.Format("02")
	hour := nowtime.Format("15")
	minutes := nowtime.Format("04")
	sencond := nowtime.Format("05")
	outstr = strings.Replace(instr, "%Y", year, -1)
	outstr = strings.Replace(outstr, "%m", month, -1)
	outstr = strings.Replace(outstr, "%d", day, -1)
	outstr = strings.Replace(outstr, "%H", hour, -1)
	outstr = strings.Replace(outstr, "%i", minutes, -1)
	outstr = strings.Replace(outstr, "%s", sencond, -1)
	return
}

//Try 错误处理 try ...catch 仿
func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}

//GetCurrentPath 获取runtime路径
func GetCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New("error: Can't find \\ or /")
	}
	return string(path[0 : i+1]), nil
}

//GetWeekMondayTime 获取本周一的同一时间的时间戳
//因为西方计算方式每周从周日开始, 中国计算方式一般从周一开始, 所以遇到周日需要特别考虑
func GetWeekMondayTime() (mtime time.Time) {
	nowTime := time.Now()
	weekDay := int(nowTime.Weekday())
	if weekDay == 0 {
		weekDay = 7
	}
	fmt.Println(nowTime.Unix())
	return nowTime.Add(time.Duration(weekDay-1) * -86400 * time.Second)
}
