package wallet

import "github.com/gofiber/fiber/v2"

type WalletHandlers struct {
	Service *WalletService
}

func NewWalletHandlers(service *WalletService) *WalletHandlers {
	return &WalletHandlers{Service: service}
}

func (h *WalletHandlers) GetAllWallets(c *fiber.Ctx) error {
	wallets, err := h.Service.GetAllWallets()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(wallets)
}

func (h *WalletHandlers) GetWalletById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid id"})
	}

	wallet, err := h.Service.GetWalletById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(wallet)
}
