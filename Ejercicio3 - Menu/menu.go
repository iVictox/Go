package main

import (
	"fmt"     // Entrada y salida
	"math"    // Funciones matemáticas
	"os"      // Sistema operativo
	"os/exec" // Ejecutar comandos externos
	"runtime" // Información del entorno
	"strings" // Manipulación de cadenas
)

/* Funciones */

func limpiarPantalla() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // Para Windows
	} else {
		cmd = exec.Command("clear") // Para Linux y otros sistemas Unix
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Potencia cuadradra de un numero
func potencia() {
	var base, resultado float64

	limpiarPantalla()
	fmt.Println("##############################")
	fmt.Println("POTENCIA CUADRADA DE UN NÚMERO")
	fmt.Println("##############################")

	fmt.Print("\n» Ingrese un número base: ")                               // Solicitar el número base
	fmt.Scan(&base)                                                         // Leer el número base
	resultado = math.Pow(base, 2)                                           // Ejecutar operación
	fmt.Printf("\n» El resultado de la operación es: %2.f \n\n", resultado) // Mostrar resultado
}

// Area de un trapecio
func areaTrapecio() {
	var baseMenor, baseMayor, altura, area float64

	limpiarPantalla()
	fmt.Println("###################")
	fmt.Println("AREA DE UN TRAPECIO")
	fmt.Println("###################")

	fmt.Print("\n» Ingrese cuanto es la base menor: ") // Solicitar la base menor
	fmt.Scan(&baseMenor)
	fmt.Print("\n» Ingrese cuanto es la base mayor: ") // Solicitar la base mayor
	fmt.Scan(&baseMayor)
	fmt.Print("\n» Ingrese cuanto es la altura: ") // Solicitar la altura
	fmt.Scan(&altura)

	area = ((baseMayor + baseMenor) / 2) * altura
	fmt.Printf("\n» El area del trapecio es: %.2f\n\n", area) // Mostrar resultado
}

// Calcular el volumen de una esfera
func volumenEsfera() {
	var radio, volumen float64

	limpiarPantalla()
	fmt.Println("#####################")
	fmt.Println("VOLUMEN DE UNA ESFERA")
	fmt.Println("#####################")

	fmt.Print("\n» Ingrese radio de la esfera: ") // Solicitar el radio
	fmt.Scan(&radio)

	volumen = (4.0 / 3.0 * math.Pi) * math.Pow(radio, 3)
	fmt.Printf("\n» El volumen de la esfera es: %.2f\n\n", volumen) // Mostrar resultado
}

func salir() {
	limpiarPantalla()
	fmt.Println()
	fmt.Println("##################################################")
	fmt.Println("Gracias por utilizar la aplicación de calculos")
	fmt.Println("Desarrollado por: Víctor Hugo De Abreu Pedrón")
	fmt.Println("##################################################")
	fmt.Println()
	os.Exit(0)
}

func operaciones() {
	var opcion int

	var continuar string
	for {
		fmt.Print("» Ingrese la opción que desea: ")
		fmt.Scan(&opcion)

		if opcion < 1 || opcion > 4 {
			fmt.Println("No ha seleccionado una opción correcta, vuélvalo a intentar.")
		} else {
			break
		}
	}
	switch opcion {
	// Potencia cuadrada de un número
	case 1:
		potencia()
	// Area de un trapecio
	case 2:
		areaTrapecio()
	// Volumen de una esfera
	case 3:
		volumenEsfera()
	// Salir
	case 4:
		salir()
	}

	// Mensaje de continuar
	for {
		fmt.Println("¿Desea continuar con el programa? (s/n)")
		fmt.Scan(&continuar)

		continuar = strings.ToLower(continuar)

		if continuar == "s" || continuar == "n" {
			if continuar == "n" {
				salir()
			} else {
				limpiarPantalla()
				break
			}
		} else {
			fmt.Println("No ha seleccionado una opción correcta, vuélvalo a intentar.")
		}
	}

}

/* Main */

func main() {
	limpiarPantalla()
	for {
		fmt.Println("##########")
		fmt.Println("   Menu")
		fmt.Println("##########")
		fmt.Println()
		fmt.Println("1) Potencia cuadrada de un número")
		fmt.Println("2) Area de un trapecio")
		fmt.Println("3) Volumen de esfera")
		fmt.Println("4) Salir")
		fmt.Println()
		// Leer, comprobar y ejecutar operaciones
		operaciones()

	}
}
