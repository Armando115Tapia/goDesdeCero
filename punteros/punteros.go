package punteros

import "fmt"

func PointersSample() {
	var p *int
	x := 10
	p = &x
	*p = 20

	fmt.Println(x)
	fmt.Println(*p)
	// En go cuando queremos crear un puntero también podemos usar la palabra reservada new
	pointer := new(int)
	*pointer = 123
	fmt.Println(*pointer)

}
