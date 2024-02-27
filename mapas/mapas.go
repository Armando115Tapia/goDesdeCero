package mapas

import "fmt"

func MostrarMapas() {
	// Asignarsion de valores
	paises := make(map[string]string)
	fmt.Println(paises)
	paises["Mexico"] = "D.F"
	paises["Argentina"] = "BBAA"
	paises["Ecuador"] = "Quito"
	fmt.Println(paises)
	fmt.Println(paises["Ecuador"])

	// Asignacion directa
	campeonato := map[string]int{
		"Barcelona":   12,
		"Real Madrid": 19,
	}

	fmt.Println(campeonato)

	// Los mapas en un range retornan la clave y el valor

	for equipo, puntaje := range campeonato {
		fmt.Printf("El equipo %s tiene un puntaje de %d \n", equipo, puntaje)
	}

	// Para quitar un elemento
	delete(campeonato, "Barcelona")

	fmt.Println(campeonato)

	// Para buscar
	puntaje, isExist := campeonato["Real MadridA"]
	fmt.Println("Puntaje:  ", puntaje)
	fmt.Println("Existe:  ", isExist)

}
