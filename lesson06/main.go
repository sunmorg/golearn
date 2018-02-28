package main

import "fmt"

func main(){
	//arr := [3]int{1, 2 ,3}
	//fmt.Println(arr)
	//
	//sliceA := []int{1, 2, 3}
	//printHelper("sliceA",sliceA)
	//
	//sliceB := make([]int,5)
	//printHelper("sliceB",sliceB)
	//
	//sliceC := make([]int, 5, 10)
	//printHelper("sliceC",sliceC)
	//
	//sliceD := arr[0:3]
	//printHelper("sliceD",sliceD)
	//
	//sliceE := sliceD[:]
	//printHelper("sliceE",sliceE)
	//
	//sliceF := []int{1, 2, 3, 4}
	//printHelper("sliceF",sliceF)
	//sliceF = append(sliceF, 3)
	//printHelper("sliceF",sliceF)

	slice := []int{1, 2, 3}
	//newSlice := slice[:]

	newSlice := make([]int,len(slice))

	copy(newSlice, slice)

	printHelper("slice",slice)
	printHelper("newSlice",newSlice)

	slice[0] = 5
	printHelper("slice",slice)
	printHelper("newSlice",newSlice)
}

func printHelper(name string, slice []int){
	fmt.Println(name,": ", slice)

	fmt.Printf("len of %v: %v\n", name, len(slice))

	fmt.Printf("cap of %v: %v\n", name, cap(slice))
}
