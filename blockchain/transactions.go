package blockchain

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"strings"

	blockchain "github.com/samricotta/go-chain/blockchain/v1"
)

type Transaction struct {
	Id        string
	Sender    *Sender
	Reciever  string
	Amount    float64
	Signature string
	Timestamp string
}

type TransactionPool struct {
	Transactions []*Transaction
}

func NewTransaction(sender *Sender, reciever string, amount float64, signature string) *Transaction {
	transaction := &Transaction{
		Sender:    sender,
		Reciever:  reciever,
		Amount:    amount,
		Signature: signature,
	}
	return transaction
}

func (t *Transaction) IsValidTransaction(tp *TransactionPool, bc *blockchain.Block) bool {
	if t == nil {
		return false
	}

	if t.Sender == nil || t.Reciever == "" || t.Amount == 0 || t.Signature == "" {
		return false
	}

	if strings.TrimSpace(t.Sender.Name) != "" {
		fmt.Println("Sender is not an empty string and not just whitespace")
		return false
	}

	if strings.TrimSpace(t.Reciever) != "" {
		fmt.Println("Receiver is not an empty string and not just whitespace")
		return false
	}

	if strings.TrimSpace(t.Signature) != "" {
		fmt.Println("Signature is not an empty string and not just whitespace")
		return false
	}

	if t.Amount <= 0 {
		fmt.Println("Amount is not greater than zero")
		return false
	}

	if t.IsDoubleSpend(tp, bc) {
		fmt.Println("Transaction is a double spend")
		return false
	}
}

func (t *Transaction) Sign() ([]byte, error) {
	jb, err := json.Marshal(t)
	if err != nil {
		return []byte{}, err
	}

	calculateTransactionHash(jb)
	pk, err := ConvertPrivateKey(t.Sender.PrivateKey)
	if err != nil {
		return []byte{}, err
	}
	signed, err := SignTransaction(jb, pk)
	if err != nil {
		return []byte{}, err
	}
	return signed, nil
}

func (t *Transaction) AddTransactionToBlock(tp *TransactionPool, blockchain *Blockchain) bool {
	for _, block := range blockchain.Blocks {
    	if !t.IsValidTransaction(tp, block) {
            fmt.Println("Transaction is invalid")
            return false
        }

		size, err := t.GetTransactionSize()
		if err != nil {
			fmt.Println("Error getting transaction size:", err)
			return false
		}
		if size > blockchain.CalculateBlockSize() {
			fmt.Println("Transaction is too large")
			return false
		}

		mb, err := json.Marshal(t)
		if err != nil {
			return false
		}
		block.Data = append(block.Data, mb...)
	}
	return true
}

func (tp *TransactionPool) GetTransactionByIndex(index int32) (*Transaction, error) {
	if index < 0 || index >= int32(len(tp.Transactions)) {
		return nil, fmt.Errorf("Index out of range")
	}
	return tp.Transactions[index], nil
}

func (tp *TransactionPool) AddToTransactionPool(t *Transaction) bool {
	if t.IsValidTransaction(tp) {
		for _, trans := range tp.Transactions {
			if trans == t {
				return false
			}
		}
		tp.Transactions = append(tp.Transactions, t)
		return true
	}
	return false
}

func BroadcastTransactionToNetwork() {

}

func (t *Transaction) IsDoubleSpend(tp *TransactionPool, bc *blockchain.Block) bool {
	for _, tp := range tp.Transactions {
		if t.hasSameInput(tp) {
			return true
		}
	}

	var transactions []*Transaction
	err := json.Unmarshal(bc.Data, &transactions)
	if err != nil {
		return false
	}
	for _, transaction := range transactions {
		if t.hasSameInput(transaction) {
			return true
		}
	}
	return false
}

func (t *Transaction) hasSameInput(ot *Transaction) bool {
	if t.Sender == ot.Sender && t.Reciever == ot.Reciever && t.Amount == ot.Amount {
		return true
	}
	return false
}

func (t *Transaction) GetTransactionSize() (int32, error) {
	jsonBytes, err := json.Marshal(t)
	if err != nil {
		return -1, err
	}
	length := len(jsonBytes)
	return int32(length), nil
}

func calculateTransactionHash(record []byte) string {
	h := sha256.New()
	h.Write(record)
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func SignTransaction(transactionHash []byte, privateKey *ecdsa.PrivateKey) (signature []byte, err error) {
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, transactionHash)
	if err != nil {
		return nil, err
	}
	signature = append(r.Bytes(), s.Bytes()...)
	return signature, nil
}

func VerifySignature(publicKey *ecdsa.PublicKey, transactionHash []byte, signature []byte) bool {
	var (
		r = big.NewInt(0).SetBytes(signature[:len(signature)/2])
		s = big.NewInt(0).SetBytes(signature[len(signature)/2:])
	)

	// Verify the signature
	return ecdsa.Verify(publicKey, transactionHash, r, s)
}

func ConvertPrivateKey(privateKeyBytes []byte) (*ecdsa.PrivateKey, error) {
	block, _ := pem.Decode(privateKeyBytes)
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}
