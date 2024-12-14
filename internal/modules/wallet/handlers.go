package wallet

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

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
		if err.Error() == fmt.Sprintf("wallet with id %d not found", id) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(wallet)
}

func (h *WalletHandlers) CreateWallet(c *fiber.Ctx) error {
	var input CreateWalletRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	wallet, err := h.Service.CreateWallet(input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(wallet)
}
