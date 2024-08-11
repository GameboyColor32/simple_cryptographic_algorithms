package algorithms

import (
	"github.com/GameboyColor32/simple_cryptographic_algorithms/interfaces"
	"math/rand"
	"reflect"
	"time"
)

func GenerateRandomCipher() interfaces.Algorithm {
	constructors := map[string]interface{}{
		"caesar":   CreateCaesarCipher,
		"vigenere": CreateVigenereCipher,
	}

	keys := make([]string, 0, len(constructors))
	for k := range constructors {
		keys = append(keys, k)
	}

	rand.Seed(time.Now().UnixNano())

	selectedCipher := keys[rand.Intn(len(keys))]
	constructor := reflect.ValueOf(constructors[selectedCipher])

	var params []reflect.Value
	if selectedCipher == "caesar" {
		params = append(params, reflect.ValueOf(rand.Intn(26)+1))
	} else if selectedCipher == "vigenere" {
		params = append(params, reflect.ValueOf("LEMON"))
	}

	result := constructor.Call(params)

	return result[0].Interface().(interfaces.Algorithm)
}
