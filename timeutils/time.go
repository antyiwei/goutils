package timeutils

import (
	"time"
)

// GetCurrentTimestamp 获取当前时间戳（秒）
func GetCurrentTimestamp() int64 {
	return time.Now().Unix()
}

// GetCurrentTimestampMs 获取当前时间戳（毫秒）
func GetCurrentTimestampMs() int64 {
	return time.Now().UnixMilli()
}

// FormatTime 格式化时间
func FormatTime(t time.Time, layout string) string {
	return t.Format(layout)
}

// ParseTime 解析时间字符串
func ParseTime(timeStr string, layout string) (time.Time, error) {
	return time.Parse(layout, timeStr)
}

// GetTimeDifference 获取两个时间的差值（秒）
func GetTimeDifference(t1, t2 time.Time) float64 {
	return t1.Sub(t2).Seconds()
}

// GetTimeDifferenceMs 获取两个时间的差值（毫秒）
func GetTimeDifferenceMs(t1, t2 time.Time) int64 {
	return t1.Sub(t2).Milliseconds()
}

// AddSeconds 在给定时间上增加指定秒数
func AddSeconds(t time.Time, seconds int64) time.Time {
	return t.Add(time.Duration(seconds) * time.Second)
}

// AddMilliseconds 在给定时间上增加指定毫秒数
func AddMilliseconds(t time.Time, milliseconds int64) time.Time {
	return t.Add(time.Duration(milliseconds) * time.Millisecond)
}
