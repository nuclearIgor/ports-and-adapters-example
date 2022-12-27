package dto

import "github.com/nuclearIgor/go-hexagonal/application"

type Product struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
	Status string `json:"status"`
}

func (p *Product) Bind(product *application.Product) (*application.Product, error) {
	if p.Id != "" {
		product.Id = p.Id
	}
	product.Name = p.Name
	product.Price = p.Price
	product.Status = p.Status

	_, err := product.IsValid()
	if err != nil {
		return &application.Product{}, err
	}
	return product, nil
}