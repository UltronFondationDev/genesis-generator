package main

import (
	"context"
	"github.com/Fantom-foundation/go-opera/opera/genesisstore"
	"github.com/UltronFoundationDev/genesis-generator/configs"
	"github.com/UltronFoundationDev/genesis-generator/generator/genesis"
	"github.com/UltronFoundationDev/genesis-generator/pkeyreader"
	cli "github.com/urfave/cli/v2"
	"log"
	"os"
)

// TODO: need to change dependencies from opera to ultron when public repo will be available

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
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
