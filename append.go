package main

import "fmt"

func NewAppend() *Append {
	fmt.Println("--creating new array...")

	return &Append{
		Data:      make([]int, 2),
		NextIndex: 0,
	}
}
