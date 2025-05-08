## 这是一个日志工具包

### 1. 介绍
    使用golang的slog包，封装了一个日志工具包，提供了多种日志输出方式，包括控制台输出、文件输出和远程输出等。

    支持多种日志级别，包括debug、info、warn、error等，方便用户根据需要选择合适的日志级别。

    支持日志格式化输出，包括json格式和文本格式，方便用户根据需要选择合适的日志格式。

    支持日志分割和压缩，方便用户管理日志文件，避免日志文件过大导致的存储问题。

    支持日志轮转，方便用户管理日志文件，避免日志文件过大导致的存储问题。

    支持日志过滤，方便用户根据需要选择合适的日志输出，避免不必要的日志输出导致的存储问题。

    支持日志输出到多个目标，方便用户根据需要选择合适的日志输出目标，避免不必要的日志输出导致的存储问题。

### 2. 安装
    go get github.com/antyiwei/goutils/slogutils

### 3. 使用

#### 3.0 slog的使用
```go
package main

import (
	"log/slog"
	"os"
	
	"github.com/antyiwei/goutils/slogutils"
)

func main() {

	textHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource:   false,
		Level:       slog.LevelDebug,
		ReplaceAttr: nil,
	})

	slog.SetDefault(slog.New(textHandler))
	
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
	slog.Info("Request processed", slogutils.WithFields(slogutils.Fields{
		"url":           "URL_ADDRESS",
		"method":        "GET",
		"response-code": 200,
	})...)
}

// output:
//time=2025-05-08T14:56:30.489+08:00 level=DEBUG msg=hello name=world
//time=2025-05-08T14:56:30.489+08:00 level=INFO msg=hello name=world
//time=2025-05-08T14:56:30.489+08:00 level=WARN msg=hello name=world
//time=2025-05-08T14:56:30.489+08:00 level=ERROR msg=hello name=world
//time=2025-05-08T14:56:30.489+08:00 level=INFO msg=hello groupData.name=world groupData.age=18
//time=2025-05-08T14:59:04.590+08:00 level=INFO msg="Request processed" url=https://example.com method=GET response-code=200
//time=2025-05-08T15:18:16.061+08:00 level=INFO msg="Request processed" response-code=200 url=URL_ADDRESS method=GET

```

    你可以使用下面的代码来初始化日志，使用默认的输出和处理函数。
```go
package main

import (
	"log/slog"
	"os"

	"github.com/antyiwei/goutils/slogutils"
)

func main() {

	jsonHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: false,
		Level:     slog.LevelDebug,
	})

	slog.SetDefault(slog.New(jsonHandler))
	
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
	slog.Info("Request processed", slogutils.WithFields(slogutils.Fields{
		"url":           "URL_ADDRESS",
		"method":        "GET",
		"response-code": 200,
	})...)
}

//{"time":"2025-05-08T15:27:01.478842+08:00","level":"DEBUG","msg":"hello","name":"world"}
//{"time":"2025-05-08T15:27:01.479053+08:00","level":"INFO","msg":"hello","name":"world"}
//{"time":"2025-05-08T15:27:01.479058+08:00","level":"WARN","msg":"hello","name":"world"}
//{"time":"2025-05-08T15:27:01.479061+08:00","level":"ERROR","msg":"hello","name":"world"}
//{"time":"2025-05-08T15:27:01.479066+08:00","level":"INFO","msg":"hello","groupData":{"name":"world","age":18}}
//{"time":"2025-05-08T15:27:01.479072+08:00","level":"INFO","msg":"Request processed","url":"https://example.com","method":"GET","response-code":200}
//{"time":"2025-05-08T15:27:01.479087+08:00","level":"INFO","msg":"Request processed","response-code":200,"url":"URL_ADDRESS","method":"GET"}

```

#### 3.1 WithFields的使用
```go
// ...... 其他代码 ......
slog.Info("Request processed", slogutils.WithFields(Fields{
"url":           "URL_ADDRESS",
"method":        "GET",
"response-code": 200,
})...)
// ...... 其他代码 ......

```
#### 3.2 slogutils 里面的4种常规使用方法
    你可以使用下面的代码来初始化日志，使用默认的输出和处理函数。
```go
package main

import (
	"io"
	"log/slog"
	"os"

	"github.com/antyiwei/goutils/slogutils"
)

func main() {
	// 初始化日志
	slogutils.UseJson(slog.LevelInfo)

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
	slog.Info("Request processed", slogutils.WithFields(slogutils.Fields{
		"url":           "URL_ADDRESS",
		"method":        "GET",
		"response-code": 200,
	})...)
	
	
}

```

#### 3.3 DIY自定义
    你可以使用下面的代码来初始化日志，使用自定义的输出和处理函数。
```go
package main

import (
	"io"
	"log/slog"
	"os"

	"github.com/antyiwei/goutils/slogutils"
)

func main() {
	// 初始化日志
	InitSLog()

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
	slog.Info("Request processed", slogutils.WithFields(slogutils.Fields{
		"url":           "URL_ADDRESS",
		"method":        "GET",
		"response-code": 200,
	})...)
}

func InitSLog() {
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

	slogutils.Diy(writerFunc, handlerFunc)
}

```

### 4. 其他
```text
参考文章：https://blog.csdn.net/holdlg/article/details/130367106

```


```go
/*
    自测代码
 */
	
func DIYLOG() {
	
	// 可定制的输出目录。
	var logFilePath string
	dir := "./slog"
	logFilePath = dir + "/logs/"
	if err := os.MkdirAll(logFilePath, 0o777); err != nil {
		log.Println(err.Error())
		return
	}

	// 将文件名设置为日期
	logFileName := "log.log"
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
		MaxSize:    2,     // 日志的最大大小，以M为单位
		MaxBackups: 3,     // 保留的旧日志文件的最大数量
		MaxAge:     28,    // 保留旧日志文件的最大天数
		Compress:   false, // 是否压缩旧日志文件
	}

	// --参考-- 
	// https://blog.csdn.net/holdlg/article/details/130367106
	
	
	// 输出到控制台
	// log := slog.New(slog.NewTextHandler(os.Stdout))
	// jsonHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
	
	// 输出到文件
	// log := slog.New(slog.NewTextHandler(w))

	// 输出到文件
	// jsonHandler := slog.NewJSONHandler(w, &slog.HandlerOptions{

	// 输出到控制台和文件
	// jsonHandler := slog.NewJSONHandler(io.MultiWriter(os.Stdout, w), &slog.HandlerOptions{

	// 输出到文件
	jsonHandler := slog.NewJSONHandler(w, &slog.HandlerOptions{
		// AddSource: true,
		Level: slog.LevelDebug,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				a.Key = "timestamp" // 自定义时间戳的key: 将time变更为timestamp
			} 
			return a
		},
	})
	
	/*
	    JSLog = slog.New(jsonHandler) 获取到一个全局日志对象。如果不启用 slog.SetDefault(JSLog) ，那么日志会有两个对象。
        JSLog和slog中默认的两个对象。  
	    如果启用 slog.SetDefault(JSLog) ，那么日志会有一个对象。
		如果不启用 slog.SetDefault(JSLog) ，那么日志会有两个对象。
	
	    两个对象时候，可以JSLog打印日志到文件，slog打印到控制台。
	
	    可以灵活使用。
	
	    slogutils工具包中，都是使用一个对象。
	    如果你需要使用两个对象，DiySpecial函数，自己定义全局变量接收。然后项目中，使用全局变量打印。
	*/
	JSLog = slog.New(jsonHandler)
	slog.SetDefault(JSLog)
}
```