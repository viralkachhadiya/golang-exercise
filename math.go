package main

import (
	"fmt"
	"math"
)

const s string = "contant"

func mathFunction() {
	fmt.Println(s)
	const n = 500000
	const d = 3e20 / n
	fmt.Println(int(d))
	fmt.Println(math.Sin(n))
}
