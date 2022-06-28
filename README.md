# To generate a genesis file `genesis.g`

Go to `cmd/mainnet` or `cmd/testnet`

Copy `pks.txt.example` to `pks.txt`. Store private keys there (line-by-line) and run generator.
```
cp pks.txt.example pks.txt
go run generate.go
```

`genesis.g` file is created, use it to start the network.