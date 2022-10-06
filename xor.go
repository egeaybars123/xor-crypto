package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	fmt.Println("Welcome to XOR-crypto!")
	key := generateRandomKey()
	fmt.Println("Your Key: ", key)

	//keyfi bir input girebilirsiniz.
	cipher := encryptXOR(93, key) //şifreli metin
	fmt.Println("Şifreli Metin: ", cipher)

	decrypted := decryptXOR(cipher, key)
	fmt.Println("Metin: ", decrypted)
}

func generateRandomKey() int {
	num := rand.Intn(256)
	return num
}

func encryptXOR(text int, key int) int {
	result := text ^ key
	return result
}

func decryptXOR(cipher int, key int) int {
	result := cipher ^ key
	return result
}
