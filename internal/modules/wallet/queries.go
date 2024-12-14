package wallet

const (
	FindAllWalletsQuery = `
		SELECT wallet_id, user_id, balance, name, description
		FROM wallets
	`

	FindWalletByIDQuery = `
		SELECT wallet_id, user_id, balance, name, description
		FROM wallets
		WHERE wallet_id = $1
	`

	CreateWalletQuery = `
		INSERT INTO wallets (user_id, balance, name, description)
		VALUES ($1, $2, $3, $4)
		RETURNING wallet_id, user_id, balance, name, description
	`
)
