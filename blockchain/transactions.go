package blockchain

import (
	"encoding/json"
	"fmt"
	"strings"

	blockchain "github.com/samricotta/go-chain/blockchain/v1"
)

type Transaction struct {
	Id        string
	Sender    string
	Reciever  string
	Amount    float64
	Signature string
	Timestamp string
}

type TransactionPool struct {
	Transactions []*Transaction
}

func NewTransaction(sender string, reciever string, amount float64, signature string) *Transaction {
	transaction := &Transaction{
		Sender:    sender,
		Reciever:  reciever,
		Amount:    amount,
		Signature: signature,
	}
	return transaction
}

func (t *Transaction) IsValidTransaction(tp *TransactionPool) bool {
	if t == nil {
		return false
	}

	if t.Sender == "" || t.Reciever == "" || t.Amount == 0 || t.Signature == "" {
		return false
	}

	if strings.TrimSpace(t.Sender) != "" {
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

	if t.IsDoubleSpend(tp) {
		fmt.Println("Transaction is a double spend")
		return false
	}
}

func (t *Transaction) Sign() string {

}

func (t *Transaction) AddTransactionToBlock(block *blockchain.Block) bool {
	if t.IsValidTransaction() {
		fmt.Println("Transaction is invalid")
		return false
	}
	if t.GetTransactionSize() > CalculateBlockSize(block) {
		fmt.Println("Transaction is too large")
		return false
	}

	mb, err := json.Marshal(t)
	if err != nil {
			return false
	}
	block.Data = append(block.Data, mb...)
	return true 
}

func (t *Transaction) GetTransactionByIndex(index int32) {
	return t[index]
}

func (tp *TransactionPool) AddToTransactionPool(t *Transaction) bool {
	if t.IsValidTransaction() {
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

func (t *Transaction) IsDoubleSpend(tp *TransactionPool, bc *Blockchain) bool {
	for _, tp := range tp.Transactions {
		if t.hasSameInput(tp) {
			return true
		}
	}

	for _, block := range bc.Blocks {
		var transactions []*Transaction
		err := json.Unmarshal(block.Data, &transactions)
		if err != nil {
			return false
		}
		for _, transaction := range transactions {
			if t.hasSameInput(transaction) {
				return true
			}
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

func (t *Transaction) GetTransactionSize() int32 {
	jsonBytes, err := json.Marshal(t)
	if err != nil {
		return err
	}
	length := len(jsonBytes)
	return int32(length)
}
