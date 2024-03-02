package main

import (
	"container/list"
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

	var myBool bool = true

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
	} else if myInt != 12 {
		fmt.Println("El valor no es 12")
	} else {
		fmt.Println("No hay mas opciones")
	}

	//Array

	var myArray [3]int

	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	fmt.Println(myArray)
	fmt.Println(myString, myArray[2])

	var myArrayString [4]string

	myArrayString[0] = "Hola"
	myArrayString[1] = "Go"
	myArrayString[2] = "!!!"
	fmt.Println(myArrayString)
	fmt.Println(myString, myArrayString[2])

	//Mapas

	myMap := make(map[string]int)

	myMap["Yony"] = 23
	myMap["Felipe"] = 21
	myMap["Fercho"] = 26

	fmt.Println(myMap)
	fmt.Println(myMap["Yony"]) //Mostramos un elemento del mapa
	
	//Otra manera de crear un mapa

	myMap2 := map[string]int{"yony":3228237649, "Maritza": 3225678909, "Edgar": 3166899748}
	fmt.Println(myMap2)

	//Haciendo Listas: Se manejan un  poco diferente

	myList := list.New()
	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)
	fmt.Println(myList.Back().Value)

	//Bucles:

for i := 0; i < len(myArray); i++ {
	fmt.Println(myArray[i])
}

//Mostramos la informacion de un objeto

for key, value:= range myMap{
	fmt.Println(key, value)
}


//Estructura ===== Clases 

type MyStruct struct {
	nombre string
	edad int
}


myStruct0 := MyStruct{"Yony", 23}
fmt.Println(myStruct0)


//Función en Go: Las construimos afuera de esta función

myFuction()
fmt.Println(myFuction2())

//Fin del punto de entrada main()
}

func myFuction()  {
	fmt.Println("Esto lo imprimió un función ")
}

func myFuction2() string {
	return "retornamos un texto en la funcion2"
}




