package main

import "fmt"

//Creamos la interfaz para calcular el área de tipo float64
type calcularArea interface {
	area() float64
}

//Creamos el struc del cuadrado
type cuadrado struct {
	lado float64
}

//Creamos el struct de rectangulo
type rectangulo struct {
	base   float64
	altura float64
}

//Creamos la funcion para calcular el area que recibe como parametro una interfaz
func calculo(f calcularArea) {
	fmt.Println("El area calculada es: ", f.area())
}

//Creamos la funcion calcular el área para el cuadrado
func (c cuadrado) area() float64 {
	return c.lado * c.lado
}

//creamos la funcion calcular el área para el rectangulo
func (r rectangulo) area() float64 {
	return r.base * r.altura
}

//Invocamos todo en el método main
func main() {
	myCuadrado := cuadrado{lado: 2}
	myRectangulo := rectangulo{base: 2, altura: 4}

	calculo(myCuadrado)
	calculo(myRectangulo)

	//Crear una lista de interfaces
	myInterfaz := []interface{}{"Hola", 5, 4}
	fmt.Println(myInterfaz)
}
