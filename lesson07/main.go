package main

import (
	"fmt"
)

func main(){
	//var cMap map[string]int
	//fmt.Println(cMap == nil)
	//
	//cMap["beijing"] = 1

	//cMap := make(map[string]int, 10)
	//fmt.Println(cMap == nil)
	//
	//cMap["beijing"] = 1
	//fmt.Println(cMap)

	//cMap := map[string]int{"beijing":1}
	//fmt.Println(cMap)
	//cMap := make(map[string]int)
	//
	//cMap["beijing"] = 1
	//fmt.Println(cMap)
	//
	//code := cMap["beijing"]
	//fmt.Println(code)
	//
	//code, ok := cMap["shanghai"]
	//if ok {
	//	fmt.Println(code)
	//} else {
	//	fmt.Println("shanghai not exist")
	//}
	//
	//delete(cMap,"beijing")
	//fmt.Println(cMap)
	//cMap := map[string]int{
	//	"shanghai": 1,
	//	"beijing": 2,
	//	"shenzhen": 3,
	//}
	//
	//for k,v := range cMap {
	//	fmt.Printf("%s is %d", k, v)
	//	fmt.Println()
	//}

	//cMap := make(map[string]int)
	//
	//var wg sync.WaitGroup
	//wg.Add(2)
	//
	//go func() {
	//	cMap["beijing"] = 1
	//	wg.Done()
	//}()
	//
	//go func() {
	//	cMap["shanghai"] = 1
	//	wg.Done()
	//}()
	//
	//wg.Wait()
	//
	//fmt.Println(cMap)
	cMap := make(map[string]map[string]int)

	cMap["beijing"] = map[string]int{
		"shanghai" : 1,
	}
	fmt.Println(cMap)

}
