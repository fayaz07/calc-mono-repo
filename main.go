package main

import (
	"fmt"

	"github.com/fayaz07/calc-mono-repo/add"
	"github.com/fayaz07/calc-mono-repo/divide"
	"github.com/fayaz07/calc-mono-repo/multiply"
	"github.com/fayaz07/calc-mono-repo/subtract"
)

func main() {
	fmt.Println(add.Add(1, 2, 3, 4, 5))
	fmt.Println(subtract.Subtract(100, 293, 384, 1, 2))
	fmt.Println(multiply.Multiply(1, 2))
	fmt.Println(divide.Divide(4, 2))
}
