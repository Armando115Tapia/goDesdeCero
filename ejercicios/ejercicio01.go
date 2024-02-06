package ejercicios

import "strconv"

func ValidateIsGraterThan(value string) (int, string) {

	valor, err := strconv.Atoi(value)

	if err != nil {
		return 0, "El valor ingresado no es un número"
	}

	if valor > 100 {
		return valor, "Es mayor que 100"
	} else {
		return valor, "Es menor o igual que 100"
	}

}
