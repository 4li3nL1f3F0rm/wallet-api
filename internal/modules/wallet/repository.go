package wallet

import (
	"database/sql"
	"errors"
	"fmt"
)

type WalletRepository struct {
	DB *sql.DB
}

func NewWalletRepository(db *sql.DB) *WalletRepository {
	return &WalletRepository{DB: db}
}

func (r *WalletRepository) FindAll() ([]Wallet, error) {
	rows, err := r.DB.Query(FindAllWalletsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var wallets []Wallet
	for rows.Next() {
		var wallet Wallet
		if err := rows.Scan(&wallet.ID, &wallet.UserID, &wallet.Balance); err != nil {
			return nil, err
		}
		wallets = append(wallets, wallet)
	}
	return wallets, nil
}

func (r *WalletRepository) FindById(id int) (Wallet, error) {
	var wallet Wallet
	err := r.DB.QueryRow(FindWalletByIDQuery, id).
		Scan(&wallet.ID, &wallet.UserID, &wallet.Balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return Wallet{}, errors.New(fmt.Sprintf("wallet with id %d not found", id))
		}
		return Wallet{}, err
	}
	return wallet, nil
}
