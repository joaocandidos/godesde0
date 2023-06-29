package ejercicios

import (
	"strconv"
)

func DoisValores(texto string) (int, string) {
	num, err := strconv.Atoi(texto)
	if err != nil {
		return 0, "Hubo un error" + err.Error()
	}
	if num > 100 {
		return num, "es mayor a 100"
	} else {
		return num, "es menor a 100"
	}

}
