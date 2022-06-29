package main

import (
	"fmt"
)

func main() {
	var numbers1 = make([]int, 0, 5) //กำหนดเเบบ limit จะเปลือง mem ถ้าไม่ได้ใช้
	numbers1 = append(numbers1, 1) 
	numbers1 = append(numbers1, 2)
	
	var numbers2 []int 
	numbers2 = append(numbers2, 1)
	numbers2 = append(numbers2, 2)
	debugslice(numbers1)
	debugslice(numbers2)
}

func debugslice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}