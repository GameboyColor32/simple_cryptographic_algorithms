# Simple Cryptographic Algorithms

The goal of my project is to learn the Go language. Interfaces, goroutines, and some networking.

Project objectives:
- [x] Caesar's cipher
- [x] Vegenere's cipher
- [ ] Atbash cipher
- [ ] ROT13
- [ ] XOR cipher
- [ ] AES encryption (optionnal as I don't know how it will be possible to decrypt it on my future project)

Each encryption will be implementing an interface, namely Encrypt and Decrypt. The server will, on request encrypt a secret message from a random encryption algorithm. The encrypted message will then be send back to the client. The goal here is to connect this server to a future project, in which the client will be presented an encrypted message, and will have to decrypt it.
