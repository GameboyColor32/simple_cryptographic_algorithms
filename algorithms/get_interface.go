package algorithms

import (
	"github.com/GameboyColor32/simple_cryptographic_algorithms/interfaces"
	"math/rand"
	"reflect"
	"time"
)

func GenerateRandomCipher() interfaces.Algorithm {
	rand.Seed(time.Now().UnixNano())

	types := []reflect.Type{
		reflect.TypeOf(CaesarCipher{shift: rand.Intn(26) + 1}),
		reflect.TypeOf(VigenereCipher{key: "LEMON"}),
	}

	randomIndex := rand.Intn(len(types))
	selectedType := types[randomIndex]
	newInstance := reflect.New(selectedType).Elem().Interface()

	return newInstance.(interfaces.Algorithm)
}
