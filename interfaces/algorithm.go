package interfaces

type Algorithm interface {
	Encrypt(text string) string
	Decrypt(text string) string
}
