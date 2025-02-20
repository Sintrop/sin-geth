package params

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/params/types/genesisT"
)

var SintropGenesisHash = common.HexToHash("0x5d035852d375b39b1eec893b7aef334c3d908999cfc9fd08de43de8160d4a42a")

func DefaultSintropGenesisBlock() *genesisT.Genesis {
	return &genesisT.Genesis{
		Config:     SintropChainConfig,
		Nonce:      hexutil.MustDecodeUint64("0x0"),
		ExtraData:  hexutil.MustDecode("0x42"),
		GasLimit:   hexutil.MustDecodeUint64("0x2fefd8"),
		Difficulty: hexutil.MustDecodeBig("0x20000000"),
		Timestamp:  1615385980,
		Alloc: genesisT.GenesisAlloc{
			common.HexToAddress("366ae7da62294427c764870bd2a460d7ded29d30"): genesisT.GenesisAccount{
				Balance: big.NewInt(42),
			},
		},
	}
}
