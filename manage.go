package main

import "fmt"

func input_student() (student){
	fmt.Println("Please input student message:")
	fmt.Println("Name  Id  Gender  MarkMath  MarkEnglish")
	var nstu student
	fmt.Scanln(&(nstu.Name),&(nstu.Id), &(nstu.Gender), &(nstu.MarkMath), &(nstu.MarkEnglish))
	return nstu
}
func inputButton() (int){
	fmt.Println("Please input function number to use:")
	var choose int
	fmt.Scanln(&choose)
	return choose
}

func showFunctionList() {
	fmt.Println("Show list : 1")
	fmt.Println("Order list : 2")
	fmt.Println("Insert a new student : 3")
	fmt.Println("show function list again : 4")
}