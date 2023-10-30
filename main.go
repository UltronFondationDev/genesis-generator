package main

import (
	"context"
	"fmt"
	"github.com/Fantom-foundation/go-opera/opera/genesisstore"
	"github.com/UltronFoundationDev/genesis-generator/configs"
	"github.com/UltronFoundationDev/genesis-generator/generator/genesis"
	"github.com/UltronFoundationDev/genesis-generator/generator/pkeys"
	"github.com/UltronFoundationDev/genesis-generator/pkeyreader"
	cli "github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

// TODO: need to change dependencies from opera to ultron when public repo will be available

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:     "count",
						Usage:    "fixes block transactions to this block",
						Required: true,
					},
				},
				Name:        "genkeys",
				Aliases:     nil,
				Description: "move events transactions to block transaction, required block_to flag to indicate which block should be fixed",
				Action: func(context *cli.Context) error {
					count := context.Int("count")
					var pkeyS []*pkeys.PKey

					for i := 1; i <= count; i++ {
						pkeysAndress, err := pkeys.GenerateKeyAndAddress()
						if err != nil {
							return cli.Exit(err, 1)
						}

						pkeyS = append(pkeyS, pkeysAndress)
					}

					folderName := fmt.Sprintf("./keys_%d", time.Now().Unix())
					if err := os.Mkdir(folderName, os.ModePerm); err != nil {
						return cli.Exit(err, 1)
					}

					for i, pkey := range pkeyS {
						iter := i
						filenamepubkey := fmt.Sprintf("%s/pubkey_%d.txt", folderName, iter)
						filenameprivkey := fmt.Sprintf("%s/privkey_%d.txt", folderName, iter)
						filenameaddr := fmt.Sprintf("%s/addr_%d.txt", folderName, iter)

						err := writeStringToFile(filenamepubkey, pkey.Public)
						if err != nil {
							return cli.Exit(err, 1)
						}

						err = writeStringToFile(filenameprivkey, pkey.Private)
						if err != nil {
							return cli.Exit(err, 1)
						}

						err = writeStringToFile(filenameaddr, pkey.Address)
						if err != nil {
							return cli.Exit(err, 1)
						}
					}

					return nil
				},
			},
			{
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "config",
						Usage:    "path to config file (you can use ready presets conf.testnet.toml or conf.devnet.toml",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "pks.path",
						Usage:    "path to file with private keys (see pks.example.txt)",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "out.filename",
						Usage:    "the name of genesis file, default: genesis.g",
						Required: false,
					},
				},
				Name:        "gengenesis",
				Aliases:     nil,
				Description: "generate genesis file from epoch to epoch",
				Action: func(ctx *cli.Context) error {
					cfgPath := ctx.String("config")
					pksPath := ctx.String("pks.path")
					outFilename := ctx.String("out.filename")

					var pkeyReader = pkeyreader.New(pksPath)
					pks, err := pkeyReader.GetPKS()
					if err != nil {
						return cli.Exit(err, 1)
					}

					cfg, err := configs.LoadConfig(cfgPath)
					if err != nil {
						return cli.Exit(err, 1)
					}

					g := genesis.New(cfg)
					gStore := g.CreateGenesisStore(context.TODO(), pks)

					if outFilename == "" {
						outFilename = "genesis.g"
					}

					fi, err := os.OpenFile(outFilename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
					if err != nil {
						return cli.Exit(err, 1)
					}
					defer fi.Close()

					err = genesisstore.WriteGenesisStore(fi, gStore)
					if err != nil {
						return cli.Exit(err, 1)
					}

					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func writeStringToFile(fName string, content string) error {
	fi, err := os.OpenFile(fName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer fi.Close()

	_, err = fi.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}
