package main

import (
	"GOPROJECT/calc"
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {
	fmt.Println(calc.Add(1, 2))
	fmt.Println(calc.Sub(1, 2))
	a := decimal.NewFromFloat(1.2)
	fmt.Println(a)

}
