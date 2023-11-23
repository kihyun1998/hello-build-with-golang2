package main

import (
	"fmt"
	"hello-build-with-golang2/calc"
	"hello-build-with-golang2/util"
)

func main() {
	num1 := util.RandomInt(1, 1000)
	num2 := util.RandomInt(1, 1000)

	num3 := calc.Sub(num1, num2)
	fmt.Println("result is ", num3)
}
