package main

import (
	"fmt"
	"math/big"
	"os"

	"github.com/Fantom-foundation/go-opera/opera/genesisstore"
	"testnet_genesis/makegenesis"
)

func main() {
	// const validatorBalance = uint64(1000000 * 1e18)
	validatorStakedAmt := new(big.Int)
	validatorStakedAmt.SetString("100000000000000000000000000", 10) //100 millions ulx * 2 validators

	validatorBalance := new(big.Int)
	validatorBalance.SetString("10000000000000000000", 10) //10 ulx

	gStore := makegenesis.CreateGenesisStore(validatorBalance, validatorStakedAmt)

	fi, err := os.OpenFile("genesis.g", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Println("Error in Create")
		panic(err)
	}

	genesisstore.WriteGenesisStore(fi, gStore)

}
