package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Say() string       { return "hi, " + p.Name } // 值接收者
func (p *Person) SetName(n string) { p.Name = n }             // 指针接收者

type Counter struct{ n int }

func (c *Counter) Inc()    { c.n++ }        // 指针接收者：要原地改
func (c Counter) Val() int { return c.n }   // 值接收者：只读复制

func mk() Person { return Person{Name: "tmp"} }

func main() {
	fmt.Println("=== Rule1 自动解引用: p 是 *Person ===")
	p := &Person{Name: "alice"}
	p.Name = "bob" // 编译器自动 (*p).Name
	fmt.Printf("  p.Name = %q | p.Say() = %q\n", p.Name, p.Say())

	fmt.Println("\n=== Rule2 自动取地址: 变量可寻址 → 指针接收者 OK ===")
	v := Person{Name: "carol"}
	v.SetName("dave") // 编译器自动 (&v).SetName
	fmt.Printf("  v.Name = %q\n", v.Name)

	fmt.Println("\n=== Rule3 可寻址: slice 元素能调指针接收者方法 ===")
	s := []Person{{Name: "eve"}}
	s[0].SetName("frank") // &s[0] 有固定位置
	fmt.Printf("  s[0].Name = %q\n", s[0].Name)

	fmt.Println("\n=== 不可寻址三行: 编译错误(见 11_ptr_errors 的 go build 实测) ===")
	fmt.Println("  - &mk()        → 函数返回值, 现炒现卖没落点")
	fmt.Println("  - m[\"k\"].Inc() → map 元素会搬家, 拿不到地址")
	fmt.Println("  - &str[0]      → string 共享只读, 不许涂改")

	fmt.Println("\n=== 惯用法: map 存指针 → 桶随便搬, 目标对象不动 ===")
	m := map[string]*Counter{}
	m["k"] = &Counter{}
	m["k"].Inc() // m["k"] 本身就是 *Counter, 连 & 都不用
	fmt.Printf("  m[\"k\"].n = %d\n", m["k"].n)

	fmt.Println("\n=== 值接收者能直接用 map 元素上(只需读, 不需地址) ===")
	vm := map[string]Counter{"a": {n: 5}}
	fmt.Printf("  vm[\"a\"].Val() = %d\n", vm["a"].Val()) // 读值复制, OK
}
