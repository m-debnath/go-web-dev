package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string
}

func main() {
	// slice
	// composite literal; slice literal
	x := []int{7, 9, 42}
	fmt.Println(x)

	y := make([]int, 0, 100)
	y = append(y, 8)
	y = append(y, 12)
	y = append(y, 43)
	fmt.Println(y)

	// map
	m := map[string]int{
		"Mukul":    32,
		"Diptajit": 33,
	}
	fmt.Println(m)
	// struct
	p := person{
		"Mukul",
		"Debnath",
	}
	fmt.Println(p)
}
