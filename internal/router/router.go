package router

import (
	"database/sql"

	"github.com/4li3nL1f3F0rm/wallet-api/internal/modules/user"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db *sql.DB) {
	routes := app.Group("/api")
	user.UserRoutes(app, db, routes)

}
