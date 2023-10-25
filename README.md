# Build and run
## Build
run command `make` to build application

## Run
Application has following commands:

1. command `./build/genesis-generator gengenesis --config=./path_to_conf --pks.path=./path_to_file_with_privatekeys --out.filename="genesis.g" (optional)` the output will be genesis file with provided name or default genesis.g
2. command `./build/genesis-generator genkeys --count=2` the output will folder with provided count of private, pub keys and addresses

# PKS File
pks file is file with private keys for yours validators, private keys should be follow line by line, see example in ./pks.example.txt

# Configuration
configuration allows to configurate network rules and some genesis data. There is couple ready presets to for devnet and testnet in conf.*.toml files, you can use your own
