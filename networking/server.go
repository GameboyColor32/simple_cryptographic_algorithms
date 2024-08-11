package networking

import (
	"github.com/GameboyColor32/simple_cryptographic_algorithms/algorithms"
	"io"
	"net/http"
)

func GetCipher(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello\n")
	//return &algorithms.CreateVigenereCipher("LEMON")
}

func GenerateCipher(w http.ResponseWriter, r *http.Request) {
	cipher := algorithms.GenerateRandomCipher()

	io.WriteString(w, cipher.Encrypt("Helloo")+"\n")
}
