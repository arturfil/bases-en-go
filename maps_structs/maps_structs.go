package mapsstructs

import (
	"fmt"

	"github.com/google/uuid"
)

type Product struct {
    Id uuid.UUID `json:"id"`
    Name string `json:"name"`
    Price float32 `json:"price"`
    Category string `json:"category"`
}

func (p *Product) returnInfo() string {
    info := fmt.Sprintf("Nombre: %v, Price: %v, Category: %v", p.Name, p.Price, p.Category)
    fmt.Println(info)
    return info
}

var productNames = []string {
    "Nintendo Switch",
    "Macbook Pro M1",
    "Mechanical Keyboard",
    "Raspberry Pi",
    "Acoustic Guitar",
}
var productPrices = []float32 {199.55, 1899.99, 201.22, 40.57, 89.99}
var productCategories = []string {"video games", "computers", "computers", "computers", "musical instrument"}

func CreateProductsArrMap() {
    // map[key]value
    pdMap := make(map[string]int)
    var productsArr []Product

    for i := 0; i < len(productNames); i++ {
        productsArr = append(productsArr, Product{
            Id: uuid.New(),
            Name: productNames[i],
            Price: productPrices[i],
            Category: productCategories[i],
        })
    }

    for _, product := range productsArr {
        fmt.Println(product)
        if _, key_exist := pdMap[product.Category]; key_exist {
            pdMap[product.Category] += 1
        } else {
            pdMap[product.Category] = 1
        }
    }
    fmt.Println("") 
    fmt.Println(pdMap)
    
    productsArr[0].returnInfo()
}

