package blockchain

import "ngepetcoin/internal/pkg/block"

type IBlockChain interface {
	ProofOfWork(int64) (bool, int64)
	Consensus() (bool, []block.Block)
}
