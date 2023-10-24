package configs

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/Fantom-foundation/go-opera/inter"
	"github.com/Fantom-foundation/go-opera/opera"
	"github.com/naoina/toml"
	"math/big"
	"os"
	"reflect"
)

var tomlSettings = toml.Config{
	NormFieldName: func(rt reflect.Type, key string) string {
		return key
	},
	FieldToKey: func(rt reflect.Type, field string) string {
		return field
	},
	MissingField: func(rt reflect.Type, field string) error {
		return fmt.Errorf("field '%s' is not defined in %s", field, rt.String())
	},
}

type GenesisConfig struct {
	EVMAccountAddress string
	FirstEpoch        int

	GenesisTime inter.Timestamp
	BlockTime   inter.Timestamp

	Rules opera.Rules

	ValidatorBalance   *big.Int
	ValidatorStakedAmt *big.Int
}

func LoadConfig(path string) (*GenesisConfig, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg GenesisConfig

	err = tomlSettings.NewDecoder(bufio.NewReader(f)).Decode(&cfg)
	// Add file name to errors that have a line number.
	if err != nil {
		if _, ok := err.(*toml.LineError); ok {
			return nil, errors.New(path + ", " + err.Error())
		} else {
			return nil, errors.New(fmt.Sprintf("TOML config file error: %v.\n", err))
		}
	}

	return &cfg, err
}
