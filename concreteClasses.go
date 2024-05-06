package main

import "fmt"

type GeneralInterface interface {
	Connect()
	Select()
	Process(top int)
	Disconnect()
}

//categories
type Categories struct {
	categoryList []string
}

func (c *Categories) Connect() {
	fmt.Println("Categories connected")
}
func (c *Categories) Select() {
	c.categoryList = append(c.categoryList, "Laptop")
	c.categoryList = append(c.categoryList, "Monitor")
	c.categoryList = append(c.categoryList, "CPU")
	c.categoryList = append(c.categoryList, "GPU")
	c.categoryList = append(c.categoryList, "RAM")
}
func (c *Categories) Process(top int) {
	fmt.Println("Categories: ")
	if top >= len(c.categoryList) {
		fmt.Println("Top out range")
		return
	}
	for i := 0; i <= top; i++ {
		fmt.Println(c.categoryList[i])
	}
}
func (c *Categories) Disconnect() {
	fmt.Println("Categories disconnected")
}

// products
type Products struct {
	productList []string
}

func (p *Products) Connect() {
	fmt.Println("Products connected")
}
func (p *Products) Select() {
	p.productList = append(p.productList, "Macbook M1 2021")
	p.productList = append(p.productList, "Asus Monitor 2022")
	p.productList = append(p.productList, "CPU ABZ")
	p.productList = append(p.productList, "GPU ABZ")
	p.productList = append(p.productList, "RAM ABZ")
}
func (p *Products) Process(top int) {
	fmt.Println("Products: ")
	if top >= len(p.productList) {
		fmt.Println("Top out range")
		return
	}
	for i := 0; i <= top; i++ {
		fmt.Println(p.productList[i])
	}
}
func (p *Products) Disconnect() {
	fmt.Println("Products disconnected")
}
