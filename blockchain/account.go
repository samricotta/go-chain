package blockchain

import (
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"math/big"
	"math/rand"

	"github.com/cosmos/btcutil/bech32"
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

func (a *Account) GenerateAddress() (string, error) {
	pubKeyBytes, err := hex.DecodeString(a.publicKey)
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256(pubKeyBytes)


	// convert the first 20 bytes of the hash to a bech32 string
	// each byte is 8 bits, we then convert 8 bits to 5 bits
	//
	//The pad argument determines how the function should handle 
	//a situation where the input data does not contain a whole number of groups
	address, err := bech32.ConvertBits(hash[:20], 8, 5, true)
	if err != nil {
		return "", err
	}

	bech32Address, err := bech32.Encode("bc", address)
	if err != nil {
		return "", err
	}

	return bech32Address, nil
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
