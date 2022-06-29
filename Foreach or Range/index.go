package main

import (
	"fmt"
)

func main(){
	courses := []string {"apple", "ios", "andrie"}
	for index, item := range courses {
		fmt.Printf("%d. %s\n", index+1, item)
	}
	for _, item := range courses{
		fmt.Printf("%s\n",item)
	}
	stop()
}

func stop(){
	fmt.Println("-------break----------")
	index := 1
	for true {
		if index > 5 {
			break
		}
		fmt.Printf("index = %d\n", index)
		index++
	}
}