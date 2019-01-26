package run

import (
	"github.com/evm-lite/kernal"
)

type ETHChainHandler struct{}

func (ethChainHandler *ETHChainHandler) GetBlockHeaderHash(uint64) kernal.Hash {
	//just return a fake value
	return kernal.HexToHash("this is a demo")
}
