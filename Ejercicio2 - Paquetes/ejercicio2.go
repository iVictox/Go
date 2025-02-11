package main

import (
	"bufio"   // Paquete para leer entradas del usuario
	"fmt"     // Paquete para imprimir en la consola
	"os"      // Paquete para interactuar con el sistema operativo
	"strconv" // Paquete para conversiones de tipos de datos
	"strings" // Paquete para manipulación de cadenas de texto
)

func main() {
	// Crear un lector para leer las entradas del usuario desde la consola
	reader := bufio.NewReader(os.Stdin)

	// Ejemplo de Atoi: Convierte una cadena de texto en un entero
	fmt.Print("Introduce un número entero (Atoi): ") // Solicitar al usuario que introduzca un número
	str, _ := reader.ReadString('\n')                // Leer la entrada del usuario
	str = strings.TrimSpace(str)                     // Eliminar los caracteres de espacio en blanco (incluyendo saltos de línea)
	num, err := strconv.Atoi(str)                    // Convertir la cadena en un entero
	if err != nil {
		// Si hay un error en la conversión, imprimir el error
		fmt.Println("Error:", err)
	} else {
		// Si la conversión es exitosa, imprimir el número
		fmt.Println("Número (Atoi):", num)
	}

	// Ejemplo de Itoa: Convierte un entero en una cadena de texto
	fmt.Print("Introduce un número entero para convertir a cadena (Itoa): ")
	str, _ = reader.ReadString('\n')
	str = strings.TrimSpace(str)
	number, err := strconv.Atoi(str) // Convertir la cadena en un entero para demostrar Itoa
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		strNum := strconv.Itoa(number)        // Convertir el entero en una cadena de texto
		fmt.Println("Cadena (Itoa):", strNum) // Imprimir la cadena resultante
	}

	// Ejemplo de ParseInt: Convierte una cadena hexadecimal en un entero
	fmt.Print("Introduce una cadena hexadecimal para convertir a entero (ParseInt): ")
	hexStr, _ := reader.ReadString('\n')
	hexStr = strings.TrimSpace(hexStr)
	parsedInt, err := strconv.ParseInt(hexStr, 16, 64) // Convertir la cadena hexadecimal en un entero
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Número (ParseInt):", parsedInt) // Imprimir el número resultante
	}

	// Ejemplo de FormatInt: Convierte un entero en una cadena hexadecimal
	fmt.Print("Introduce un número entero para convertir a hexadecimal (FormatInt): ")
	str, _ = reader.ReadString('\n')
	str = strings.TrimSpace(str)
	intVal, err := strconv.ParseInt(str, 10, 64) // Convertir la cadena en un entero
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		formattedStr := strconv.FormatInt(intVal, 16)    // Convertir el entero en una cadena hexadecimal
		fmt.Println("Cadena (FormatInt):", formattedStr) // Imprimir la cadena resultante
	}

	// Ejemplo de ParseFloat: Convierte una cadena de texto en un número de punto flotante
	fmt.Print("Introduce un número de punto flotante (ParseFloat): ")
	floatStr, _ := reader.ReadString('\n')
	floatStr = strings.TrimSpace(floatStr)
	parsedFloat, err := strconv.ParseFloat(floatStr, 64) // Convertir la cadena en un número de punto flotante
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Número (ParseFloat):", parsedFloat) // Imprimir el número resultante
	}

	// Ejemplo de FormatFloat: Convierte un número de punto flotante en una cadena de texto
	fmt.Print("Introduce un número de punto flotante para convertir a cadena (FormatFloat): ")
	str, _ = reader.ReadString('\n')
	str = strings.TrimSpace(str)
	floatVal, err := strconv.ParseFloat(str, 64) // Convertir la cadena en un número de punto flotante
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		formattedFloatStr := strconv.FormatFloat(floatVal, 'f', 2, 64) // Convertir el número de punto flotante en una cadena de texto con formato
		fmt.Println("Cadena (FormatFloat):", formattedFloatStr)        // Imprimir la cadena resultante
	}
}
