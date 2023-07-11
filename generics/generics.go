package generics

import (
	"fmt"
)

type Ropa struct {
    tipo string
    precio int
}

type Comida struct {
    nombre string
    precio float64
    calorias int
}

type Item interface {
    Comida | Ropa
}

type Monto interface {
    int | float64
}

func ObtenerTotal[T Item, M Monto](grupo []T, getPrecio func(obj T) M) M {
    sum := M(0)    
    for _, obj := range grupo {
        sum += getPrecio(obj) 
    }
    return sum
}

func MainGenerics() {

    cuenta := []Comida {
        {"Pizza", 8.99, 1400},
        {"Hamburgesa", 7.20, 500},
        {"Tacos", 11.20, 800},
    }

    carrito := []Ropa {
        {"playera", 15},
        {"jeans", 35},
        {"hoodie", 24},
    }

    cuentaTot := ObtenerTotal(cuenta, func(comida Comida) float64 {
        return comida.precio
    })
    carritoTot := ObtenerTotal(carrito, func (ropa Ropa) int  {
       return ropa.precio 
    })

    fmt.Printf("El total del carrito es:\t%v\n", carritoTot)
    fmt.Printf("El total de la cuenta es:\t%v\n", cuentaTot)
}
