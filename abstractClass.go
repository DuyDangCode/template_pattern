package main

// abstract class
type DataAccessor struct {
	GeneralInterface
}

func (da *DataAccessor) run(top int) {
	da.Connect()
	da.Select()
	da.Process(top)
	da.Disconnect()
}
