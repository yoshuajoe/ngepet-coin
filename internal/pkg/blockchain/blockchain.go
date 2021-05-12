package blockchain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"ngepetcoin/internal/pkg/block"
	"ngepetcoin/pkg/milis"
	"os"
	"reflect"
)

type SimpleBlockChain struct {
	chain BlockChain
	nodes []string
}

func New(chain BlockChain, nodes []string) IBlockChain {
	return &SimpleBlockChain{
		chain: chain,
		nodes: nodes,
	}
}

func (bc *SimpleBlockChain) ProofOfWork(lastProof int64) (bool, int64) {
	inc := lastProof + 1
	startTime := milis.MakeTimestamp()

	for (inc%8020 != 0) && (inc%lastProof != 0) {
		inc++

		if (milis.MakeTimestamp()-startTime)%30 == 0 {
			isNewBlockChain, newBlockChain := bc.Consensus()
			if isNewBlockChain {
				bc.chain.Blocks = newBlockChain
				return false, -1
			}
		}
	}
	return true, inc
}

func (bc *SimpleBlockChain) FindNewChains() []BlockChain {
	other := []BlockChain{}
	for _, val := range bc.nodes {
		blockChain := fetchBlockChainFromURLPeer(val)
		if validateBlockChain(blockChain) {
			other = append(other, blockChain)
		}
	}
	return other
}

func (bc *SimpleBlockChain) Consensus() (bool, []block.Block) {
	otherChains := bc.FindNewChains()
	longest := bc.chain.Blocks

	for _, chain := range otherChains {
		if len(longest) < len(chain.Blocks) {
			longest = chain.Blocks
		}
	}

	if reflect.DeepEqual(longest, bc.chain.Blocks) {
		return false, []block.Block{}
	}

	return true, longest
}

func validateBlockChain(b BlockChain) bool {
	return true
}

func fetchBlockChainFromURLPeer(url string) BlockChain {
	response, responseErr := http.Get(url)

	if responseErr != nil {
		fmt.Print(responseErr.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	arrBlock := []block.Block{}
	jsonErr := json.Unmarshal(responseData, &arrBlock)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return BlockChain{
		Blocks: arrBlock,
	}
}
