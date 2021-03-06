package car

// Car ...
type Car struct {
	Name  string
	Price int
	Color string
}

// GetName ...
func (c *Car) GetName() string {
	return c.Name
}

// UpdateColor ...
func (c *Car) UpdateColor(color string) {
	c.Color = color
}

// GetColor ...
func (c *Car) GetColor() string {
	return c.Color
}

// SetPrice ...
func (c *Car) SetPrice(price int) {
	c.Price = price
}

// SetPrice2 ...
func (c Car) SetPrice2(price int) Car {
	c.Price = price
	return c
}

// New ...
func New(name string, color string) *Car {

	return &Car{
		Name:  name,
		Color: color,
	}
}
