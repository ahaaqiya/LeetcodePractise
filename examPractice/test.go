package main

import (
	"fmt"
	"reflect"
)
func changeSlice(s []int)  {
	fmt.Printf("func： %p \n", &s)
	fmt.Println(reflect.ValueOf(s))
	s[1] = 111
	s = append(s,3)
	s[1] = 222
}

func main()  {
	slice := make([]int,4,6)
	slice = []int{0,1,2,3}
	fmt.Printf("slice: %v  slice addr %p \n",slice,&slice)
	changeSlice(slice)
	fmt.Printf("slice: %v  slice addr %p \n",slice,&slice)
}