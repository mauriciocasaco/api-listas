package Listas

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Task struct {
	Checked     bool   `json:"checked"`
	Description string `json:"description"`
}

type List struct {
	Titulo string  `json:"titulo"`
	Tareas *[]Task `json:"tareas"`
}

// Crear una lista nueva
func CreateList(context *fiber.Ctx) error {
	lista := List{}
	err := context.BodyParser(&lista)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	if lista.Titulo == "" {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "falta titulo"})
		return errors.New("Falta titulo")
	}
	/*done, Validar informacion
	Crear objeto lista.
	Guardarlo en base de datos
	Responder con el objeto creado/*/
	fmt.Println("Se llamo exitosamente a la funcion crear de listas")
	context.Status(http.StatusOK).JSON(lista) //&fiber.Map{
	//"message": "Estamos constrruyendo la funcion", "title": lista.Titulo,

	return nil
}
