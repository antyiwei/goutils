package rlockutils

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
)

// SimpleLock 极简Redis分布式锁工具
// 提供最基础的锁功能，易于使用

// Lock 获取锁并执行函数（推荐使用）
// client: Redis客户端
// key: 锁的键名
// fn: 需要在锁保护下执行的函数
// 默认：锁过期时间30秒，获取锁超时10秒
func Lock(client *redis.Client, key string, fn func() error) error {
	return LockWithTimeout(client, key, 30*time.Second, 10*time.Second, fn)
}

// LockWithTimeout 获取锁并执行函数（自定义超时）
// client: Redis客户端
// key: 锁的键名
// expiry: 锁过期时间
// timeout: 获取锁的超时时间
// fn: 需要在锁保护下执行的函数
func LockWithTimeout(client *redis.Client, key string, expiry, timeout time.Duration, fn func() error) error {
	// 创建redsync实例
	pool := goredis.NewPool(client)
	rs := redsync.New(pool)
	mutex := rs.NewMutex(key, redsync.WithExpiry(expiry))

	// 获取锁（阻塞式，带超时）
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err := mutex.LockContext(ctx)
	if err != nil {
		return fmt.Errorf("获取锁失败: %w", err)
	}

	// 启动自动续期
	renewCtx, renewCancel := context.WithCancel(context.Background())
	go func() {
		ticker := time.NewTicker(expiry / 3)
		defer ticker.Stop()
		for {
			select {
			case <-renewCtx.Done():
				return
			case <-ticker.C:
				if ok, _ := mutex.Extend(); !ok {
					return
				}
			}
		}
	}()

	// 确保释放锁
	defer func() {
		renewCancel() // 停止续期
		if ok, unlockErr := mutex.Unlock(); !ok || unlockErr != nil {
			// 静默处理解锁错误，避免影响业务逻辑
		}
	}()

	// 执行业务函数
	return fn()
}

// TryLock 尝试获取锁并执行函数（非阻塞）
// client: Redis客户端
// key: 锁的键名
// fn: 需要在锁保护下执行的函数
// 如果锁已被占用，立即返回错误
func TryLock(client *redis.Client, key string, fn func() error) error {
	// 创建redsync实例
	pool := goredis.NewPool(client)
	rs := redsync.New(pool)
	mutex := rs.NewMutex(key, redsync.WithExpiry(30*time.Second))

	// 尝试获取锁（非阻塞）
	err := mutex.TryLock()
	if err != nil {
		return fmt.Errorf("锁已被占用: %w", err)
	}

	// 启动自动续期
	renewCtx, renewCancel := context.WithCancel(context.Background())
	go func() {
		ticker := time.NewTicker(10 * time.Second) // 每10秒续期一次
		defer ticker.Stop()
		for {
			select {
			case <-renewCtx.Done():
				return
			case <-ticker.C:
				if ok, _ := mutex.Extend(); !ok {
					return
				}
			}
		}
	}()

	// 确保释放锁
	defer func() {
		renewCancel() // 停止续期
		if ok, unlockErr := mutex.Unlock(); !ok || unlockErr != nil {
			// 静默处理解锁错误，避免影响业务逻辑
		}
	}()

	// 执行业务函数
	return fn()
}
