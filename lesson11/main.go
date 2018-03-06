package main

import "fmt"

type  Student struct {
	name string
	age int
}

func (s Student) sayHello() {
	fmt.Println("hello")
}

// 自定义类型 或者结构体

func (s *Student) addAge(){
	s.age +=1
}

func (s Student) printAge(){
	fmt.Println(s.name,"age is",s.age)
}

func main(){
	s := Student{"xiaomin",8}
	s.printAge()
	s.addAge()
	s.printAge()
}

