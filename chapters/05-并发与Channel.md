# 并发与 Channel

> 一句话心智模型：**channel 是一条数据流水线，生产者放货、消费者取货；关闭 channel = 广播"没有更多数据了"。**

## 一、channel 基础语法

```go
ch := make(chan int)      // 无缓冲：发送必须等接收者就绪（或反之），否则阻塞
ch := make(chan int, 4)   // 有缓冲：缓冲有空间就能发，有数据就能收
ch <- v                   // 发送
<-ch                      // 接收
v, ok := <-ch             // ok=false 表示通道已关闭、拿到的是零值
close(ch)                 // 关闭（只有发送方能 close）
```

## 二、channel 内部实现（hchan）

`hchan` 结构（`runtime/chan.go`）：一个环形缓冲 + 两个等待队列 + 一把锁。

```
hchan {
  buf      环形缓冲（dataqsiz / qcount / sendx / recvx）
  sendq    阻塞的发送者队列（waitq）
  recvq    阻塞的接收者队列（waitq）
  closed   是否已关闭
  lock     mutex（保护一切）
}
```

### 发送（chansend）：先查 recvq，再查缓冲

```
lock
① recvq 非空?（有阻塞的接收者）
     → 有：直接交接给接收者，唤醒它，完成
② recvq 空 → 缓冲有空位?（qcount < dataqsiz）
     → 有：放进 buffer，完成
③ 都没有 → 阻塞：把自己加进 sendq，park
unlock
```

### 接收（chanrecv）：先查 sendq，再查缓冲

```
lock
① sendq 非空?（有阻塞的发送者）
     → 有：交接（无缓冲直接拿；有缓冲从 buffer 取 + 把发送者的值挪进空位），唤醒它
② sendq 空 → 缓冲有数据?（qcount > 0）
     → 有：从 buffer 取一个
③ 都没有 → 阻塞：把自己加进 recvq，park
unlock
```

**为什么先查等待队列**：直接交接（handoff）比走缓冲快——省一次"写缓冲 + 读缓冲"的拷贝，还能立刻唤醒等待者。**有等待者就优先直接递，缓冲区只是"暂时没对接上"时的暂存区。**

### 验证：单 goroutine 用 select 对同一个 channel 读写

```go
// 无缓冲 → 死锁 panic
ch := make(chan int)
select {
case v := <-ch:  // 需要"有阻塞的发送者"才 ready
case ch <- 1:    // 需要"有阻塞的接收者"才 ready
}
// 唯一的 goroutine 是你自己 → 两边都等不到对方 → all goroutines asleep - deadlock!

// 有缓冲、部分满(0<len<cap) → 读写都 ready，随机选一个
ch := make(chan int, 4)
ch <- 1; ch <- 2
for i := 0; i < 12; i++ {
    select {
    case v := <-ch: fmt.Println("read", v)
    case ch <- 9:  fmt.Println("write")
    }
}
// 输出 read/write 交替无规律（非确定）
```

**结论**：单 goroutine 对同一 channel 读写几乎总是 bug——无缓冲死锁，有缓冲部分满随机乱流。channel 的发送方和接收方应该是两个不同的执行体。

## 三、关闭 channel 的内部（closechan）

```
closechan:
  ① c == nil?            → panic("close of nil channel")
  ② c.closed != 0?       → panic("close of closed channel")   // 双重关闭检测
  ③ c.closed = 1          // 翻转标志（在锁内，原子）
  ④ 释放 recvq：接收者的目标内存清零（typedmemclr）→ 接收者拿到零值
  ⑤ 释放 sendq：发送者的值作废（elem=nil）→ 发送者醒来 panic
  ⑥ 解锁 → 唤醒所有被释放的 goroutine（goready，在锁外）
```

**不对称待遇**（Go 的设计决策）：

| 等待者 | 唤醒后 | 语义 |
|--------|--------|------|
| **接收者**（recvq） | 拿到**零值**（`ok=false`） | **接收永远安全**——关闭后还能读（缓冲先读光，再拿零值），`for range` 靠它退出 |
| **发送者**（sendq） | **panic**（"send on closed channel"） | **发送永远危险**——值注定丢失是 bug，要大声失败 |

**经典用法**：只有发送方 close；`for range` 读到 ok=false 自动退出；close 作为"通知所有等待者"的信号。

## 四、互斥 vs 同步（happens-before）

- **互斥**（Mutual Exclusion）管"**同时**"：同一时刻最多一个执行体进临界区
- **同步**（Synchronization）管"**先后 + 可见**"：建立 happens-before，让一个执行体的写对另一个可见且有序
- **互斥 ⊂ 同步**：互斥是同步的一种

**为什么锁/原子"也有同步语义"**：它们不只防踩踏，还建立 happens-before：

```go
var x int
var mu sync.Mutex
// goroutine A
x = 42        // 写
mu.Unlock()   // 释放
// goroutine B
mu.Lock()     // 获取
_ = x         // 保证读到 42
```

`Unlock → Lock` 构成 happens-before 边：**A 在 Unlock 前写的一切，B 在 Lock 后保证看见**。这不是概率、不是"谁先跑"，是 acquire/release 配对的结构保证。

**统一内核**：所有并发原语都在做同一件事——建立 happens-before。

| 原语 | 释放端 | 获取端 | happens-before |
|------|--------|--------|----------------|
| mutex | `Unlock` | `Lock` | Unlock → Lock |
| channel | `send` | `receive` | send → receive |
| atomic | `Store` | `Load` | Store → Load |
| WaitGroup | `Done` | `Wait` 返回 | Done → Wait |

> 心智模型：**互斥解决"谁先进门"，同步解决"进门后能不能看见别人留下的东西"。** 锁和原子变量同时干这两件事——一个管"同时"，一个管"先后 + 可见"。

## 五、mutex vs channel 取舍

一句话规则：**传数据用 channel，护状态用 mutex。**

| 场景 | 用啥 |
|------|------|
| 把数据从一个 goroutine 交给另一个 | **channel** |
| 生产者-消费者 / 流水线 / 任务分发 | **channel** |
| 等待多个事件 / 取消 / 超时 | **channel + select** |
| 一对多广播 | **channel** |
| 天然背压 / 流控（满了自动阻塞） | **channel**（有缓冲） |
| 多个 goroutine 共享 map/slice/计数器 | **mutex** |
| 读多写少 | **RWMutex** |
| 单纯同步（等对方干完，不传数据） | **sync.WaitGroup**（不是 channel） |
| 高频小临界区（计数 +1、check-then-set） | **mutex / atomic** |

**常见误区**："Go 哲学 = 只用 channel" 是误读——官方明确说纯同步用 mutex/WaitGroup，channel 不是 mutex 的替代品。

**实际 SFU 控制面**：每连接一个 goroutine + channel 传信令消息 + mutex 护连接注册表 map + atomic 数活跃流——两者混着用是常态。

## 六、手写验证

```bash
# 无缓冲 select 同 channel → 死锁 panic
go run select_deadlock.go   # fatal error: all goroutines are asleep - deadlock!

# 有缓冲部分满 → 读写随机
go run select_same_channel.go  # read/write 交替无规律
```
