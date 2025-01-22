/*
################################################################################
                                    Ejercicio 1
(Elaborar un programa que lea la cedula, nombre, venta y calcule el IVA (16%))
################################################################################
*/

package main

import "fmt"

func main() {
	// Variables
	var cedula int
	var nombre string
	var venta float32
	var iva float32

	// Proceso
	fmt.Println("* Ingrese cédula del cliente:")
	fmt.Scanln(&cedula)

	fmt.Println("* Ingrese nombre del cliente:")
	fmt.Scanln(&nombre)

	fmt.Println("* Ingrese el monto de la venta:")
	fmt.Scanln(&venta)

	iva = venta * 16 / 100 // Calcular IVA

	// Imprimir resultado
	fmt.Println("############################")
	fmt.Println("* Nombre: ", nombre)
	fmt.Println("* Cedula: ", cedula)
	fmt.Println("* Venta: ", venta, " Bs")
	fmt.Println("* IVA: ", iva, " Bs")
	fmt.Println("")
	fmt.Println("* Precio final: ", venta+iva, " Bs")
	fmt.Println("############################")
}
