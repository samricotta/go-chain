package blockchain

import (
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"math/big"
	"math/rand"
)

type Account struct {
	Name         string
	Address      string
	Balance      float64
	privateKey   string
	publicKey    string
	Transactions []*Transaction
}

func NewAccount(name string, address string, balance float64, privateKey string, publicKey string) *Account {
	account := &Account{
		Name:       name,
		Address:    address,
		Balance:    balance,
		privateKey: privateKey,
		publicKey:  publicKey,
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

func (a *Account) GeneratePublicKey() (string, error) {
	block, _ := pem.Decode([]byte(a.privateKey))
	if block == nil {
		return "", fmt.Errorf("failed to decode PEM block containing private key")
	}

	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	pubKey := append(privateKey.PublicKey.X.Bytes(), privateKey.PublicKey.Y.Bytes()...)
	return hex.EncodeToString(pubKey), nil
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

func IsValidPrivateKey(privateKey string) bool {

}
