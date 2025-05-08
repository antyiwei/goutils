package slogutils

import (
	"io"
	"log"
	"log/slog"
	"os"
	"path"

	"gopkg.in/natefinch/lumberjack.v2"
)

// UseText 文本日志输出格式(控制台)
func UseText(level slog.Level) {

	textHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource:   false,
		Level:       level,
		ReplaceAttr: nil,
	})

	slog.SetDefault(slog.New(textHandler))

}

// UseJson json日志输出格式(控制台)
func UseJson(level slog.Level) {

	jsonHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: false,
		Level:     level,
	})

	slog.SetDefault(slog.New(jsonHandler))

}

type Fields map[string]interface{}

func WithFields(fields Fields) []any {
	var attrs []any
	for k, v := range fields {
		attrs = append(attrs, slog.Any(k, v))
	}
	return attrs
}

// UseTextWithFile 文本日志输出格式(文件)
func UseTextWithFile(level slog.Level, dir, logFileName string) {

	// 可定制的输出目录。
	var logFilePath string
	if dir == "" {
		dir = "./slog"
	}

	logFilePath = dir + "/logs/"
	if err := os.MkdirAll(logFilePath, 0o777); err != nil {
		log.Println(err.Error())
		return
	}

	// 将文件名设置为日期
	if logFileName == "" {
		logFileName = "log.log"
	}
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			log.Println(err.Error())
			return
		}
	}

	w := &lumberjack.Logger{
		Filename:   fileName,
		LocalTime:  true,
		MaxSize:    50,    // 日志的最大大小，以M为单位
		MaxBackups: 5,     // 保留的旧日志文件的最大数量
		MaxAge:     28,    // 保留旧日志文件的最大天数
		Compress:   false, // 是否压缩旧日志文件
	}

	textHandler := slog.NewTextHandler(w, &slog.HandlerOptions{
		AddSource:   false,
		Level:       level,
		ReplaceAttr: nil,
	})

	slog.SetDefault(slog.New(textHandler))

}

// UseJsonWithFile json日志输出格式(文件)
func UseJsonWithFile(level slog.Level, dir, logFileName string) {

	// 可定制的输出目录。
	var logFilePath string
	if dir == "" {
		dir = "./slog"
	}

	logFilePath = dir + "/logs/"
	if err := os.MkdirAll(logFilePath, 0o777); err != nil {
		log.Println(err.Error())
		return
	}

	// 将文件名设置为日期
	if logFileName == "" {
		logFileName = "log.log"
	}
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			log.Println(err.Error())
			return
		}
	}

	w := &lumberjack.Logger{
		Filename:   fileName,
		LocalTime:  true,
		MaxSize:    50,    // 日志的最大大小，以M为单位
		MaxBackups: 5,     // 保留的旧日志文件的最大数量
		MaxAge:     28,    // 保留旧日志文件的最大天数
		Compress:   false, // 是否压缩旧日志文件
	}

	textHandler := slog.NewJSONHandler(w, &slog.HandlerOptions{
		AddSource:   false,
		Level:       level,
		ReplaceAttr: nil,
	})

	slog.SetDefault(slog.New(textHandler))

}

// Diy 自定义日志输出格式和输出位置
func Diy(writerFunc func() io.Writer, handlerFunc func(io.Writer) slog.Handler) {
	slog.SetDefault(slog.New(handlerFunc(writerFunc())))
}

func DiySpecial(writerFunc func() io.Writer, handlerFunc func(io.Writer) slog.Handler) *slog.Logger {
	return slog.New(handlerFunc(writerFunc()))
}
