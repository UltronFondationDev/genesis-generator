package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/Fantom-foundation/go-opera/opera/genesisstore"
	"github.com/UltronFoundationDev/genesis-generator/configs"
	"github.com/UltronFoundationDev/genesis-generator/generator"
	"github.com/UltronFoundationDev/genesis-generator/pkeyreader"
	"log"
	"os"
)

// TODO: need to change dependencies from opera to ultron when public repo will be available

func main() {
	var cfgPath string
	var pksPath string
	var outFilename string

	flag.StringVar(&cfgPath, "config", "", "config path")
	flag.StringVar(&pksPath, "pks.path", "", "path of the file with public keys")
	flag.StringVar(&outFilename, "out.filename", "genesis.g", "the name of the file to write the genesis to")

	flag.Parse()

	if pksPath == "" {
		log.Fatal("--pks.path is required")
	}
	if cfgPath == "" {
		log.Fatal("--config is required")
	}

	cfg, err := configs.LoadConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	var pkeyReader = pkeyreader.New(pksPath)
	pks, err := pkeyReader.GetPKS()
	if err != nil {
		log.Fatal(err)
	}

	g := generator.New(cfg)
	gStore := g.CreateGenesisStore(context.TODO(), pks)
	fi, err := os.OpenFile(outFilename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Println("Error in Create")
		panic(err)
	}

	err = genesisstore.WriteGenesisStore(fi, gStore)
	if err != nil {
		panic(err)
	}
}
