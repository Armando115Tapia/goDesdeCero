package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	Nombre = "Armando"
	Estado = true
	Sueldo = 1800
	Fecha = time.Now()

	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Fecha)
	fmt.Println(Sueldo)
}

func ConviertoaTexto(numero int) (bool, string) {
	text := strconv.Itoa(numero)

	return true, text
}
