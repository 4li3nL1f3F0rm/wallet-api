package wallet

type WalletService struct {
	Repo *WalletRepository
}

func NewWalletService(repo *WalletRepository) *WalletService {
	return &WalletService{Repo: repo}
}

func (s *WalletService) GetAllWallets() ([]Wallet, error) {
	return s.Repo.FindAll()
}
