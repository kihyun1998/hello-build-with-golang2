package main

import (
	"fmt"

	"github.com/TeamTestCodeowner/calc"
	rand "github.com/TeamTestCodeowner/rand"
)

func main() {
	num1 := rand.RandomInt(1, 1000)
	num2 := rand.RandomInt(1, 1000)

	num3 := calc.Sub(num1, num2)
	fmt.Printf("%d - %d = %d\n", num1, num2, num3)
}
