package main

import "fmt"

type GeneralInterface interface {
	Connect()
	Select()
	Process(int)
	Disconnect()
}

type DataAccessor struct {
	data []string
}

func (*DataAccessor) Connect() {
	fmt.Println("Connected")
}
func (da *DataAccessor) Process(top int) {
	if top >= len(da.data) {
		fmt.Println("Out of range")
		return
	}
	for i := 0; i <= top; i++ {
		fmt.Println(da.data[i])
	}
}
func (*DataAccessor) Disconnect() {
	fmt.Println("Disconnect")
}
