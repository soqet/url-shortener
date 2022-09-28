package shortlinkgen

import (
	"crypto/sha256"
	"fmt"
	"github.com/itchyny/base58-go"
	"math/big"
)

func hashString(str string) []byte {
	alg := sha256.New()
	alg.Write([]byte(str))
	return alg.Sum(nil)
}

func encodeBase58(bytes []byte) (string, error) {
	encoded, err := base58.BitcoinEncoding.Encode(bytes)
	if err != nil {
		return "", err
	}
	return string(encoded), nil
}

func GenerateShortLink(url string) (string, error) {
	hash := hashString(url)
	number := new(big.Int).SetBytes(hash).Uint64()
	short, err := encodeBase58([]byte(fmt.Sprintf("%d", number)))
	if err != nil {
		return "", err
	}
	return short[:10], nil
}
