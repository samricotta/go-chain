package blockchain


type Transaction struct {
	Id  string
	Sender string
	Reciever string
	Amount float64
	Signature string
	Timestamp  string
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

	if IsDoubleSpend {
		fmt.println("Transaction is a double spend")
		return false
	}
}

func(t *transaction)Sign() string {
	
}

func(t *transaction)AddTransactionToBlock(block Block) {
	t.transaction.IsValidTransaction() {
		fmt.Println("Transaction is valid")
		block.Transaction = append(block.Transaction, t.transaction)
	}
}

func(t *transaction)GetTransactionByIndex(index int32) {
 return t.transaction[index]
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

func(t *Transaction)IsDoubleSpend(tp *TransactionPool) bool {
	for _, tp := range tp.Transactions{
		if t.hasSameInput(tp) {
			return true
		}
	}

	for _, block := range bc.Blocks {
		if t.hasSameInput(block.Transaction) {
			return true
		}
	}
}

func(t *Transaction)hasSameInput(ot OtherTransaction) bool {
	if t.Sender == ot.Sender && t.Reciever == ot.Reciever && t.Amount == ot.Amount {
		return true
	}
	return false
}

func(t *Transaction)GetTransactionSize() int {
	jsonBytes, err := json.Marshal(t.transaction)
	if err != nil {
		return err
	}
	length := len(jsonBytes)
	return length
}

