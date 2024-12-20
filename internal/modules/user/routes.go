package user

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, db *sql.DB, routes fiber.Router) {
	userRepo := NewUserRepository(db)
	userService := NewUserService(userRepo)
	userHandlers := NewUserHandlers(userService)

	routes.Get("/users", userHandlers.GetAllUsers)
	routes.Get("/users/:id", userHandlers.GetUser)
	routes.Post("/users", userHandlers.CreateUser)

}
