package main

import (
	"fmt"
)

func main() {
	if check := dosomething(); check == "ok" {
		fmt.Println("Done")
	} else{
		fmt.Println("No")
	}
	switch2()
}

func dosomething() string{
	return "ok"
}

func switch2(){
	index := 3
	switch index {
	case 0:
		fmt.Println("zero")
		break
	case 1:
		fmt.Println("one")
		break
	case 2:
		fmt.Println("two")
		break
	default:
		fmt.Println("default")
		break
	}
}