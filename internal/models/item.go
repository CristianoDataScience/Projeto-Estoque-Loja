package models

import "fmt"

type Item struct {
	ID       int
	Name     string
	Quantily int
	Price    float64
}

func (i Item) Info() string {
	return fmt.Sprintf("ID: %d | Name: %s | Quantily: %d | Price: %.2f",
		i.ID, i.Name, i.Quantily, i.Price)
}
