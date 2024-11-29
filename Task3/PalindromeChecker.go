package main

import "fmt"

func main(){
	s := "goodMorning"
	var rev string
	for i:=len(s)-1;i>=0;i--{
		rev += string(s[i])
	}
	if rev==s{
		fmt.Println("Palindrome")
	}else{
		fmt.Println("Not Palindrome")
	}
}