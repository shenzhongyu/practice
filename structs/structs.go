package main

import (
	"fmt"
	"math"
)

var fp = fmt.Println

// 接口实现

// 定义任何几何都需要实现的接口
type geometry interface {
	area() float64
	perim() float64
}

// struct定义矩形需要提供的用户数据
type rect struct {
	width, height float64
}

// struct定义圆需要提供的用户数据
type circle struct {
	radius float64
}

// 面积具体实现
func (r rect) area() float64 {
	return r.width * r.height
}

// 周长具体实现
func (r rect) perim() float64 {
	return (r.height + r.width) * 2
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return math.Pi * c.radius * 2
}

// 提供一个通用函数可以调用接口中的方法
func measure(g geometry) {
	fp(g)
	fp(g.area())
	fp(g.perim())
}

func main() {
	r := rect{width: 3, height: 5}
	c := circle{radius: 3}
	measure(r)
	measure(c)
}