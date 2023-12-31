package blockchain

import (
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
	privateKey := "Test Private Key"
	publicKey := "Test Public Key"

	account := NewAccount(name, address, balance, privateKey, publicKey)

	return &TestFixture{
		account: account,
	}
}

func TestNewAccount(t *testing.T) {
	f := setup()

	require.Equal(t, "Test Name", f.account.Name)
	require.Equal(t, "Test Address", f.account.Address)
	require.Equal(t, 100.0, f.account.Balance)
	require.Equal(t, "Test Private Key", f.account.privateKey)
	require.Equal(t, "Test Public Key", f.account.publicKey)
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

func TestGeneratePrivateKey(t *testing.T) {
	privateKey, err := GeneratePrivateKey()
	require.NoError(t, err)
	validKey := IsValidKey(privateKey)
	require.True(t, validKey)
}

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

func TestIsValidKey(t *testing.T) {
	f := setup()
	validKey := IsValidKey(f.account.privateKey)
	require.True(t, validKey)
}
