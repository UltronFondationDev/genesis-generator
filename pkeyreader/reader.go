package pkeyreader

import (
	"bufio"
	"bytes"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"io"
	"os"
)

type PKeyReader struct {
	Path string
}

func New(path string) *PKeyReader {
	return &PKeyReader{
		Path: path,
	}
}

func (r *PKeyReader) GetPKS() ([]*ecdsa.PrivateKey, error) {
	l, err := readLines(r.Path)
	if err != nil {
		return nil, fmt.Errorf("failed to read lines: %v", err)
	}

	var pks []*ecdsa.PrivateKey

	for i, line := range l {
		pkey, err := crypto.HexToECDSA(line)
		if err != nil {
			return nil, fmt.Errorf("failed to make pkey from line %d, err: %v", i+1, err)
		}

		pks = append(pks, pkey)
	}

	return pks, nil
}

// TODO: little refactoring needed
func readLines(path string) (lines []string, err error) {
	var (
		file   *os.File
		part   []byte
		prefix bool
	)

	if file, err = os.Open(path); err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 0))
	for {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(part)
		if !prefix {
			lines = append(lines, buffer.String())
			buffer.Reset()
		}
	}
	if err == io.EOF {
		err = nil
	}
	return
}
