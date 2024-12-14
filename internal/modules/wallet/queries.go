package wallet

const (
	FindAllWalletsQuery = `
		SELECT wallet_id, user_id, balance
		FROM wallets
	`

	FindWalletByIDQuery = `
		SELECT wallet_id, user_id, balance
		FROM wallets
		WHERE wallet_id = $1
	`
)
