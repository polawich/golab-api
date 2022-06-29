package main
 
import "fmt"

func main() {
	var number = []int{1,2,3,4,5,6}
	debugslice(number)

	number = number[1: len(number)]
	debugslice(number)
	number = number[1: len(number)]
	debugslice(number)

	number = number[0: len(number)-1]
	debugslice(number)
	 
}

func debugslice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}