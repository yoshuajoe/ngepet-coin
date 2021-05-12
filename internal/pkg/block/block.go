package block

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

type SimpleBlock struct {
	block Block
	hash  string
}

func New(newBlock Block) IBlock {
	return &SimpleBlock{
		block: newBlock,
	}
}

func (sb *SimpleBlock) HashBlock() (string, error) {
	dataStr, convertError := json.Marshal(sb.block.Data)
	if convertError != nil {
		return "", convertError
	}

	hasher := sha256.New()
	blockStr := fmt.Sprintf("%d-%d-%s-%s", sb.block.Index, sb.block.TimeStamp, dataStr, sb.block.PrevHash)
	hasher.Write([]byte(blockStr))

	h := string(hasher.Sum(nil))
	sb.hash = h
	return h, nil
}
