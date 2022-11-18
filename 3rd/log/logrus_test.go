package log

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"io"
	"io/fs"
	"os"
	"testing"
)

// https://zhuanlan.zhihu.com/p/105759117 查看更多用法
/**
日志级别从上向下依次增加，Trace最大，Panic最小。
logrus有一个日志级别，高于这个级别的日志不会输出
。默认的级别为InfoLevel。所以为了能看到Trace和Debug日志，我们在main函数第一行设置日志级别为TraceLevel。
*/

func TestSetOutput(t *testing.T) {
	// path := "./storage/logs/log.log"? // 在项目中./代表当前项目根目录
	path := "../../storage/logs/log.log" //在单元测试中./路径是当前文件路径
	logger := logrus.New()
	writer, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, fs.ModePerm)
	if err != nil {
		logger.Error(err)
	}

	logger.SetReportCaller(true)       //设置在输出日志中添加文件名和方法信息：
	logger.SetOutput(writer)           //设置输出
	logger.SetLevel(logrus.TraceLevel) //设置输出等级

	logger.Trace("test trace msg")
	logger.Debug("test debug msg")
	logger.Info("test info msg")
	logger.Warn("test Warn msg")
	logger.Error("test Error msg")
	logger.Fatal("test Fatal msg") //程序会退出，会导致单元测试不通过
	logger.Panic("test panic  msg")
}

// 多方输出
func TestSetMultiOutput(t *testing.T) {
	writer1 := &bytes.Buffer{}
	writer2 := os.Stdout

	fileName := "../../storage/logs/log.log"
	writer3, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, fs.ModePerm)

	logger := logrus.New()

	if err != nil {
		logger.Fatalf("create file log.txt failed: %v", err)
	}
	logger.SetOutput(io.MultiWriter(writer1, writer2, writer3))
	logger.Info("info msg")
}

// 设置输出格式
func TestSetOutputFormatter(t *testing.T) {
	path := "../../storage/logs/log.log" //在单元测试中./路径是当前文件路径
	logger := logrus.New()
	writer, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, fs.ModePerm)
	if err != nil {
		logger.Error(err)
	}

	logger.SetReportCaller(true)                 //设置在输出日志中添加文件名和方法信息：
	logger.SetOutput(writer)                     //设置输出
	logger.SetLevel(logrus.TraceLevel)           //设置输出等级
	logger.SetFormatter(&logrus.JSONFormatter{}) //设置输出格式，默认TextFormatter
	logger.Trace("test trace msg")
}

//添加字段
//可以通过调用logrus.WithField和logrus.WithFields实现。logrus.WithFields接受一个logrus.Fields类型的参数，其底层实际上为map[string]interface{}：
func TestWithField(t *testing.T) {
	path := "../../storage/logs/log.log" //在单元测试中./路径是当前文件路径
	logger := logrus.New()
	writer, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, fs.ModePerm)
	if err != nil {
		logger.Error(err)
	}

	logger.SetReportCaller(true)                 //设置在输出日志中添加文件名和方法信息：
	logger.SetOutput(writer)                     //设置输出
	logger.SetLevel(logrus.TraceLevel)           //设置输出等级
	logger.SetFormatter(&logrus.JSONFormatter{}) //设置输出格式，默认TextFormatter
	logger.WithFields(logrus.Fields{
		"name": "张三",
		"age":  18,
	}).Info("info msg")

}
