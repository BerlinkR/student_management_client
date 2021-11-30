package main

import "fmt"

func input_student() (student){
	fmt.Println("Please input student message: ")
	var nstu student
	fmt.Scanln(&(nstu.Name),&(nstu.Id), &(nstu.Gender), &(nstu.MarkMath), &(nstu.MarkEnglish))
	return nstu
}