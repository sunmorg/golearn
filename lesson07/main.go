package main

import "fmt"

func main(){
	//var cMap map[string]int
	//fmt.Println(cMap == nil)
	//
	//cMap["beijing"] = 1

	cMap := make(map[string]int, 10)
	fmt.Println(cMap == nil)

	cMap["beijing"] = 1
	fmt.Println(cMap)
}
