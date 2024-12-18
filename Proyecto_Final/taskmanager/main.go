package main

import (
	"fmt"
)

type Task struct {
	ID          int
	Title       string
	Description string
	Completed   bool
}

var tasks = []Task{}
var idCounter int

func main() {
	for {
		fmt.Println("Task Manager")
		fmt.Println("1. Crear Tarea")
		fmt.Println("2. Ver Tareas")
		fmt.Println("3. Actualizar Tarea")
		fmt.Println("4. Eliminar Tarea")
		fmt.Println("5. Salir")
		fmt.Println("Seleccione una opción: ")

		var choice int
		fmt.Scan(&choice)
		fmt.Println(choice)

		switch choice {
		case 1:
			CreateTask()
		case 2:
			ListTask()
		case 3:
			UpdateTask()
		case 4:
			DeleteTask()
		case 5:
			fmt.Println("Hasta luego!")
			return
		default:
			fmt.Println("La opción no es valida.")
		}

	}
}

func DeleteTask() {
	var id int
	fmt.Println("Ingrese el ID de la tarea: ")
	fmt.Scan(&id)
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:id], tasks[i+1:]...)
			fmt.Println("Tarea eliminada!")
			return
		}

	}
}

func UpdateTask() {
	var id int
	fmt.Println("Ingrese el ID de la tarea: ")
	fmt.Scan(&id)

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = !tasks[i].Completed
			fmt.Println("Tarea actualizada!")
			return
		}

	}
	fmt.Println("Tarea no encontrada")
}

func ListTask() {
	if len(tasks) == 0 {
		fmt.Println("No hay tareas")
		return
	}

	fmt.Println("Lista de tareas:")
	for _, task := range tasks {
		status := "Pendiente"
		if task.Completed {
			status = "Completada"
		}
		fmt.Printf("%d - %s - %s - %s\n", task.ID, task.Title, task.Description, status)
	}
}

func CreateTask() {
	var title, description string
	fmt.Println("Ingrese el titulo de la tarea: ")
	fmt.Scan(&title)
	fmt.Println("Ingrese la descripción de la tarea: ")
	fmt.Scan(&description)

	idCounter++
	newTask := Task{
		ID:          idCounter,
		Title:       title,
		Description: description,
		Completed:   false,
	}

	tasks = append(tasks, newTask)
	fmt.Println("La tarea se ha creado con éxito!")

}
