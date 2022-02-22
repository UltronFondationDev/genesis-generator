package makegenesis

import (
	"math/big"
	"time"

	"github.com/Fantom-foundation/go-opera/inter"
	"github.com/Fantom-foundation/go-opera/opera"
)

const (
	EventGas uint64 = 28000
)

func CreateRules() opera.Rules {
	return opera.Rules{
		Name:      "ultron",
		NetworkID: 0x4ce,
		Dag: opera.DagRules{
			MaxParents:     10,
			MaxFreeParents: 3,
			MaxExtraData:   128,
		},
		Epochs: opera.EpochsRules{
			MaxEpochGas:      1500000000,                     // high?
			MaxEpochDuration: inter.Timestamp(4 * time.Hour), //inter.Timestamp(10 * time.Minute)
		},
		Economy: opera.EconomyRules{
			BlockMissedSlack: 50,
			MinGasPrice:      big.NewInt(1e9),
			Gas: opera.GasRules{
				MaxEventGas:  10000000 + EventGas,
				EventGas:     EventGas,
				ParentGas:    2400,
				ExtraDataGas: 25,
			},
			LongGasPower: opera.GasPowerRules{
				AllocPerSec:        100 * EventGas, // AllocPerSec:        100000 * EventGas
				MaxAllocPeriod:     inter.Timestamp(60 * time.Minute),
				StartupAllocPeriod: inter.Timestamp(5 * time.Second),
				MinStartupGas:      EventGas * 20,
			},
			ShortGasPower: opera.GasPowerRules{
				// 2x faster allocation rate, 6x lower max accumulated gas power
				AllocPerSec:        100 * EventGas * 2,
				MaxAllocPeriod:     inter.Timestamp(60*time.Minute) / (2 * 6),
				StartupAllocPeriod: inter.Timestamp(5*time.Second) / 2,
				MinStartupGas:      EventGas * 20,
			},
		},
		Blocks: opera.BlocksRules{
			MaxBlockGas:             20500000,
			MaxEmptyBlockSkipPeriod: inter.Timestamp(1 * time.Minute),
		},
	}
}
