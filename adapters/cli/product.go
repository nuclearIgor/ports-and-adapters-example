package cli

import (
	"fmt"
	"github.com/nuclearIgor/go-hexagonal/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, price int) (string, error) {
	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product Id %s with the name %s has been created with price %d and status %s",
			product.GetId(), product.GetName(), product.GetPrice(), product.GetStatus())
	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		res, err := service.Enable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s has been enabled", res.GetName())
	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		res, err := service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s has been disabled", res.GetName())

	default:
		res, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product Id: %s\nName: %s\nPrice: %d\nStatus: %s",
			res.GetId(), res.GetName(), res.GetPrice(), res.GetStatus())
	}
	return result, nil
}
