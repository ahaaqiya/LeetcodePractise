package main

import (
	"fmt"
	"strconv"
)

func main()  {
	var str string
	fmt.Scanln(&str)
	haveBash := false
	for i:=0;i<len(str);i++{
		if str[i]=='#'{
			haveBash = true
			break
		}
	}
	if !haveBash{
		if str[0]=='0'{
			if str[1]=='X' || str[1]=='x'{
				i:=len(str)-1
				res := 0
				count := 0
				for (str[i]>='0' && str[i]<='9') || (str[i]>='a' && str[i]<='f'){
					 if str[i]>='0' && str[i]<='9'{
					 	num,_ := strconv.Atoi(str[i:i+1])
					 	res += power(16,count)*num
					 }
					 i++
				}
				if i!=len(str){
					fmt.Println("ERROR")
				}
			}
		}
	}

	fmt.Println(power(3,4))
}

func power(a,b int) int{
	res := 1
	for b>=1{
		res = res*a
		b--
	}
	return res
}
