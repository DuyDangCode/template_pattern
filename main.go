package main

import "fmt"

func main() {
	categories := Categories{}
	da := DataAccessor{&categories}
	da.run(7)

	fmt.Println()
	products := Products{}
	da = DataAccessor{&products}
	da.run(3)
}
