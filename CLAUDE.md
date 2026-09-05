# Go 学习计划（go_learning_plan）

目标：深入 Go 的并发模型与语言哲学，最终能熟练用于 **SFU 控制面/信令服务**（配合绿联入职：mediasoup + Go 控制面）。

## 教学方式（会话协议）

每次学习按这个循环推进：

```
1. 用户提问 或 发表自己的理解
2. Claude 追踪 Go 运行时源码验证（$(go env GOROOT)/src/runtime/）
3. 确认或修正用户的理解（卡住时先跳出代码，用比喻/心智模型讲透）
4. 手写小程序验证（跑通 / 看输出 / 看崩溃）
5. 沉淀为网页课程（courses/ 下的 HTML）
```

核心原则：
- **源码优先**：不凭记忆答，追踪 `chan.go`、`proc.go`、`runtime2.go` 等真实实现
- **先概念后代码**：用户卡住时，先跳出代码用流水线/比喻讲透，再回源码
- **代码验证**：每个概念都写小程序验证（跑通/看输出/看崩溃），不空谈
- **站点沉淀**：每个主题沉淀成 `chapters/*.md`，由 Go 学习站点（main.go）托管渲染，形成可回顾的资料库
- **例程验证姿势（2026-09-05 定）**：新小节的核心结论**默认写成 `verify/verify_test.go` 里的 `TestXxx` 断言**（`go test ./verify/` 或 `go test ./...` 一键回归，无 main、无多 main 红波浪）。只有"边改边看输出"的探索性调试才开独立 `func main`，且**一个文件夹最多一个 main**（一个目录 = 一个包，多 main 会让 gopls 报 redeclared 红波浪）

## 站点驱动（学习即实战）

学习 Go 的方式 = **用 Go 搭建并不断完善学习站点本身**（`go_learning_plan/`：main.go + chapters/*.md + static/ + courses/layout.html）。

飞轮：学一个 Go 概念 → 沉淀成 `chapters/*.md` → 顺手用新学的特性改进站点代码 → 站点更完善 → 学下一个。

每个 Go 概念有双重价值：既深化理解，又能立刻在站点真实代码里用上（例：学 `html/template` → 做站点的服务端渲染；学 goroutine → 给站点加并发处理）。**站点是实验室，也是资料库。**

## 最终目标：通用知识库框架

从"单课程 Go 站点"升级为**通用框架**：放一个课程 zip（含 `course.yaml` + `chapters/*.md`）→ 自动生成一个课程页面。

```
courses/
  Go 学习课程.zip   → 生成 Go 课程页
  SFU.zip           → 生成 SFU 内容页
```

- 框架用 `archive/zip` 读 zip、`go-yaml` 解析元数据、`net/http` 多课程路由
- Go 学习过程 = 不断完善这个框架；每个新学的 Go 概念（并发/缓存/模板/上传）都沉淀到框架
- 课程内容（chapters/*.md）随时补充，形成个人技术知识库
- 最终部署云端，作为简历作品

## 已完成（2026-08-28）

| 主题 | 要点 | 验证 |
|------|------|------|
| channel 内部实现（send/recv） | send 先查 recvq、recv 先查 sendq，再查缓冲区；等待者优先直接交接 | 源码 + 小程序 |
| 互斥 vs 同步 | 互斥管"同时"、同步管"先后+可见"；happens-before；acquire/release | 源码 + 心智模型 |
| mutex vs channel 取舍 | 传数据用 channel、护状态用 mutex；决策表 | 决策表 |
| 单 goroutine select 同 channel | 无缓冲死锁；有缓冲部分满随机读写 | 小程序验证 |
| 关闭 channel 内部 | 翻 closed 标志 → 接收者拿零值(ok=false) → 发送者 panic → 全部唤醒 | 源码 + 流水线比喻 |

## 学习路线（章节化，由浅入深）

> **推进方式：按章节顺序，从第一章开始**。已掌握的章快速验证（问答过一遍，卡住才深挖）；已学项（3.5 chan、4.3 select、7.1 互斥vs同步、7.6 取舍）到章时当复习。目标从头到尾线性走完，不留夹生盲区。

### 第一章：语言基础（语法与常用数据结构）
- 1.1 变量 / 常量 / iota（枚举技巧）
- 1.2 string（用法）
- 1.3 struct 与方法
- 1.4 slice（用法）
- 1.5 map（用法）
- 1.6 控制流（if / for / switch）
- 1.7 函数（多返回值、闭包）

### 第二章：语法糖与编译还原
- 2.1 `:=` 类型推断 / 短变量声明
- 2.2 `...` 变参与切片展开
- 2.3 多返回值
- 2.4 `for range` 语法糖（array/slice/map/chan/string/int）
- 2.5 `make` vs `new`
- 2.6 复合字面量（struct/map/array 构造糖）
- 2.7 指针自动解引用 / 自动取地址
- 2.8 方法值 / 方法表达式
- 2.9 无条件的 switch（等价 if-else 链）
- 2.10 `go` / `defer` / `select` / `range` 的编译还原

### 第三章：核心数据结构实现原理
- 3.1 string 内部（不可变、字符串头部、[]byte 转换）
- 3.2 slice 内部（array/len/cap、扩容、共享底层数组的坑）
- 3.3 map 内部（hmap/bmap 桶、扩容、并发读写 panic）
- 3.4 struct 内部（内存布局、字段对齐、内嵌）
- 3.5 chan 内部（hchan、send/recv/close）✅ 已学

### 第四章：控制流实现原理
- 4.1 defer（注册与执行、参数求值时机）
- 4.2 range（对 slice/map/chan 遍历的实现）
- 4.3 select（多路复用、随机性、非阻塞 default）✅ 部分已学
- 4.4 panic / recover（栈展开）

### 第五章：协程与调度（Go 的招牌）
- 5.1 goroutine 模型（轻量、栈初始 2KB）
- 5.2 GMP 调度器（proc.go）
- 5.3 栈增长与抢占（asyncPreempt）
- 5.4 goroutine 泄漏与排查

### 第六章：定时器
- 6.1 time.Timer / Ticker / After / Sleep 用法
- 6.2 运行时 timer 实现（per-P 最小堆）
- 6.3 定时器与调度器（checkTimers / findrunnable 里跑 timer）
- 6.4 timer 唤醒与 goroutine 调度
- 6.5 高性能注意点（Ticker 泄漏、定时器堆）
- 6.6 手写验证（Sleep vs Timer、精度）

### 第七章：并发控制（同步原语）
- 7.1 互斥 vs 同步（happens-before）✅ 已学
- 7.2 sync.Mutex 内部（自旋、信号量、饥饿模式）
- 7.3 sync.RWMutex
- 7.4 sync.WaitGroup / Once / Cond
- 7.5 atomic 与内存序
- 7.6 channel vs mutex 取舍 ✅ 已学
- 7.7 context 取消传播
- 7.8 Go 内存模型完整规则

### 第八章：类型系统与接口
- 8.1 interface 隐式实现
- 8.2 空接口 interface{} 与类型断言
- 8.3 泛型（Go 1.18+）
- 8.4 组合 vs 继承

### 第九章：内存管理
- 9.1 逃逸分析
- 9.2 内存分配（mcache/span 缓存层级）
- 9.3 GC（三色标记、写屏障、STW）
- 9.4 性能：pprof、benchmark、零分配优化

### 第十章：工程实践
- 10.1 错误处理（error as value、errors.Is/As、panic 边界）
- 10.2 go test / benchmark / fuzz
- 10.3 反射（reflect：Type/Value、tag）
- 10.4 标准库精读（net/http、io.Reader/Writer、time）
- 10.5 实战：SFU 信令控制面（每连接 goroutine + channel + mutex + context 超时）

## 网页课程

- 每个主题完成后，在 `courses/` 下生成一个 HTML 网页课程
- 网页结构：**概念讲解**（比喻/心智模型）→ **源码追踪**（关键代码 + 注释）→ **手写验证**（可跑代码）→ **小结**
- 网页自包含（内联 CSS/JS），便于分享和复习

## 参考资料

- Go 运行时源码：`$(go env GOROOT)/src/runtime/`（chan.go、proc.go、runtime2.go 等）
- 官方：The Go Memory Model、Effective Go、Go 官方博客
- 关联项目：`/home/ydqun/workspace/lession/xrtc-server2.0`（SFU 学习，绿联 mediasoup + Go 控制面实战前置）
