package main

type Item struct {
	Name  string
	Price float64
}

type Cart struct {
	Items []Item
}

func InitializeCart() *Cart {
	return &Cart{Items: []Item{}}
}

func (c *Cart) AddItem(item Item) {
	c.Items = append(c.Items, item)
}

func (c *Cart) TotalPrice() (total float64) {
	for _, val := range c.Items {
		total += val.Price
	}

	if total > 1000 {
		total *= 0.9 // 10% discount
	}
	return
}
