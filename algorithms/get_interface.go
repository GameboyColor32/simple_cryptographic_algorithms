package algorithms

import (
	"github.com/GameboyColor32/simple_cryptographic_algorithms/interfaces"
	"math/rand/v2"
)

func CreateRandomKey() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	length := rand.IntN(13)
	key := make([]rune, length)
	for i := range key {
		key[i] = letters[rand.IntN(len(letters))]
	}
	return string(key)
}

func GenerateRandomCipher() interfaces.Algorithm {
	constructors := map[string]func() interfaces.Algorithm{
		"caesar":   CreateRandomCaesarCipher,
		"vigenere": CreateRandomVigenereCipher,
		"xor":      CreateRandomXorCipher,
	}

	keys := make([]string, 0, len(constructors))
	for k := range constructors {
		keys = append(keys, k)
	}

	selectedCipher := keys[rand.IntN(len(keys))]

	return constructors[selectedCipher]()
}
