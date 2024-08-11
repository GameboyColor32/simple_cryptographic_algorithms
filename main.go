package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/GameboyColor32/simple_cryptographic_algorithms/algorithms"
	"github.com/GameboyColor32/simple_cryptographic_algorithms/networking"
)

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	algo := algorithms.CreateXorCipher("test")
	http.HandleFunc("/", networking.GenerateCipher) //algo := algorithms.CreateCaesarCipher(3)
	fmt.Println(algo.Encrypt("Hello"), algo.Decrypt(algo.Encrypt("Hello")))
	//args := os.Args

	err := http.ListenAndServe(":3333", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
	//l := algorithms.CreateVigenereCipher("LEMON")
	//fmt.Println(l.Encrypt("Hello"))
	//for _, arg := range args[1:] {
	//	fmt.Println(algo.Encrypt(arg), "decrypted: ", algo.Decrypt(algo.Encrypt(arg)))
	//}
}
