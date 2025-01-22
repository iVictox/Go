/*
################################################################################
Ejercicio 1
(Elaborar un programa que lea la cedula, nombre, venta y calcule el IVA (16%))

Ejercicio realizado por: Víctor Hugo De Abreu Pedrón (V- 30.839.445)
################################################################################
*/

package main

import "fmt"

func main() {
	// Variables
	var cedula int
	var nombre string
	var venta float64
	var iva float64

	// Proceso
	fmt.Println("############")
	fmt.Println("VENTAS - GO")
	fmt.Println("############")

	fmt.Println("* Ingrese cédula del cliente:")
	fmt.Scanln(&cedula)

	fmt.Println("* Ingrese nombre del cliente:")
	fmt.Scanln(&nombre)

	fmt.Println("* Ingrese el monto de la venta:")
	fmt.Scanln(&venta)

	iva = venta * 16 / 100 // Calcular IVA

	// Imprimir resultado
	fmt.Println("############################")
	fmt.Printf("* %-10s : %s\n", "Nombre", nombre)
	fmt.Printf("* %-10s : V-%d\n", "Cedula", cedula)
	fmt.Println("")
	fmt.Printf("* %-10s : %.2f Bs\n", "Venta", venta)
	fmt.Printf("* %-10s :  %.2f Bs\n", "IVA (16%)", iva)
	fmt.Printf("* %-10s : %.2f Bs\n", "Total", venta+iva)
	fmt.Println("############################")
}
