# lockutils 包说明

`lockutils` 包提供了一个简单的尝试锁（try lock）实现，用于在并发编程中控制对共享资源的访问。

## 结构体

### `Lock`

`Lock` 结构体表示一个尝试锁，内部使用一个容量为 1 的通道来实现。

```go
type Lock struct {
    c chan struct{}
}
```
