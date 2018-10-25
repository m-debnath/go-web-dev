package main

import (
	"fmt"
)

type person struct {
	fname   string
	lname   string
	age     int
	sayings []string
}

func main() {
	// slice
	// composite literal; slice literal
	x := []int{7, 9, 42}
	for i, _ := range x {
		fmt.Println(i, "-", x[i])
	}

	y := make([]int, 0, 10)
	y = append(y, 777)

	for i, v := range y {
		fmt.Println(i, "-", v)
	}

	// map
	m := []map[string]int{
		{
			"Mukul":    32,
			"Diptajit": 33,
		},
		{
			"James":      27,
			"Moneypenny": 25,
		},
		{
			"David": 81,
			"Stan":  93,
		},
	}
	for i, v := range m {
		fmt.Println(i)
		for x, n := range v {
			fmt.Println("\t", x, ":", n)
		}
	}

	// struct
	p1 := person{
		fname: "Mukul",
		lname: "Debnath",
		age:   32,
		sayings: []string{
			"I've lost my brain!",
			"I think, therefore I am.",
		},
	}
	p2 := person{
		fname: "Diptajit",
		lname: "Pal",
		age:   33,
		sayings: []string{
			"What a lovely day!",
			"Why so serious?",
		},
	}
	people := []person{p1, p2}
	for _, v := range people {
		fmt.Println(v.fname, v.lname, "of age", v.age, "says")
		for _, s := range v.sayings {
			fmt.Println("\t", s)
		}
	}
}
