package main

import (
	"fmt"
)

// 根据第一个字节判断宽度
func runeWidth(b byte) int {
	switch {
	case b < 0x80:
		return 1 // 0开头
	case b < 0xE0:
		return 2 // 110
	case b < 0xF0:
		return 3 // 1110
	default:
		return 4 // 11110
	}
}

func main() {
	s := "你好,Go"

	// 坑1：len是字节数
	fmt.Println("字节数:", len(s))

	// 坑2a：按字节遍历 -- 中文会碎成一堆字节
	// [0]=228(ä) [1]=189(½) [2]=160( ) [3]=229(å) [4]=165(¥) [5]=189(½) [6]=44(,) [7]=71(G) [8]=111(o)
	for i := 0; i < len(s); i++ {
		fmt.Printf("[%d]=%v(%c) ", i, s[i], s[i])
	}
	fmt.Println()

	// 坑2b：range 按 rune 走 -- 一个字符一次
	// [0]='你' [3]='好' [6]=',' [7]='G' [8]='o'
	for i, r := range s {
		fmt.Printf("[%d]=%q ", i, r)
	}
	fmt.Println()

	// 坑3：切片切的是字节位
	// 你好
	fmt.Printf("%q\n", s[0:6])

	// 题目2
	str := "A中B"
	for i := 0; i < len(str); i++ {
		fmt.Printf("[%d]=%c:%x ", i, str[i], str[i])
	}
	fmt.Println()

	for i, r := range str {
		fmt.Printf("[%d]=%q ", i, r)
	}
	fmt.Println()

	for i := 0; i < len(str); {
		w := runeWidth(str[i])
		fmt.Printf("[%d]=%q ", i, string(str[i:i+w]))
		i += w
	}
	fmt.Println()

	// 题目3
	a := "Go"
	b := a + "lang"
	fmt.Println(a, b)

	// 题目4
	str = "你好"
	tmp := str[0:2]
	fmt.Println(tmp) // 乱码

}
