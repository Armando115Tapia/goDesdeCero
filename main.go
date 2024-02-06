package main

import (
	"fmt"
	"github.com/Armando115Tapia/goDesdeCero/ejercicios"
)

func main() {

	//estado, valor := variables.ConviertoaTexto(14)
	//fmt.Println(estado)
	//fmt.Println(valor)
	//
	//if os := runtime.GOOS; os == "darwin" {
	//
	//	fmt.Println("Loco este Go, ya vamos a cachar las gorutinas")
	//}
	//
	//switch os := runtime.GOOS; os {
	//case "darwin":
	//	fmt.Println("Este de mac")
	//case "Linux.":
	//	fmt.Println("Este es de Linux")
	//
	//}

	valor, msj := ejercicios.ValidateIsGraterThan("PEPE")
	fmt.Println(valor)
	fmt.Println(msj)

}
