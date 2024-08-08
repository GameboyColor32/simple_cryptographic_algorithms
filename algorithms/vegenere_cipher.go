package algorithms

import (
	"strings"
)

type VigenereCipher struct {
	key string
}

func CreateVigenereCipher(key string) *VigenereCipher {
	return &VigenereCipher{key: key}
}

func (self VigenereCipher) modifyKey(text string) string {
	var sb strings.Builder

	for idx := range text {
		sb.WriteByte(self.key[idx%len(self.key)])
	}
	return strings.ToLower(sb.String())
}

func (self VigenereCipher) Encrypt(text string) string {
	newKey := self.modifyKey(text)
	var message strings.Builder
	alph := "abcdefghijklmnopqrstuvwxyz"

	for i := range newKey {
		if text[i] < 'a' || text[i] > 'z' {
			message.WriteRune(rune(text[i]))
			continue
		}
		message.WriteRune(rune(alph[(text[i]-'a'+newKey[i]-'a')%26]))
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
