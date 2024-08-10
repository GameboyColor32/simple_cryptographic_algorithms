package algorithms

import (
	"fmt"
	"strings"
)

type VigenereCipher struct {
	key string
}

func CreateVigenereCipher(key string) *VigenereCipher {
	return &VigenereCipher{key: key}
}

func (self VigenereCipher) modifyKey(text string) string {
	fmt.Println("Vigenere", text, self.key)
	var sb strings.Builder

	for idx := range text {
		fmt.Println("Vigenere", text, self.key, idx%len(self.key))
		sb.WriteByte(self.key[idx%len(self.key)])
	}
	return strings.ToLower(sb.String())
}

func (self VigenereCipher) Encrypt(text string) string {
	fmt.Println("Vigenere here")

	newKey := self.modifyKey(text)
	var message strings.Builder
	alph := "abcdefghijklmnopqrstuvwxyz"

	fmt.Println("Vigenere")
	for i := range newKey {
		if text[i] < 'a' || text[i] > 'z' {
			message.WriteRune(rune(text[i]))
			continue
		}
		message.WriteRune(rune(alph[(text[i]-'a'+newKey[i]-'a')%26]))
		fmt.Println("Vigenere")
	}
	return message.String()
}

func (self VigenereCipher) Decrypt(text string) string {
	newKey := self.modifyKey(text)
	var message strings.Builder
	alph := "abcdefghijklmnopqrstuvwxyz"

	for i := range newKey {
		if text[i] < 'a' || text[i] > 'z' {
			message.WriteRune(rune(text[i]))
			continue
		}
		message.WriteRune(rune(alph[(text[i]-'a'-(newKey[i]-'a')+26)%26]))
	}
	return message.String()
}
