package wallet

type Wallet struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Balance     int    `json:"balance"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateWalletRequest struct {
	UserID      int    `bind:"required" json:"user_id"`
	Name        string `bind:"required" json:"name"`
	Balance     int    `json:"balance"`
	Description string `json:"description"`
}
