package blockchain


type Sender struct {
	Name          string
	Address       string
	Balance       float64
	PrivateKey    []byte // This should be securely stored
	PublicKey     []byte
	Transactions  []*Transaction
}


