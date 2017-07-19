package main

import (
	"fmt"
)

func main() {
	s := []string{"1", "2"}
	fmt.Println(s)

	for i := range s {
		fmt.Println(i)
	}
}
