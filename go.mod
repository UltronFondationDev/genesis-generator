module github.com/UltronFoundationDev/genesis-generator

go 1.14

require (
	github.com/Fantom-foundation/go-opera v1.0.2-rc.5
	github.com/Fantom-foundation/lachesis-base v0.0.0-20210721130657-54ad3c8a18c1
	github.com/ethereum/go-ethereum v1.9.22
	github.com/naoina/toml v0.1.2-0.20170918210437-9fafd6967416
	github.com/urfave/cli/v2 v2.25.7
)

replace github.com/ethereum/go-ethereum => github.com/Fantom-foundation/go-ethereum v1.9.7-0.20210827160629-07563551b4c0

replace github.com/dvyukov/go-fuzz => github.com/guzenok/go-fuzz v0.0.0-20210103140116-f9104dfb626f
