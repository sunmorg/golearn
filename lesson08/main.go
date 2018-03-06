package main

import "fmt"

func plus(a,b int )(res int){
	return a + b
}

func multi() (string,int){
	return "age:", 18
}

func nameReturnValue() (name string ,height int){
	name = "zhangsan"
	height = 180
	return
}

func Sum(nums ...int)(res int){
	fmt.Println("len of nums is:",len(nums))
	for _,v := range nums{
		res += v
	}
	return
}

func sayHello(name string){
	fmt.Println("hello ", name)
}
func logger(f func(string string),name string){
	fmt.Println("start calling method sayHello")
	f(name)
	fmt.Println("end calling method sayHello")
}

func sendValue(name string){
	name = "zhangsan"
}

func sendAddress(name *string){
	*name = "zhangsan"
}

func addInt(n int) func() int {
	i := 0
	return func() int {
		i += n
		return i
	}
}

func main(){
	//c := plus(2,3)
	//fmt.Println(c)
	//str,age := multi()
	//fmt.Println(str)
	//fmt.Println(age)
	//
	//name,height := nameReturnValue()
	//fmt.Println(name,"",height)
	//fmt.Println(Sum(1))
	//fmt.Println(Sum(1,2))
	//fmt.Println(Sum(1,2,3))
	//func(name string){
	//	fmt.Println(name)
	//}("beijing")
	//logger(sayHello,"zhangsan")

	//str := "lisi"
	//fmt.Println("befor value ", str)
	//sendValue(str)
	//fmt.Println("end value", str)
	//
	//fmt.Println("befor add ", str)
	//sendAddress(&str)
	//fmt.Println("end add", str)


	addOne := addInt(1)
	fmt.Println(addOne())
	fmt.Println(addOne())
	fmt.Println(addOne())

	addTwo := addInt(2)
	fmt.Println(addTwo())
	fmt.Println(addTwo())
	fmt.Println(addTwo())
}
