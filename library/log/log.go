package log

import (
	"fmt"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"gosharp/library/config"
	"gosharp/library/file"
	"path"
	"time"
)

var Logger *logrus.Logger

func GetLogger() *logrus.Logger {
	return Logger
}

func Init(logPath string) {
	Logger = logrus.New()
	level, err := logrus.ParseLevel(config.Viper.GetString("LOG_LEVEL"))
	if err != nil {
		level = logrus.DebugLevel
	}
	Logger.SetLevel(level)
	if logPath == "" {
		logPath = "logs"
	}
	ConfigLocalFilesystemLogger(logPath, "", time.Hour*24, time.Hour*24)
}

func ConfigLocalFilesystemLogger(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) {
	err := file.MkdirIfNotExist(logPath)

	baseLogPaht := path.Join(logPath, logFileName)
	writer, err := rotatelogs.New(
		baseLogPaht+"/"+"%Y%m%d"+".log",
		//rotatelogs.WithLinkName(baseLogPaht),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		panic(fmt.Sprintf("config local file system logger error. %+v", errors.WithStack(err)))
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.TextFormatter{DisableColors: true, TimestampFormat: "2006-01-02 15:04:05.000"})

	Logger.AddHook(lfHook)
}
