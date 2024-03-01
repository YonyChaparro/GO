package main

import (
	"fmt"
	"reflect"
)

func main() {

	// VARIABLES:

	var myString string = "Esto es una cadena de texto"
	fmt.Println(myString)

	myString = "Cambiamos el valor de la cadena"

	fmt.Println(myString)

	var myInt int = 5
	myInt = myInt + 7
	fmt.Println(myInt)

	fmt.Println(myInt, myString)
	fmt.Println(reflect.TypeOf(myInt))

	var myFloat = 6.5
	println(myFloat)
	fmt.Println(reflect.TypeOf(myFloat))
	fmt.Println(myFloat, myInt)

	//Para datos Booleanos;

	var myBool bool=true

	println(myBool)
	// =: Variable e inicializada de manera abreviada :=
	variableAutomatica := "Esto es una cadena"
	println(variableAutomatica)

	//Constantes:

	const myConst = "Esto es una constante"

	println(myConst)

	//Controles de flujo

	if myInt == 12 && myString == "Cambiamos el valor de la cadena" {
		fmt.Println("el valor es 12")
	} else if myInt != 12{
		fmt.Println("El valor no es 12")
	}else{
		fmt.Println("No hay mas opciones")
	}
		
	}
}
