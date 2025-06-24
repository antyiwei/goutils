package rlockutils

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

// 创建测试用的Redis客户端
func createTestClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       1, // 使用测试数据库
	})
}

// TestEasyLock 测试基础锁功能
func TestEasyLock(t *testing.T) {
	client := createTestClient()
	defer client.Close()

	lockKey := fmt.Sprintf("test:simple:lock:%d", time.Now().Unix())
	executionCount := 0

	err := Lock(client, lockKey, func() error {
		executionCount++
		t.Logf("业务逻辑执行，次数: %d", executionCount)
		time.Sleep(100 * time.Millisecond)
		return nil
	})

	if err != nil {
		t.Fatalf("EasyLock执行失败: %v", err)
	}

	if executionCount != 1 {
		t.Fatalf("执行次数错误: 期望1, 实际%d", executionCount)
	}
}

// TestEasyTryLock 测试非阻塞锁功能
func TestEasyTryLock(t *testing.T) {
	client := createTestClient()
	defer client.Close()

	lockKey := fmt.Sprintf("test:simple:trylock:%d", time.Now().Unix())
	executionCount := 0

	err := TryLock(client, lockKey, func() error {
		executionCount++
		t.Logf("EasyTryLock业务逻辑执行，次数: %d", executionCount)
		time.Sleep(100 * time.Millisecond)
		return nil
	})

	if err != nil {
		t.Fatalf("EasyTryLock执行失败: %v", err)
	}

	if executionCount != 1 {
		t.Fatalf("执行次数错误: 期望1, 实际%d", executionCount)
	}
}

// Test_LockWithTimeout 测试自定义超时锁功能
func Test_LockWithTimeout(t *testing.T) {
	client := createTestClient()
	defer client.Close()

	lockKey := fmt.Sprintf("test:simple:timeout:%d", time.Now().Unix())
	executionCount := 0

	err := LockWithTimeout(client, lockKey, 1*time.Minute, 5*time.Second, func() error {
		executionCount++
		t.Logf("EasyLockWithTimeout业务逻辑执行，次数: %d", executionCount)
		time.Sleep(200 * time.Second)
		return nil
	})

	if err != nil {
		t.Fatalf("EasyLockWithTimeout执行失败: %v", err)
	}

	if executionCount != 1 {
		t.Fatalf("执行次数错误: 期望1, 实际%d", executionCount)
	}
}

// TestEasyConcurrentLock 测试并发锁竞争
func TestEasyConcurrentLock(t *testing.T) {
	client := createTestClient()
	defer client.Close()

	lockKey := fmt.Sprintf("test:simple:concurrent:%d", time.Now().Unix())
	successCount := 0
	failCount := 0
	var mu sync.Mutex

	// 启动多个goroutine竞争同一个锁
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			err := TryLock(client, lockKey, func() error {
				t.Logf("Goroutine %d 获得锁", id)
				time.Sleep(35 * time.Second)
				return nil
			})

			mu.Lock()
			if err != nil {
				failCount++
				t.Logf("Goroutine %d 获取锁失败: %v", id, err)
			} else {
				successCount++
				t.Logf("Goroutine %d 执行成功", id)
			}
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	t.Logf("并发测试完成: 成功 %d 次, 失败 %d 次", successCount, failCount)

	// 至少应该有一个成功
	if successCount == 0 {
		t.Fatal("所有并发请求都失败了")
	}

	// 由于是互斥锁，成功次数应该只有1个
	if successCount != 1 {
		t.Logf("警告: 成功次数为 %d，期望为 1", successCount)
	}
}

// TestEasyLockBusinessError 测试业务逻辑错误处理
func TestEasyLockBusinessError(t *testing.T) {
	client := createTestClient()
	defer client.Close()

	lockKey := fmt.Sprintf("test:simple:error:%d", time.Now().Unix())

	err := Lock(client, lockKey, func() error {
		t.Log("模拟业务逻辑错误")
		time.Sleep(15 * time.Second)
		panic("业务逻辑错误")
		return fmt.Errorf("业务逻辑执行失败")
	})

	if err == nil {
		t.Fatal("期望返回错误，但没有错误")
	}

	if err.Error() != "业务逻辑执行失败" {
		t.Fatalf("错误信息不匹配: %v", err)
	}

	t.Logf("正确处理业务逻辑错误: %v", err)
}
