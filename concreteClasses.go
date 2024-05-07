package main

type Categories struct {
	*DataAccessor
}

func (c *Categories) Select() {
	c.data = append(c.data, "Red")
	c.data = append(c.data, "Green")
	c.data = append(c.data, "Blue")
	c.data = append(c.data, "Yellow")
	c.data = append(c.data, "Purple")
	c.data = append(c.data, "White")
	c.data = append(c.data, "Black")
}

type Products struct {
	*DataAccessor
}

func (p *Products) Select() {
	p.data = append(p.data, "Car")
	p.data = append(p.data, "Bike")
	p.data = append(p.data, "Boat")
	p.data = append(p.data, "Truck")
	p.data = append(p.data, "Moped")
	p.data = append(p.data, "Rollerskate")
	p.data = append(p.data, "Stroller")
}
