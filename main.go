package main

import "fmt"

func run(g GeneralInterface, top int) {
	g.Connect()
	g.Select()
	g.Process(top)
	g.Disconnect()
}

func main() {
	products := Products{&DataAccessor{data: make([]string, 0, 10)}}
	categories := Categories{&DataAccessor{data: make([]string, 0, 10)}}

	run(&categories, 5)
	fmt.Println()
	run(&products, 2)
}

