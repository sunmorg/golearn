package main

import "fmt"

type Person struct {
	Name string
	Age int
}

//type Student struct {
//	Age int
//	Name string
//}

type Student struct {
	Person
}

func main(){
	//stu := Student{
	//	Age: 18,
	//	Name: "zhangsan",
	//}
	//
	//stu1 := Student{
	//	18,"zhangsan",
	//}
	//
	//fmt.Println(stu)
	//fmt.Println(stu1)

	stu := Student{
		Person{
			Age: 18,
			Name: "zhangsan",
		},
	}

	fmt.Println(stu)

	fmt.Println(stu.Age)
	fmt.Println(stu.Name)

}