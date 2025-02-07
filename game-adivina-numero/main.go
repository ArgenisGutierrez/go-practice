package main

import (
	"fmt"
	"math/rand"
)

// main Juego adivina el numero
func main() {
	fmt.Println("Bienvenido a Adivina el numero.")
	fmt.Println("Tendras 10 oportunidades de adivinar un numero aleatorio entre 1 y 100")
	fmt.Println("Te gustaria jugar: (y/n)")
	var opcion string
	fmt.Scan(&opcion)
	switch opcion {
	case "y":
		juego()
	default:
		fmt.Println("Entiendo, intentenmoslo en otro momento")
	}

}

func juego() {
	numeroAleatorio := rand.Intn(100)
	var numeroIngresado int

	for i := 1; i <= 10; i++ {
		fmt.Println("Intento numero :", i)
		fmt.Scan(&numeroIngresado)

		if numeroIngresado == numeroAleatorio {
			fmt.Println("Felicidades acertaste.")
			break
		} else if i != 10 {
			fmt.Println("Error intentalo de nuevo")
			fmt.Println("Pista el numero es: ", pista(numeroAleatorio, numeroIngresado))
		} else {
			fmt.Println("Perdiste el numero era: ", numeroAleatorio)
		}
	}

	fmt.Println("Fin del juego")
	fmt.Println("Te gustatia jugar de nuevo (y/n)")
	var opcion string
	fmt.Scan(&opcion)
	if opcion == "y" {
		juego()
	} else {
		fmt.Println("Gracias por jugar")
	}
}

func pista(secretNumber, numero int) string {
	var mensaje string
	if numero < secretNumber {
		mensaje = "Mas grande"
	} else {
		mensaje = "Mas chico"
	}
	return mensaje
}
