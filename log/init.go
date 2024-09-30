package log

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

// init Logger
func init() {
	Logger = logrus.New()
	Logger.Infoln()
	Logger.SetFormatter(&LiteyukiFormatter{})
}

const (
	TraceLevel = logrus.TraceLevel
	DebugLevel = logrus.DebugLevel
	InfoLevel  = logrus.InfoLevel
	WarnLevel  = logrus.WarnLevel
	ErrorLevel = logrus.ErrorLevel
	FatalLevel = logrus.FatalLevel
	PanicLevel = logrus.PanicLevel
)

const (
	Reset   = "\033[0m"
	Black   = "\033[30m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"

	BlackBold   = "\033[30;1m"
	RedBold     = "\033[31;1m"
	GreenBold   = "\033[32;1m"
	YellowBold  = "\033[33;1m"
	BlueBold    = "\033[34;1m"
	MagentaBold = "\033[35;1m"
	CyanBold    = "\033[36;1m"
	WhiteBold   = "\033[37;1m"
)

var Logger *logrus.Logger

// LiteyukiFormatter 标准输出格式化
type LiteyukiFormatter struct{}

// Format 标准输出格式
func (f *LiteyukiFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b bytes.Buffer
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	// 时间为绿色
	timeText := Color(Green, timestamp)
	logLevel := entry.Level.String()
	message := entry.Message
	// 定义颜色
	var color string
	switch entry.Level {
	case logrus.TraceLevel:
		color = Black
	case logrus.DebugLevel:
		color = Cyan
	case logrus.InfoLevel:
		color = Blue
	case logrus.WarnLevel:
		color = Yellow
	case logrus.ErrorLevel:
		color = Red
	case logrus.FatalLevel:
		color = RedBold
	case logrus.PanicLevel:
		color = RedBold
	default:
		color = "\033[0m"
	}
	levelText := Color(color, "["+logLevel+"]")
	// 格式化日志信息并添加颜色
	b.WriteString(fmt.Sprintf("%s %s %s\n", timeText, levelText, message))
	return b.Bytes(), nil
}

func Trace(args ...interface{}) {
	Logger.Trace(args)
}

func Traceln(args ...interface{}) {
	Logger.Traceln(args)
}

func Tracef(format string, args ...interface{}) {
	Logger.Tracef(format, args)
}

func Debug(args ...interface{}) {
	Logger.Debug(args)
}

func Debugln(args ...interface{}) {
	Logger.Debugln(args)
}

func Debugf(format string, args ...interface{}) {
	Logger.Debugf(format, args)
}

func Info(args ...interface{}) {
	Logger.Info(args)
}

func Infoln(args ...interface{}) {
	Logger.Infoln(args)
}

func Infof(format string, args ...interface{}) {
	Logger.Infof(format, args)
}

func Warn(args ...interface{}) {
	Logger.Warn(args)
}

func Warnln(args ...interface{}) {
	Logger.Warnln(args)
}

func Warnf(format string, args ...interface{}) {
	Logger.Warnf(format, args)
}

func Error(args ...interface{}) {
	Logger.Error(args)

}

func Errorln(args ...interface{}) {
	Logger.Errorln(args)
}

func Errorf(format string, args ...interface{}) {
	Logger.Errorf(format, args)
}

func Fatal(args ...interface{}) {
	Logger.Fatal(args)
}

func Fatalln(args ...interface{}) {
	Logger.Fatalln(args)
}

func Fatalf(format string, args ...interface{}) {
	Logger.Fatalf(format, args)
}

func Panic(args ...interface{}) {
	Logger.Panic(args)
}

func Panicln(args ...interface{}) {
	Logger.Panicln(args)
}

func Panicf(format string, args ...interface{}) {
	Logger.Panicf(format, args)
}

// Color 字符串添加颜色
func Color(color, text string) string {
	return fmt.Sprintf("%s%s%s", color, text, Reset)
}

// SetLevel 设置日志级别
func SetLevel(level logrus.Level) {
	Logger.SetLevel(level)
}

// SetFormatter 设置日志格式
func SetFormatter(formatter logrus.Formatter) {
	Logger.SetFormatter(formatter)
}

// SetOutput 设置日志输出
func SetOutput(out *bytes.Buffer) {
	Logger.SetOutput(out)
}

// getCurrentDate 获取当前日期, 返回格式为yyyy-mm-dd
func getCurrentDate() string {
	return time.Now().Format("2006-01-02")
}
