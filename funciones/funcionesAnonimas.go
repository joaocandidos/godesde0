package funciones

import "fmt"

func Calculos() {
	suma := func(numero1 int, numero2 int) int {
		return numero1 + numero2
	}

	fmt.Println(suma(18, 12))
}
