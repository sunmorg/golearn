package main

import (
	"math"
	"fmt"
)

type geometry interface{
	area() float32
	perim() float32
}

type rect struct {
	len,width float32
} 

func (r *rect) area() float32{
	return r.len * r.width
}

func (r *rect) perim() float32{
	return 2*(r.len + r.width)
}

type  circle struct {
	radius float32
} 

func (c *circle) area() float32{
	return math.Pi * c.radius * c.radius
}

func (c *circle) perim() float32 {
	return 2*math.Pi * c.radius
}

func show(name string,params interface{}){
	switch  params.(type) {
	case geometry:
		fmt.Println("area of %v is %v \n", name, params.(geometry).area())
		fmt.Println("perim of %v is %v \n", name, params.(geometry).perim())
	default:
		fmt.Println("wrong type!!!")

	}
}

func main(){
	r := &rect{
		len: 1,
		width: 2,
	}

	show("rect", r)

	cir := &circle{
		radius: 1,
	}

	show("circle", cir)

}