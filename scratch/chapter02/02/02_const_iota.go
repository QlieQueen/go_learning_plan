package main

import "fmt"

const (
	Mon = iota
	Tue
	Wed
	Thu
	Fri
	Sat
	Sun
)

const (
	A = iota
	_
	B
	C
)

const (
	FlagA = 1 << iota
	FlagB
	FlagC
	FlagD
)

func main() {
	fmt.Println(Mon, Tue, Wed, Thu, Fri, Sat, Sun)
	fmt.Println(A, B, C)
	fmt.Println(FlagA, FlagB, FlagC, FlagD)
}
