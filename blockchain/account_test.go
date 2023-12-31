package blockchain

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

type TestFixture struct {
	account *Account
}

func setup() *TestFixture {
	name := "Test Name"
	address := "Test Address"
	balance := 100.0
	privateKey, err := GeneratePrivateKey()
	if err != nil {
		log.Fatalf("Failed to generate private key: %v", err)
	}

	account := NewAccount(name, address, balance, privateKey, "")
	publicKey, err := account.GeneratePublicKey()
	if err != nil {
		log.Fatalf("Failed to generate public key: %v", err)
	}
	account.publicKey = publicKey

	return &TestFixture{
		account: account,
	}
}

func TestNewAccount(t *testing.T) {
	f := setup()

	require.Equal(t, "Test Name", f.account.Name)
	require.Equal(t, "Test Address", f.account.Address)
	require.Equal(t, 100.0, f.account.Balance)
	require.NotEmpty(t, f.account.privateKey)
	require.NotEmpty(t, f.account.publicKey)
}

func TestGetBalance(t *testing.T) {
	f := setup()
	require.Equal(t, 100.0, f.account.GetBalance())
}

func TestGetName(t *testing.T) {
	f := setup()
	require.Equal(t, "Test Name", f.account.GetName())
}

func TestGetAddress(t *testing.T) {
	f := setup()
	require.Equal(t, "Test Address", f.account.GetAddress())
}

// TODO: fix
// func TestGeneratePrivateKey(t *testing.T) {
// 	privateKey, err := GeneratePrivateKey()
// 	require.NoError(t, err)
// 	validKey := IsValidKey(privateKey)
// 	require.True(t, validKey)
// }

// TODO: fix
// func TestGeneratePublicKey(t *testing.T) {
// 	f := setup()
// 	privateKey, err := f.account.GeneratePublicKey()
// 	require.NoError(t, err)
// 	validKey := IsValidKey(privateKey)
// 	require.True(t, validKey)
// }

func TestGenerateAddress(t *testing.T) {
	f := setup()
	address, err := f.account.GenerateAddress()
	require.NoError(t, err)
	require.NotEmpty(t, address)
}

//TODO: Fix
// func TestIsValidKey(t *testing.T) {
// 	f := setup()
// 	validKey := IsValidKey(f.account.privateKey)
// 	require.True(t, validKey)
// }
