package pkeys

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type PKey struct {
	Public  string
	Private string
	Address string
}

func GenerateKeyAndAddress() (*PKey, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, err
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return &PKey{
		Public:  hexutil.Encode(publicKeyBytes)[4:],
		Private: hexutil.Encode(privateKeyBytes)[2:],
		Address: address,
	}, nil
}
