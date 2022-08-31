package main

import "fmt"

func NewAppend() *Append {
	fmt.Println("--creating new array...")

	return &Append{
		Data:      make([]int, 2),
		NextIndex: 0,
	}
}

func (a *Append) Insert(newData int) {
	fmt.Printf("--inserting new data : %v\n", newData)

	if a.NextIndex > cap(a.Data)-1 {
		newCapacity := cap(a.Data) * 2
		newArray := make([]int, newCapacity)
		copy(newArray, a.Data)
		a.Data = newArray
	}

	a.Data[a.NextIndex] = newData
	a.NextIndex++
}
