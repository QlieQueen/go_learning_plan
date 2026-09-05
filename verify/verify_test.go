// Package verify 把各章节的"核心结论"固化成自动断言。
// 用法：go test ./verify/   （或在仓库根目录 go test ./verify）
// 目的：Go 版本/语义若有变化，一键回归确认我们的理解仍成立。
// 每个 TestXxx 对应一个学习结论（注释里标注章节）。
package verify

import "testing"

// ===== Part1 变量/常量/iota (ch02 一节) =====

// iota 可被 _ 跳过编号；1<<iota 是位标志
func TestIotaSkipAndShift(t *testing.T) {
	const (
		A = iota
		_ // 跳过 1
		B
		C
	)
	if A != 0 || B != 2 || C != 3 {
		t.Fatalf("iota skip wrong: A=%d B=%d C=%d", A, B, C)
	}

	const (
		FlagA = 1 << iota // 1
		FlagB             // 2
		FlagC             // 4
	)
	if FlagA != 1 || FlagB != 2 || FlagC != 4 {
		t.Fatalf("1<<iota flags wrong: %d %d %d", FlagA, FlagB, FlagC)
	}
}

// ===== Part2 string (ch02 二节) =====

// string 按字节算 len；range 按 rune 迭代
func TestStringIsBytes(t *testing.T) {
	if len("hello") != 5 {
		t.Fatal("ascii len should be 5 bytes")
	}
	if len("你好世界") != 12 { // 4 个 rune × 3 字节 = 12
		t.Fatalf("chinese string should be 12 bytes, got %d", len("你好世界"))
	}
	n := 0
	for range "你好世界" {
		n++
	}
	if n != 4 {
		t.Fatalf("range over string should yield 4 runes, got %d", n)
	}
}

// ===== Part3 struct 与方法 (ch02 三节) =====

type ageBox struct{ n int }

func (a ageBox) growValue() { a.n++ } // 值接收者：改副本
func (a *ageBox) growPtr()  { a.n++ } // 指针接收者：改原对象

func TestStructValueVsPointerReceiver(t *testing.T) {
	a := ageBox{n: 10}
	a.growValue()
	if a.n != 10 {
		t.Fatalf("value receiver must not change field, got %d", a.n)
	}
	a.growPtr()
	if a.n != 11 {
		t.Fatalf("pointer receiver must change field, got %d", a.n)
	}
}

// ===== Part4 slice (ch02 五节) =====

// 窗口共享：b := a[1:4]，改 b 元素会改 a
func TestSliceWindowSharing(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5}
	b := a[1:4]
	b[0] = 99
	if a[1] != 99 {
		t.Fatal("b[0] should alias a[1] (window sharing)")
	}
}

// append 扩容 = 断开共享（搬到新数组）
func TestSliceReallocBreaksSharing(t *testing.T) {
	c := []int{1, 2, 3} // cap=3
	d := c
	d = append(d, 4) // 扩容搬家
	d[0] = 99
	if c[0] != 1 {
		t.Fatal("after realloc, d must be detached from c")
	}
}

// ===== Part5 map (ch02 六节) =====

// nil map 读安全（返回零值 + ok=false）
func TestNilMapReadSafe(t *testing.T) {
	var m map[string]int // nil map
	v, ok := m["x"]
	if ok || v != 0 {
		t.Fatalf("reading nil map should give zero value, got v=%d ok=%v", v, ok)
	}
}

// nil map 写会 panic
func TestNilMapWritePanics(t *testing.T) {
	var m map[string]int
	defer func() {
		if recover() == nil {
			t.Fatal("writing to nil map must panic")
		}
	}()
	m["a"] = 1 //nolint:staticcheck // 故意触发 panic
}

// ===== Part6 控制流 (ch03) =====

// switch 顺序匹配第一个为真的 case（不是跳表）
func TestSwitchFirstMatchOrder(t *testing.T) {
	x := 10
	got := ""
	switch x {
	case x / 2: // 表达式 case：=5 不匹配
		got = "half"
	case 10:
		got = "ten"
	}
	if got != "ten" {
		t.Fatalf("switch should match second case, got %q", got)
	}
}

// switch 里的 break 只退 switch，不退 for
func TestBreakInSwitchOnlyExitsSwitch(t *testing.T) {
	count := 0
	for i := 0; i < 3; i++ {
		switch i {
		case 1:
			break // 只退 switch
		}
		count++
	}
	if count != 3 {
		t.Fatalf("for must run all 3 iterations, count=%d", count)
	}
}

// fallthrough 不重新求值下一个 case 的条件，直接穿进 body
func TestFallthroughDoesNotReeval(t *testing.T) {
	evaluated := false
	probe := func() int { evaluated = true; return 2 }
	got := ""
	switch 1 {
	case 1:
		got += "one"
		fallthrough
	case probe():
		got += "two"
	}
	if got != "onetwo" {
		t.Fatalf("fallthrough should run next case body, got %q", got)
	}
	if evaluated {
		t.Fatal("fallthrough must NOT evaluate next case expression")
	}
}

// ===== Part7 函数 (ch04) =====

// Go 1.22：循环变量每次迭代独立，闭包各捕各的
func TestLoopClosurePerIteration(t *testing.T) {
	// go.mod 声明 go 1.22 → 走新语义
	var funcs []func() int
	for i := 0; i < 3; i++ {
		funcs = append(funcs, func() int { return i })
	}
	for i, f := range funcs {
		if got := f(); got != i {
			t.Fatalf("closure[%d]=%d, want %d (Go1.22 per-iteration)", i, got, i)
		}
	}
}

// defer 夹在"return 求值"和"真正返回"之间，能改命名返回值
func TestDeferCanModifyNamedReturn(t *testing.T) {
	f := func() (n int) {
		defer func() { n++ }()
		return 5
	}
	if got := f(); got != 6 {
		t.Fatalf("defer should bump named return to 6, got %d", got)
	}
}

// 函数内 append 扩容后：本地 slice 写入与外部断连
func TestSliceReallocDivorceInFunc(t *testing.T) {
	tweak := func(s []int) []int {
		s[0] = 99          // 扩容前：共享，外部可见
		s = append(s, 100) // cap 满 → 扩容搬家
		s[1] = 98          // 扩容后：写新数组，外部不可见
		return s
	}
	s := []int{1, 2} // cap=2
	inner := tweak(s)
	if s[0] != 99 || s[1] != 2 {
		t.Fatalf("outside slice should be [99 2], got %v", s)
	}
	if len(inner) != 3 || inner[1] != 98 || inner[2] != 100 {
		t.Fatalf("inner slice after realloc should be [99 98 100], got %v", inner)
	}
}

// ===== Part8 := 与变参 (ch06 一节) =====

// 空变参调用 → nums 是 nil（遍历天然安全）
func TestVariadicEmptyIsNil(t *testing.T) {
	first := func(nums ...int) (bool, int) { return nums == nil, len(nums) }
	isNil, n := first()
	if !isNil || n != 0 {
		t.Fatalf("empty variadic should be nil slice, got nil=%v len=%d", isNil, n)
	}
}

// ===== Part9 make vs new (ch06 三节) =====

// make 初始化内脏可写；new 只给指向 nil 的指针
func TestMakeVsNew(t *testing.T) {
	s := make([]int, 3)
	s[0] = 1 // 能写
	ns := new([]int)
	if *ns != nil {
		t.Fatal("new([]int) should point to nil slice")
	}

	m := make(map[string]int)
	m["a"] = 1 // 能写
	nm := new(map[string]int)
	if *nm != nil {
		t.Fatal("new(map) should point to nil map")
	}
}

// ===== Part10 指针可寻址性 (ch06 四节) =====

type counter struct{ n int }

func (c *counter) inc() { c.n++ }

// map 存指针 → 桶搬家不影响，能原地调用指针接收者方法
func TestMapStorePointerForMutation(t *testing.T) {
	m := map[string]*counter{}
	m["k"] = &counter{}
	m["k"].inc() // 无需 & —— m["k"] 本身就是 *counter
	if m["k"].n != 1 {
		t.Fatal("pointer value in map should allow in-place mutation")
	}
}

// 值接收者/读操作可以直接用于 map 元素（只需读副本，不需地址）
func TestValueReceiverOnMapElement(t *testing.T) {
	read := func(c counter) int { return c.n } // 值语义：读副本
	m := map[string]counter{"a": {n: 5}}
	if read(m["a"]) != 5 {
		t.Fatal("value read on map element should work")
	}
}

// ===== Part10.5 数组 (ch02 四节) =====

// 数组赋值 = 深拷贝；数组可比较
func TestArrayValueCopy(t *testing.T) {
	a := [3]int{1, 2, 3}
	b := a
	b[0] = 99
	if a[0] != 1 {
		t.Fatal("array assignment must deep-copy (value type)")
	}
	if [2]int{1, 2} != [2]int{1, 2} {
		t.Fatal("comparable arrays should be == on equal elements")
	}
}
