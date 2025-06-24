package slogutils

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"testing"
)

func TestUseText(t *testing.T) {

	UseText(slog.LevelDebug)

	// 打印更漂亮的对象
	slog.Debug("hello", "name", "world")
	slog.Info("hello", "name", "world")
	slog.Warn("hello", "name", "world")
	slog.Error("hello", "name", "world")

	// 使用组，可以在组中添加多个字段
	groupLog := slog.Default().WithGroup("groupData")
	groupLog.Info("hello", "name", "world", "age", 18)

	// 常规使用
	slog.Info("Request processed",
		slog.String("url", "https://example.com"),
		slog.String("method", "GET"),
		slog.Int("response-code", 200),
	)

	// 使用 WithFields 方法添加多个字段
	slog.Info("Request processed", WithFields(Fields{
		"url":           "URL_ADDRESS",
		"method":        "GET",
		"response-code": 200,
	})...)

}

func TestUseJson(t *testing.T) {

	UseJson(slog.LevelDebug)

	// 打印更漂亮的对象
	slog.Debug("hello", "name", "world")
	slog.Info("hello", "name", "world")
	slog.Warn("hello", "name", "world")
	slog.Error("hello", "name", "world")

	// 使用组，可以在组中添加多个字段
	groupLog := slog.Default().WithGroup("groupData")
	groupLog.Info("hello", "name", "world", "age", 18)

	// 常规使用
	slog.Info("Request processed",
		slog.String("url", "https://example.com"),
		slog.String("method", "GET"),
		slog.Int("response-code", 200),
	)

	// 使用 WithFields 方法添加多个字段
	slog.Info("Request processed", WithFields(Fields{
		"url":           "URL_ADDRESS",
		"method":        "GET",
		"response-code": 200,
	})...)
}

func TestUseTextWithFile(t *testing.T) {

	UseTextWithFile(slog.LevelDebug, "./slog", "log.log")

	// 打印更漂亮的对象
	slog.Debug("hello", "name", "world")
	slog.Info("hello", "name", "world")
	slog.Warn("hello", "name", "world")
	slog.Error("hello", "name", "world")

	// ...... 其他日志代码 ......

}

func TestUseJsonWithFile(t *testing.T) {

	UseJsonWithFile(slog.LevelDebug, "./slog", "log.log")

	// 打印更漂亮的对象
	slog.Debug("hello", "name", "world")
	slog.Info("hello", "name", "world")
	slog.Warn("hello", "name", "world")
	slog.Error("hello", "name", "world")

	// ...... 其他日志代码 ......
}

func TestDiy(t *testing.T) {
	initSLog() // 控制台|文本输出 （方式1）
	//initSlogFile() // 文件|文本输出 （方式2）

	// 打印更漂亮的对象
	slog.Debug("hello", "name", "world")
	slog.Info("hello", "name", "world")
	slog.Warn("hello", "name", "world")
	slog.Error("hello", "name", "world")

}

func initSLog() {
	writerFunc := func() io.Writer {
		return os.Stdout
	}

	handlerFunc := func(w io.Writer) slog.Handler {
		return slog.NewTextHandler(w, &slog.HandlerOptions{
			AddSource:   false,
			Level:       slog.LevelDebug,
			ReplaceAttr: nil,
		})
	}

	Diy(writerFunc, handlerFunc)
}

func initSlogFile() {

	f, err := os.OpenFile(filepath.Join("foo.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Errorf("error opening file: %v", err))
	}

	writerFunc := func() io.Writer {
		return f
	}

	handlerFunc := func(w io.Writer) slog.Handler {
		return slog.NewTextHandler(w, &slog.HandlerOptions{
			AddSource:   false,
			Level:       slog.LevelDebug,
			ReplaceAttr: nil,
		})
	}

	Diy(writerFunc, handlerFunc)
}
