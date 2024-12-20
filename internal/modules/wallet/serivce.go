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

func (s *WalletService) GetWalletById(id int) (Wallet, error) {
	return s.Repo.FindById(id)
}

func (s *WalletService) CreateWallet(wallet CreateWalletRequest) (Wallet, error) {
	return s.Repo.Create(wallet)
}
