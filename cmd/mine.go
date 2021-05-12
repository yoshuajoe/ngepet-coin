package main

import (
	"ngepetcoin/internal/pkg/block"
	"ngepetcoin/internal/pkg/blockchain"
	"ngepetcoin/pkg/milis"
)

var PENDING_TRANSACTION = []block.BlockTransaction{}

func main() {
	blockChain := blockchain.BlockChain{}
	blockChain.Blocks = append(blockChain.Blocks, generateGenesisBlock())

	// blockchain.New(blockChain, []string{
	// 	"http://localhost:3000",
	// })

}

func generateGenesisBlock() block.Block {
	return block.Block{
		Index:     0,
		TimeStamp: milis.MakeTimestamp(),
		Data: block.BlockData{
			ProofOfWork:  9,
			Transactions: block.BlockTransaction{},
		},
		PrevHash: "0",
	}
}
