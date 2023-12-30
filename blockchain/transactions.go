package blockchain


type Transaction struct {
	Sender string
	Reciever string
	Amount float64
	Signature string
	Timestamp 
} 

type TransactionPool struct {
	Transactions []*Transaction
}


func NewTransaction(sender string, reciever string, amount float64, signature string) *Transaction{} {
	transaction := &Transaction{
		Sender: sender,
		Reciever: reciever,
		Amount: amount,
		Signature: signature,
	}
	return transaction
}

func(t *transaction)IsValidTransaction() bool {
	if transaction == nil || "" {
		return false
	}

	if transaction.Sender == "" || transaction.Reciever == "" || transaction.Amount == 0 || transaction.Signature == "" {
		return false
	}

	if strings.TrimSpace(transaction.Sender) != "" {
		fmt.println("Sender is not an empty string and not just whitespace")
		return false 
	}

	if strings.TrimSpace(transaction.Receiver) != "" {
		fmt.println("Receiver is not an empty string and not just whitespace")
		return false 
	}

	if strings.TrimSpace(transaction.Signature) != "" {
		fmt.println("Signature is not an empty string and not just whitespace")
		return false 
	}

	if t.Amount <= 0 {
		fmt.println("Amount is not greater than zero")
		return false
	}

	if t.Receiver {
}


}

func(t *transaction)Sign() string {
	
}

func(t *transaction)AddTransactionToBlock() {
	
}

func(t *transaction)GetTransactionByIndex() {

}

func(tp *TransactionPool)AddToTransactionPool(transaction *Transaction) bool {
	if transaction.isValidTransaction() {
		for _, trans := range tp.Transactions {
			if trans == transaction {
				return false
			}
		}
		t.Transaction :=  append(tp.Transactions, transaction)
		return true
	}
	return false
}

func BroadcastTransactionToNetwork(transaction){

}
