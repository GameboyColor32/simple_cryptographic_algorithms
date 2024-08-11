package algorithms

import (
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
	encr := []rune(text)
	key := []rune(self.key)
	keyLen := len(key)

	for i, char := range encr {
		encr[i] = char ^ key[i%keyLen]
	}
	return string(encr)
}

func (self XorCipher) Decrypt(text string) string {
	return self.Encrypt(text)
}
