package main

import "fmt"

func main() {
	var name string
	var totalSubject int
	fmt.Println("enter your name:")
	fmt.Scan(&name)

	fmt.Println("How many Subject?")
	fmt.Scan(&totalSubject)

	m:=make(map[int]int)

	var grade int
	var average int
	for i:=0 ;i < totalSubject; i++{
		fmt.Println("Enter Subject ",i+1)
		fmt.Scan(&grade)
		m[i+1] = grade
		average += grade
	}
	average = average / totalSubject
	
	fmt.Println("name", name)
	for i:=0;i<totalSubject;i++{
		fmt.Println("Subject ",i+1,":", m[i+1])
	}
	fmt.Println("Average:", average)

}
