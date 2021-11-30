package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func ShowList(){
	fetch("http://localhost:8000/show")
}
func ShowOrder(){
	fetch("http://localhost:8000/order")
}

func RemoteInsert(stu student) (error){
	err := writepost("http://localhost:8000/insert",stu)
	if err != nil{
		return err
	}
	return nil
}

func fetch(url string){
	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	fmt.Printf("%s", b)
}

func generateUrlMap(stu student) (map[string][]string){
	stuMap := make(map[string][]string)
	stuMap["Id"] = []string{1:stu.Id}
	stuMap["Name"] = []string{1:stu.Name}
	stuMap["Gender"] = []string{1:stu.Gender}
	stuMap["MarkMath"] = []string{1:strconv.Itoa(stu.MarkMath)}
	stuMap["MarkEnglish"] = []string{1:strconv.Itoa(stu.MarkEnglish)}
	return stuMap
}

func writepost(url string, stu student) (error){

	stuMap := generateUrlMap(stu)
	resp, err := http.PostForm(url,stuMap)
	if err != nil{
		fmt.Println("err in witepost")
		return err
	}
	defer resp.Body.Close()
	return nil
}