package blockchain

import (
	"math/big"
	"math/rand"
)

type Account struct {
	Name         string
	Address      string
	Balance      float64
	private_key  string
	public_key   string
	Transactions []*Transaction
}

func NewAccount(name string, address string, balance float64, private_key string, public_key string) *Account {
	account := &Account{
		Name:        name,
		Address:     address,
		Balance:     balance,
		private_key: private_key,
		public_key:  public_key,
	}
	return account
}

func (a *Account) GetBalance() float64 {
	return a.Balance
}

func (a *Account) GetName() string {
	return a.Name
}

func (a *Account) GetAddress() string {
	return a.Address
}

func (a *Account) GetTransactions() []*Transaction {
	return a.Transactions
}

func GeneratePrivateKey() (*big.Int, error) {
	privateKey, err := generateRandomNumber()
	if err != nil {
		return nil, err
	}

	for privateKey.Cmp(big.NewInt(0)) == 0 {
		privateKey, err = generateRandomNumber()
		if err != nil {
			return nil, err
		}
	}

	return privateKey, nil
}
func (a *Account) GeneratePublicKey() string {
}

func (a *Account) GenerateAddress() string {
}

func generateRandomNumber() (*big.Int, error) {
	upperLimit := new(big.Int).Lsh(big.NewInt(1), 256).Int64()

	privateKey := rand.Int63n(upperLimit)

	for privateKey == 0 {
		privateKey = rand.Int63n(upperLimit)
	}

	return big.NewInt(privateKey), nil
}


 func IsValidPrivateKey(*Account.privateKey) {

 }