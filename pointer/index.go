package main

import "fmt"

func main() {
	msg := "some message"
	var msgPointer *string = &msg
	fmt.Println(msg)
	fmt.Println(*msgPointer)
	changemessage(&msg)
	fmt.Println(msg)
}

func changemessage(apointer *string){
	*apointer = "newfuncmessage"
}ิ