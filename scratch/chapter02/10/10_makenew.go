package main

import "fmt"

// Part9: make vs new（+ 复合字面量）验证
// panic 演示保持注释，想看崩溃就取消注释

type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("=== 题1: make 非指针 vs new 指针 ===")
	s := make([]int, 3)
	s[0] = 7 // 能写：底层数组已就位
	fmt.Printf("  make([]int,3): %v (%T)\n", s, s)

	ns := new([]int)
	fmt.Printf("  new([]int): *ns == nil? %v (%T)\n", *ns == nil, ns)
	// (*ns)[0] = 1  // ✗ panic: index out of range [0] with length 0

	fmt.Println("\n=== 题2: new(map) 写会 panic；make(Person) 编译错 ===")
	nm := new(map[string]int)
	fmt.Printf("  new(map): *nm == nil? %v\n", *nm == nil)
	// (*nm)["a"] = 1  // ✗ panic: assignment to entry in nil map
	// make(Person)    // ✗ 编译错误：invalid argument: cannot make Person

	fmt.Println("\n=== 题3: &T{} 一步构造 vs new+两步赋值 ===")
	p1 := new(Person) // 两步：先零值指针，再逐个赋
	p1.Name = "alice"
	p2 := &Person{Name: "bob", Age: 20} // 一步：表达式直接可 inline
	fmt.Printf("  new+赋值:  %#v\n", p1)
	fmt.Printf("  &Person{}: %#v\n", p2)
	makePerson := func() *Person { return &Person{Name: "inline"} } // &T{} 能当返回值
	fmt.Printf("  inline 返回: %#v\n", makePerson())

	fmt.Println("\n=== 题4: [...]数组(值类型) vs []slice(窗口) ===")
	arr := [...]int{1, 2, 3}
	sl := []int{1, 2, 3}
	fmt.Printf("  [...]int{...} → %T（长度3写死在类型里）\n", arr)
	fmt.Printf("  []int{...}     → %T（窗口，可长可短）\n", sl)

	arr2 := arr // 数组赋值 = 整个复制（值类型）
	arr2[0] = 99
	fmt.Printf("  数组赋值后: arr[0]=%d, arr2[0]=%d（互不影响）\n", arr[0], arr2[0])

	sl2 := sl // slice 赋值 = 复制窗口（共享底层）
	sl2[0] = 99
	fmt.Printf("  slice赋值后: sl[0]=%d, sl2[0]=%d（共享底层！）\n", sl[0], sl2[0])
}
