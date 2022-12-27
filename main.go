package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	db2 "github.com/nuclearIgor/go-hexagonal/adapters/db"
	"github.com/nuclearIgor/go-hexagonal/application"
	"log"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db2.NewProductDb(db)

	productService := application.NewProductService(productDbAdapter)
	product, err := productService.Create("Product Example", 30)
	if err != nil {
		log.Fatal()
	}

	fmt.Println(product)

	productService.Enable(product)
	productService.Disable(product)
	productService.Enable(product)
	productService.Disable(product)
	productService.Enable(product)
	productService.Disable(product)
}
