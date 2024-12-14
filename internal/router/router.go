package router

import (
	"database/sql"

	"github.com/4li3nL1f3F0rm/wallet-api/internal/modules/user"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db *sql.DB) {
	userRepo := user.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	userHandlers := user.NewUserHandlers(userService)

	userRoutes := app.Group("/users")
	userRoutes.Get("/", userHandlers.GetAllUsers)
	userRoutes.Get("/:id", userHandlers.GetUser)
	userRoutes.Post("/", userHandlers.CreateUser)

}
