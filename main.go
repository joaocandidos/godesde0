package main

import (
	"fmt"

	"github.com/joaocandidos/godesde0/variables"
)

func main() {
	//variables.MostroEnteros()
	//variables.RestoVariables()

	estado, texto := variables.ConviertoaTexto(2345)
	fmt.Println(estado, texto)
}
