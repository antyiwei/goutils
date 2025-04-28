package timeutils

import (
	"testing"
	"time"
)

func TestGetTimeDifferenceMs(t *testing.T) {
	// 测试用例 1: 时间差为正
	now := time.Now()
	past := now.Add(-100 * time.Millisecond)
	diff := GetTimeDifferenceMs(now, past)
	if diff != 100 {
		t.Errorf("期望时间差为 100 毫秒，实际得到 %d 毫秒", diff)
	}

	// 测试用例 2: 时间差为负
	future := now.Add(50 * time.Millisecond)
	diff = GetTimeDifferenceMs(now, future)
	if diff != -50 {
		t.Errorf("期望时间差为 -50 毫秒，实际得到 %d 毫秒", diff)
	}

	// 测试用例 3: 时间差为零
	diff = GetTimeDifferenceMs(now, now)
	if diff != 0 {
		t.Errorf("期望时间差为 0 毫秒，实际得到 %d 毫秒", diff)
	}
}
