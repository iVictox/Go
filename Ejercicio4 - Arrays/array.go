package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var resp string
	var numeros []int

	for {
		var numero int

		// Solicitar un número al usuario
		fmt.Println("Ingrese un número:")
		fmt.Scan(&numero)
		numeros = append(numeros, numero)

		// Preguntar si desea agregar otro número
		for {
			fmt.Println("¿Desea agregar otro número? (s/n)")
			fmt.Scan(&resp)

			resp = strings.ToLower(resp)

			if resp == "s" || resp == "n" {
				break
			} else {
				fmt.Println("No ha seleccionado una opción correcta, vuélvalo a intentar.")
			}
		}

		if resp == "n" {
			break
		}
	}

	// Ordenar los números de menor a mayor
	sort.Ints(numeros)
	fmt.Println("Números ordenados de menor a mayor:", numeros)

	// Ordenar los números de mayor a menor
	sort.Sort(sort.Reverse(sort.IntSlice(numeros)))
	fmt.Println("Números ordenados de mayor a menor:", numeros)
}
