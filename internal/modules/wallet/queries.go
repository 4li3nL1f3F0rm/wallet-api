package wallet

const (
	FindAllWalletsQuery = `
		SELECT wallet_id, user_id, balance
		FROM wallets
	`
)
