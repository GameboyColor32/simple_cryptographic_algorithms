package main

import (
	"fmt"
	"github.com/GameboyColor32/simple_cryptographic_algorithms/algorithms"
	"os"
)

func main() {
	algo := algorithms.CreateVigenereCipher("LEMON")
	//algo := algorithms.CreateCaesarCipher(3)
	args := os.Args

	for _, arg := range args[1:] {
		fmt.Println(algo.Encrypt(arg), "decrypted: ", algo.Decrypt(algo.Encrypt(arg)))
	}
}
