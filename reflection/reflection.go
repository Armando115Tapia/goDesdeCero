package reflection

import (
	"fmt"
	"reflect"
)

type Gopher struct {
	Name  string
	color string
	Year  int
}

func ReflectSample1() {
	g := Gopher{Name: "Adam", color: "Pink", Year: 1992}
	rType := reflect.TypeOf(g)
	rKind := rType.Kind()
	rValue := reflect.ValueOf(g)

	fmt.Println("Type: ==> ", rType)
	fmt.Println("Kind: ==> ", rKind)
	fmt.Println("Value: ==> ", rValue)
	name := rValue.Interface().(Gopher).Name
	fmt.Println("Name: ==> ", name)
}

func ReflectSample2() {
	g := Gopher{Name: "Jhon", color: "Yellow", Year: 123}
	gType := reflect.TypeOf(g)
	fmt.Printf("Type: %s kind %s \n", gType.Name(), gType.Kind())

	ptrG := &g
	ptrType := reflect.TypeOf(ptrG)
	fmt.Printf("Type: %s Kind: %s\n", ptrType.Elem(), ptrType.Kind())

}
