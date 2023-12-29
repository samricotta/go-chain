package blockchain


type Transaction struct {
	Sender string
	Reciever string
	Amount float64
	Signature string
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

func(t *transaction)IsValid() bool {
	
}

func(t *transaction)Sign() string {
	
}

func(t *transaction)AddTransactionToBlock() {
	
}

func(t *transaction)GetTransactionByIndex() {

}

func(t *transaction)AddToTransactionPool() {

}