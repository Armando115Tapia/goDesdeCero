package recursion

import "fmt"

func ExponenciaDeDos(val int) {
	if val > 10000 {
		return
	}
	fmt.Println(val)
	ExponenciaDeDos(val * 2)

}

var acc int

func ExponenciaDeNvalor(val, exp int) int {
	acc := 1
	acc = acc * val
	if exp == 1 {
		return acc
	}

	//exp_aux := exp
	//fmt.Println("Exp", exp_aux)
	//fmt.Println("Acc", acc)
	return ExponenciaDeNvalor(val, exp-1)

}

//func ExponenciaDeNvalorGTP(val, exp int) int {
//	if exp == 0 {
//		return 1
//	}
//	return val * ExponenciaDeNvalor(val, exp-1)
//}
