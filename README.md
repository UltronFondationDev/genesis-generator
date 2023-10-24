# Build and run
1. run command `make` to build application
2. run command `./build/genesis-generator --config=./path_to_conf --pks.path=./path_to_file_with_privatekeys --out.filename="genesis.g" (optional)`

# PKS File
pks file is file with private keys for yours validators, private keys should be follow line by line, see example in ./pks.example.txt

# Configuration
configuration allows to configurate network rules and some genesis data. There is couple ready presets to for devnet and testnet in conf.*.toml files, you can use your own
