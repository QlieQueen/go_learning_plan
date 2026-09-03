---
name: "go-learning"
description: "深入 Go 并发与语言哲学：用户提问/发表理解 → 追踪 Go 运行时源码验证 → 完善理解 → 手写代码验证 → 沉淀网页课程（courses/）"
---

# Go 学习（go-learning）

按 `/home/ydqun/workspace/lession/go_learning_plan/CLAUDE.md` 的教学协议推进 Go 学习。

## 教学循环

```
用户提问/发表理解
  → 源码验证 (go env GOROOT/src/runtime)
  → 确认或修正（卡住先跳代码、用比喻讲透）
  → 手写小程序验证
  → 沉淀 HTML 网页课程到 courses/
```

## 进度（章节化）

- ✅ 第三章 3.5 chan 内部（send/recv 等待队列优先 + 缓冲、close 释放逻辑）
- ✅ 第七章 7.1 互斥 vs 同步、happens-before、acquire/release
- ✅ 第七章 7.6 channel vs mutex 取舍
- ✅ 第四章 4.3 select 部分（单 goroutine 同 channel：死锁/随机）
- ✅ 第一章 安装与环境搭建（直接补充成文，未走学习循环——已有环境）
- ✅ 第二章 Part1-5 已学已沉淀：变量/常量/iota、string（rune=编号 string=编码、UTF-8 自同步）、struct 与方法、slice（三字段/窗口共享/扩容规则）、map（nil vs 空、零值陷阱、类型名vs复合字面量）
- ✅ 第二章 Part6 控制流已学已沉淀：if 初始化作用域、for 唯一循环三种形态、switch 顺序匹配/表达式case/fallthrough（不重新求值）/无条件=if-else链、break只退switch、label 跳出多层 → `chapters/03-控制结构.md`
- ✅ Part7 函数已学已沉淀：多返回值/命名返回值、一切传值(slice/map传"心脏")、闭包按引用共享 + Go1.22 循环变量新语义、defer(LIFO+参数立即求值+改命名返回值的时序) → `chapters/04-函数与方法.md`
- ⬜ **下一步：继续第一章/第二章收尾后的主干——make vs new、复合字面量、指针解引用等语法糖 Part 逐个推进；方法部分已在 ch02 三节覆盖**

完整章节见 `CLAUDE.md`（一语言基础 → 二语法糖 → 三数据结构 → 四控制流 → 五协程调度 → 六定时器 → 七并发控制 → 八接口 → 九内存 → 十工程）。

## 关键原则

1. 源码优先：`$(go env GOROOT)/src/runtime/chan.go` 等，不凭记忆
2. 卡住先跳代码：用流水线/比喻讲透心智模型，再回源码
3. 每个概念写小程序验证（跑通/看崩溃/看输出）
4. 每个主题沉淀一个 HTML 网页课程到 `courses/`
