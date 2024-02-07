package punteros

import "fmt"

func PointersSample() {
	var p *int
	x := 10
	p = &x
	*p = 20

	fmt.Println(x)
	fmt.Println(*p)

}
