package main

import "fmt"

//type City string
//
//type (
//	B1 int
//	B2 int
//)

//type Name string

type (
	Height int
	Width  int
)



//func printName(name string){
//	fmt.Println("my name is ",name)
//}

func main (){
	height := Height(180)
	width := Width(10)
	//myCity := City("shanghai")

	//fmt.Println("mycity is ",myCity)
	//
	//fmt.Println("B1 is ",B1(1))
	//fmt.Println("B2 is ",B2(10))

	//fmt.Println(myCity + "is mycity")
	//fmt.Println(len(myCity))
	//
	//if B1(5) > B1(4) {
	//	fmt.Println("B1(5) is bigger than B1(4)")
	//}

	//myName := Name("zhangsan")
	//printName(string(myName))

	fmt.Println("height/width :",int(height)/int(width))

}