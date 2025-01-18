package controllers

import (
	"github.com/amentobing/todo-app-go/database"
	"github.com/amentobing/todo-app-go/models"
	"github.com/gofiber/fiber/v2"
)
func GetTodos(c *fiber.Ctx) error {
    var todos []models.ToDo
    database.DB.Find(&todos)
    return c.JSON(todos)
}
func GetTodoByID(c *fiber.Ctx) error {
    id := c.Params("id")
    var todo models.ToDo
    result := database.DB.First(&todo, id)
    if result.Error != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Todo not found",
        })
    }
    return c.JSON(todo)
}
func CreateTodo(c *fiber.Ctx) error {
    todo := new(models.ToDo)
    if err := c.BodyParser(todo); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }
    database.DB.Create(&todo)
    return c.JSON(todo)
}
func UpdateTodo(c *fiber.Ctx) error {
    id := c.Params("id")
    var todo models.ToDo
    if result := database.DB.First(&todo, id); result.Error != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Todo not found",
        })
    }
    if err := c.BodyParser(&todo); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }
    database.DB.Save(&todo)
    return c.JSON(todo)
}
func DeleteTodo(c *fiber.Ctx) error {
    id := c.Params("id")
    var todo models.ToDo
    if result := database.DB.First(&todo, id); result.Error != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Todo not found",
        })
    }
    database.DB.Delete(&todo)
    return c.SendStatus(fiber.StatusNoContent)
}