package model

import "github.com/satori/go.uuid"

type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewProduct() *Product {
	return &Product{
		ID: uuid.NewV4().String(),
	}
}

type Products struct {
	Product []*Product `json:"products"`
}

func (p *Products) Add(product *Product) {
	p.Product = append(p.Product, product)
}

func NewProducts() *Products {
	return &Products{}
}
