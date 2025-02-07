package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Contacto struct {
	Nombre string `json:"name"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
}

// saveContactsToFile Crea el archivo de contactos
func saveContactsToFile(contacts []Contacto) error {
	file, err := os.Create("contacts.json")
	if err != nil {
		return err
	}

	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(contacts)
	if err != nil {
		return err
	}

	return nil
}

func loadContactFromFile(contacts *[]Contacto) error {
	file, err := os.Open("contacts.json")
	if err != nil {
		return err
	}

	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var contacts []Contacto
	err := loadContactFromFile(&contacts)
	if err != nil {
		fmt.Println("Error al cargar los contactos: ", err)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("~~~ Gestor de Contactos ~~~\n",
			"1. Agregar Contacto \n",
			"2. Mostrar todos los contactos \n",
			"3. Salir \n",
			"Elige una opcion:",
		)

		var opcion int
		_, err := fmt.Scan(&opcion)
		if err != nil {
			fmt.Println("Error al leer la opcion")
			return
		}

		switch opcion {
		case 1:
			var c Contacto
			fmt.Println("Nombre:")
			c.Nombre, _ = reader.ReadString('\n')
			fmt.Println("Email:")
			c.Email, _ = reader.ReadString('\n')
			fmt.Println("Telefono:")
			c.Phone, _ = reader.ReadString('\n')
			contacts = append(contacts, c)
			if err := saveContactsToFile(contacts); err != nil {
				fmt.Println("Error al guadar el contacto: ", err)
			}

		case 2:
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
			for index, contac := range contacts {
				fmt.Printf(
					"%d. Nombre: %s Email: %s Telefono: %s",
					index,
					contac.Nombre,
					contac.Email,
					contac.Phone,
				)
			}
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

		case 3:
			return
		default:
			fmt.Println("Opcion invalida")
		}
	}
}
