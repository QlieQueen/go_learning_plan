package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) Greet() string { return "Hi, " + u.Name }
func (u *User) Birthday()    { u.Age++ }

func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

type Counter struct {
	Num int
}

func NewCounter(val int) *Counter {
	return &Counter{Num: val}
}

func (c Counter) Add(x int) {
	c.Num += x
}

func (c *Counter) Inc(x int) {
	c.Num += x
}

func main() {
	u1 := User{"阿绿", 25}
	u2 := User{Name: "阿绿", Age: 25}
	u3 := User{Name: "阿绿"}
	fmt.Printf("u1:%v, u2:%v, u3:%v\n", u1, u2, u3)

	u := NewUser("阿绿", 25)
	fmt.Println(u.Greet())
	u.Birthday()
	fmt.Println(u.Age)

	v := u
	v.Age = 100
	fmt.Println(u.Age, v.Age) // 100 100

	// 题目1
	c := Counter{Num: 100}
	c.Add(100)
	fmt.Println("c.Num:", c.Num) // 100
	c.Inc(100)
	fmt.Println("c.Num:", c.Num) // 200

	c1 := NewCounter(10)
	c1.Add(100)
	fmt.Println("c1.Num:", c1.Num) // 没有效果，仍然是10
	c1.Inc(100)
	fmt.Println("c1.Num:", c1.Num) // 生效，110

	// 题目2
	// User{Name: "x", Age: 18}.Birthday() // 报错 cannot call pointer method Birthday on User

	// 题目3
	a := User{Name: "xxx", Age: 18}
	b := a
	b.Birthday()
	fmt.Printf("a.Age:%d, b.Age:%d\n", a.Age, b.Age) // a.Age:18, b.Age=19
}
