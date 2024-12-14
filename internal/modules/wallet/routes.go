package wallet

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func WalletRoutes(app *fiber.App, db *sql.DB, routes fiber.Router) {

	walletRepo := NewWalletRepository(db)
	walletService := NewWalletService(walletRepo)
	walletHandlers := NewWalletHandlers(walletService)

	routes.Get("/wallets", walletHandlers.GetAllWallets)
	routes.Get("/wallets/:id", walletHandlers.GetWalletById)
	routes.Post("/wallets", walletHandlers.CreateWallet)

}
