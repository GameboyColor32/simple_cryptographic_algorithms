package algorithms

import (
	"encoding/hex"
	"fmt"
	"github.com/GameboyColor32/simple_cryptographic_algorithms/interfaces"
)

type XorCipher struct {
	key string
}

func CreateRandomXorCipher() interfaces.Algorithm {
	fmt.Println("XOR cipher")
	return XorCipher{key: CreateRandomKey()}
}

func CreateXorCipher(key string) interfaces.Algorithm {
	fmt.Println("XOR cipher")
	return XorCipher{key: key}
}

func (self XorCipher) Encrypt(text string) string {
	encr := []byte(text)
	key := []byte(self.key)
	keyLen := len(key)

	for i := range encr {
		encr[i] ^= key[i%keyLen]
	}
	return hex.EncodeToString(encr)
}

func (self XorCipher) Decrypt(text string) string {
	decoded, _ := hex.DecodeString(text)
	encr := []byte(decoded)
	key := []byte(self.key)
	keyLen := len(key)

	for i := range encr {
		encr[i] ^= key[i%keyLen]
	}
	return string(encr)

	//return self.Encrypt(text)
}
