package main

import (
	"fmt"
	"math/rand"

	"golang.org/x/text/unicode/norm"
)

func main() {
	word := "cão" // "dog" in portuguese

	nfc := norm.NFC.String(word)
	nfd := norm.NFD.String(word)

	fmt.Printf("NFC/NFD: %s/%s\n", nfc, nfd)
	fmt.Printf("This is a random number: %f", rand.Float64())
}
