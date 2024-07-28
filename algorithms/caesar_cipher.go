package algorithms

type CaesarCipher struct {
	shift int
}

func CreateCaesarCipher(shift int) *CaesarCipher {
	cc := CaesarCipher{shift: shift}
	return &cc
}

func (self CaesarCipher) Encrypt(text string) string {
	shifted := []rune(text)

	for i, char := range text {
		if char >= 'A' && char <= 'Z' {
			shifted[i] = 'A' + (char-'A'+rune(self.shift))%26
		} else if char >= 'a' && char <= 'z' {
			shifted[i] = 'a' + (char-'a'+rune(self.shift))%26
		} else {
			shifted[i] = char
		}
	}
	return string(shifted)
}

func (self CaesarCipher) Decrypt(text string) string {
	shifted := []rune(text)

	for i, char := range text {
		if char >= 'A' && char <= 'Z' {
			shifted[i] = 'A' + (char-'A'-rune(self.shift)+26)%26
		} else if char >= 'a' && char <= 'z' {
			shifted[i] = 'a' + (char-'a'-rune(self.shift)+26)%26
		} else {
			shifted[i] = char
		}
	}
	return string(shifted)
}
