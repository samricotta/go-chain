package blockchain


type Sender struct {
	Name          string
	Address       string
	Balance       float64
	PrivateKey    []byte // This should be securely stored
	PublicKey     []byte
	Transactions  []*Transaction
}


func NewSender(name string, address string, balance float64, privateKey []byte, publicKey []byte) *Sender {
	sender := &Sender{
		Name:          name,
		Address:       address,
		Balance:       balance,
		PrivateKey:    privateKey,
		PublicKey:     publicKey,
	}
	return sender
}

