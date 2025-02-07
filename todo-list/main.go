package main

import (
	"bufio"
	"fmt"
	"os"
)

// Task estructura
type Task struct {
	titulo      string
	descripcion string
	estado      bool
}

type TaskList struct {
	tarea []Task
}

// addTask agrega tareas
func (l *TaskList) addTask(task Task) {
	l.tarea = append(l.tarea, task)
}

// completeTask completar tarea
func (l *TaskList) completeTask(index int) {
	l.tarea[index].estado = true
}

func (l *TaskList) editTask(index int, task Task) {
	l.tarea[index] = task
}

func (l *TaskList) deleteTask(index int) {
	l.tarea = append(l.tarea[:index], l.tarea[index+1:]...)
}

func main() {
	lista := TaskList{}

	leer := bufio.NewReader(os.Stdin)
	for {
		showTask(lista)
		var opciones int
		fmt.Println("Selecciona una Opcion: \n",
			"1. Agregar tarea\n",
			"2. Completar tarea\n",
			"3. Editar tarea\n",
			"4. Eliminar tarea\n",
			"5. Salir",
		)
		fmt.Scan(&opciones)

		switch opciones {
		case 1:
			var tarea Task
			fmt.Println("Ingresa el nombre de la tarea:")
			tarea.titulo, _ = leer.ReadString('\n')
			fmt.Println("Ingresa la descripcion de la tarea:")
			tarea.descripcion, _ = leer.ReadString('\n')
			lista.addTask(tarea)

		case 2:
			var index int
			fmt.Println("Ingrese el numero de la tarea que desea marcar como completada")
			fmt.Scan(&index)
			lista.completeTask(index)

		case 3:
			var index int
			var tarea Task
			fmt.Println("Ingresa el numero de la tarea que deseas editar")
			fmt.Scan(&index)
			fmt.Println("Ingresa el nombre de la tarea")
			tarea.titulo, _ = leer.ReadString('\n')
			fmt.Println("Ingresa la descripcion de la tarea:")
			tarea.descripcion, _ = leer.ReadString('\n')
			lista.editTask(index, tarea)
			fmt.Println("Tarea Editada")

		case 4:
			var index int
			fmt.Println("Ingresa el numero de la tarea que deseas eliminar")
			fmt.Scan(&index)
			lista.deleteTask(index)
			fmt.Println("Tarea borrada")

		case 5:
			fmt.Println("Saliendo del programa")
			return
		default:
			fmt.Println("Opcion invalida")
		}
	}

}

func showTask(lista TaskList) {

	fmt.Println("Lista de Tareas:")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	for index, task := range lista.tarea {
		fmt.Printf(
			"%d. %s - %s - Completado %t \n",
			index,
			task.titulo,
			task.descripcion,
			task.estado,
		)
	}

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

}
