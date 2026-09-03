package main

import "fmt"

// 第二章语法糖: := 短变量声明 + ... 变参
// 重点验证：部分重声明 / 遮蔽vs复用 / 类型推断 / 空变参

func main() {
	fmt.Println("=== 题1: 部分重声明规则 ===")
	x := 1
	x, y := 2, 3    // ① y 是新变量 → 合法
	// x := 4       // ② 没有新变量 → 编译错误：no new variables on left side of :=
	x, y = 5, 6     // ③ 这是赋值(=)，不是声明 → 合法
	x, z := 7, 8    // ④ z 是新变量 → 合法
	fmt.Printf("  x=%d y=%d z=%d\n", x, y, z) // 7 8 0

	fmt.Println("\n=== 题2: 同作用域【复用】 vs 嵌套作用域【遮蔽】 ===")
	// 同作用域：第二行 fB 新 + errA 同作用域【复用】→ errA 变成 flat2
	fA, errA := demo("flat1")
	fmt.Printf("  fA=%d errA=%v\n", fA, errA) // errA=flat1
	fB, errA := demo("flat2")                 // fB 新，errA 复用
	fmt.Printf("  fB=%d，复用后 errA=%v\n", fB, errA) // errA=flat2

	// 嵌套作用域：if 块里的 err2 是【新变量遮蔽外层】→ 外层不受影响
	f2, err2 := demo("outer")
	if f2 == 0 {
		f3, err2 := demo("inner") // f3 新；if 块内没有 err2 → err2 也是新的（遮蔽外层）
		fmt.Printf("  if 内: f3=%d err2=%v\n", f3, err2) // err2=inner
	}
	fmt.Printf("  if 外: err2=%v ← 外层没被嵌套块改掉\n", err2) // err2=outer

	fmt.Println("\n=== 题3: := 推断的默认类型 ===")
	i := 1
	c := 'A' // rune，但 rune 是 int32 的别名
	s := "hi"
	f := 3.14
	fmt.Printf("  i=%T c=%T s=%T f=%T\n", i, c, s, f) // int int32 string float64

	fmt.Println("\n=== 题4: 空变参 + slice 展开 ===")
	fmt.Println("  sum() =", sum()) // 0
	fmt.Println("  sum(1,2,3) =", sum(1, 2, 3))
	s2 := []int{4, 5, 6}
	fmt.Println("  sum(s2...) =", sum(s2...))
}

func demo(label string) (int, error) { return 0, fmt.Errorf("%s", label) }

func sum(nums ...int) int {
	fmt.Printf("    (进入 sum: nums==nil? %v, len=%d)\n", nums == nil, len(nums))
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}
