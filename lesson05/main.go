package main

import "fmt"

func main(){

	// 数组长度是数组类型的一部分

	var arr1 [5]int
	for i := 0; i < 5; i++ {
		arr1[i] = i
	}
	fmt.Println(arr1)

	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)
	printHelper("arr2",arr2)


	arr3 := [...]int{1,0,3,5,3}
	printHelper("arr3", arr3)


	arr4 := [...]int{4: -1}
	printHelper("arr4", arr4)

	var twoD [2][3]int
	for i:=0;i<2;i++{
		for j:=0;j<2;j++{
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)

}

func printHelper(name string , arr [5]int){
	fmt.Println(name,"",arr)

	fmt.Printf("len of %v: %v\n", name, len(arr))

	fmt.Printf("cap of %v: %v\n", name, cap(arr))
}
