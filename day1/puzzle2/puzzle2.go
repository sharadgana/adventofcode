package main

import (
	"fmt"
	"strings"
)

//INPUT is the floor that Santa should go to
const INPUT = "()())"

//START is the starting floor of Santa
const START = 0

func main() {
	floor := START
	for index, f := range strings.Split(INPUT, "") {
		if f == "(" {
			floor++
		} else if f == ")" {
			floor--
		}
		if floor == -1 {
			fmt.Printf("The position at which Santa entered the basement is %d\n", index+1)
		}
	}
	fmt.Printf("The floor that Santa should go to is %d\n", floor)
}
