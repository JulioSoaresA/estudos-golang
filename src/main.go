package main

import (
	"fmt"
	"list/arraylist"
)

func main() {
	array := arraylist.ArrayList{}
	array.Init(5)
	fmt.Println(array)
	array.Add(1)
	array.Add(2)
	array.Add(3)
	array.Add(4)
	array.Add(5)
	array.Add(6)
	array.Display()
	array.RemoveOnIndex(2)
	array.Display()
}
