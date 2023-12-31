package blockchain

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"math/big"

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

func GeneratePrivateKey() (string, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return "", fmt.Errorf("failed to generate private key: %v", err)
	}

	der, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to marshal private key: %v", err)
	}

	block := &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: der,
	}

	privatePEM := pem.EncodeToMemory(block)
	return string(privatePEM), nil
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
	upperLimit := new(big.Int).Lsh(big.NewInt(1), 256)
	return rand.Int(rand.Reader, upperLimit)
}

func IsValidKey(key string) bool {
	if !isHexString(key) {
		fmt.Println("The key is not a valid hexadecimal string")
		return false
	}
	if !isCorrectLength([]byte(key)) {
		fmt.Println("The key is not the correct length")
		return false
	}
	// if !isValidOnCurve([]byte(key)) {
	// 	fmt.Println("The key is not a valid public key")
	// 	return false
	// }
	return true
}

func isHexString(key string) bool {
	dst := make([]byte, hex.DecodedLen(len(key)))
	_, err := hex.Decode(dst, []byte(key))
	if err != nil {
		fmt.Print("The key is not a valid hexadecimal string")
		return false
	}
	return true
}

func isCorrectLength(keyBytes []byte) bool {
	// For example, using the P256 curve which expects 64 bytes for an uncompressed public key
	return len(keyBytes) == 2*elliptic.P256().Params().BitSize/8
}

// func isValidOnCurve(keyBytes []byte) bool {
// 	pubKey, err := x509.ParsePKIXPublicKey(keyBytes)
// 	if err != nil {
// 		return false
// 	}

// 	ecPubKey, ok := pubKey.(*ecdsa.PublicKey)
// 	if !ok {
// 		return false
// 	}

// 	// Check if the key is on the curve
// 	return ecPubKey.Curve.IsOnCurve(ecPubKey.X, ecPubKey.Y)
// }
