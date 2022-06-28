# To generate a genesis file `genesis.g`

Go to `cmd/mainnet`, `cmd/testnet` or `cmd/devnet`

Copy `pks.txt.example` to `pks.txt`, add your validators' private keys there (line-by-line) and run generator.

```bash
cp pks.txt.example pks.txt
# edit pks.txt
go run generate.go
```

The `genesis.g` file is created, use it to start your network.