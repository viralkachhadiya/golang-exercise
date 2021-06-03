package main

import "fmt"

func loopFunction() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	i = 1
	for {
		if i == 5 {
			break
		}
		fmt.Println("loop")
		i++
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		} else {
			fmt.Println("....")
		}
		fmt.Println(n)
	}
}
