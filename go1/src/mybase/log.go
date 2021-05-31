package mybase

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"path/filepath"
	"runtime"
	"strings"
)

const TimeFmtLog = "2006/01/02 15:04:05.000" //毫秒保留3位有效数字

var (
	LOG = logrus.New()

	LogDir           = "./log"
	LogName          = "log"
	LogFile *os.File = nil
)

var logSaveDay = 365 * 24 * time.Hour //默认日志保存365天

type PrintHook struct {
}

//过滤等级
func (imp *PrintHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.DebugLevel, logrus.InfoLevel, logrus.WarnLevel,
		logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel}
}

//打印
func (imp *PrintHook) Fire(entry *logrus.Entry) error {
	//fmt.Println(entry.Time.Date())
	//go time Format 必须使用这个时间 2006-01-02 15:04:05.000
	if entry.Logger.Out == os.Stderr { //如果输出还没有到其他地方，那么我们就不用再打印一遍
		return nil
	}

	if strings.HasSuffix(entry.Message, "\n") {
		entry.Message = entry.Message[0 : len(entry.Message)-1]
	}

	if len(entry.Data) != 0 {
		fmt.Printf("[%s]%s:%s,%s\n", entry.Level, entry.Time.Format(TimeFmtLog), entry.Message, entry.Data)
	} else {
		fmt.Printf("[%s]%s:%s\n", entry.Level, entry.Time.Format(TimeFmtLog), entry.Message)
	}

	return nil
}

/**
@duration 日志时间保留最近多少天的。
*/
func initLogFile() error {
	now := time.Now()
	logFilePath := fmt.Sprintf("%s/%s-%s.log", LogDir, LogName, now.Format(DateFmtDB))
	remFilePath := fmt.Sprintf("%s/%s-%s.log", LogDir, LogName, now.Add(-logSaveDay).Format(DateFmtDB))
	_ = os.Remove(remFilePath)

	// You could set this to any `io.Writer` such as a file
	//D("LogFile =", LogFile)
	if LogFile != nil { //检查之前的File
		//D("True")
		_ = LogFile.Close()
	} else {
		//D("False")
	}

	LogFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		LOG.Panicf("创建/打开日志文件[%s]失败: err=%s\n", logFilePath, err)
	} else {
		LOG.Out = LogFile
	}
	return err
}

/**
@day 保留多少天的日志
*/
func initLogDir(linuxPath, fileName string, day int, isInProduct bool) error {
	var osname = string(runtime.GOOS)
	fmt.Println("os is", osname)
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
		return err
	}
	if runtime.GOOS == "linux" && isInProduct { //linux需要制定一个比较大的磁盘
		//strings.Replace(dir, "\\", "/home", -1)
		dir = linuxPath + dir
	}
	dir += "/log"
	err = os.MkdirAll(dir, os.ModeDir|os.ModePerm)
	if err != nil && !os.IsExist(err) {
		LOG.Panicf("创建日志目录[%s]失败: err=%s\n", dir, err)
		return err
	}

	LogDir = dir
	LogName = fileName
	logSaveDay = time.Duration(day) * 24 * time.Hour

	return initLogFile()
}

func InitLogModule(linuxPath, fileName string, day int, isInProduct bool) error {
	// do something here to set environment depending on an environment variable
	// or command-line flag
	if isInProduct {
		//LOG.Formatter = &logrus.JSONFormatter{
		//	TimestampFormat: TimeFmtLog,
		//} //为了方便使用Logstash
		LOG.Formatter = &logrus.TextFormatter{
			TimestampFormat: TimeFmtLog,
		}
		LOG.SetLevel(logrus.InfoLevel)
	} else {
		LOG.Formatter = &logrus.TextFormatter{
			TimestampFormat: TimeFmtLog,
		}
		LOG.SetLevel(logrus.DebugLevel)
	}
	InitNoFile()
	return initLogDir(linuxPath, fileName, day, isInProduct)
}

func InitNoFile() {
	//添加监听Hook
	LOG.Hooks.Add(new(PrintHook))
}

func CheckDay() {
	_ = initLogFile()
}

func wrapFormat(format, file string, line int) string {
	formatNew := fmt.Sprintf("%s:%d %s", file, line, format)
	return formatNew
}

func D(format string, args ...interface{}) {
	//0的话获取的是129行调用，我们要获取外层调用的位置
	_, file, line, ok := runtime.Caller(1)
	if ok {
		LOG.Debugf(wrapFormat(format, file, line), args...)
	}
}

func I(format string, args ...interface{}) {
	//0的话获取的是129行调用，我们要获取外层调用的位置
	_, file, line, ok := runtime.Caller(1)
	if ok {
		LOG.Infof(wrapFormat(format, file, line), args...)
	}
}

func W(format string, args ...interface{}) {
	//0的话获取的是129行调用，我们要获取外层调用的位置
	_, file, line, ok := runtime.Caller(1)
	if ok {
		LOG.Warningf(wrapFormat(format, file, line), args...)
	}
}

func E(format string, args ...interface{}) {
	//0的话获取的是129行调用，我们要获取外层调用的位置
	_, file, line, ok := runtime.Caller(1) //funcName
	//fmt.Println("Func Name=" + runtime.FuncForPC(funcName).Name())
	if ok {
		LOG.Errorf(wrapFormat(format, file, line), args...)
	}
}
