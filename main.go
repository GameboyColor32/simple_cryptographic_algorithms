package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	//"github.com/GameboyColor32/simple_cryptographic_algorithms/algorithms"
	"github.com/GameboyColor32/simple_cryptographic_algorithms/networking"
)

func main() {
	//algo := algorithms.CreateVigenereCipher("LEMON")
	http.HandleFunc("/", networking.GenerateCipher) //algo := algorithms.CreateCaesarCipher(3)
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
