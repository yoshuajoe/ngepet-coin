package block

// Block is a struct of block within blockchain
type Block struct {
	Index     int64
	TimeStamp int64
	Data      BlockData
	PrevHash  string
}

// BlockData is struct of Data inside Block
type BlockData struct {
	ProofOfWork  int64
	Transactions BlockTransaction
}

// BlockTransaction is struct of Transactions inside BlockData
type BlockTransaction struct {
	From   string
	To     string
	Amount float64
}
