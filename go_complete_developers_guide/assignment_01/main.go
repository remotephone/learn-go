package main

import (
	"fmt"
	"strconv"
)

func main() {
// First had to define a range, got range [10]intP{} wrong, 11 was necessary to get to 10
	for i := range [11]int{} {
	// simple comparison, similar to python modulus function
	if i % 2 == 0 {
		// This required import and FormatInt, the 10 is weird, this is base 10.
		fmt.Println(strconv.FormatInt(int64(i), 10) +" is even")
	} else {
		fmt.Println(strconv.FormatInt(int64(i), 10) +" is odd")
		}
	}
}


