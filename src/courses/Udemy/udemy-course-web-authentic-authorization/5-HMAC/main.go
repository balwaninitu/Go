package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
)

var key = []byte{}

//HMAC is cryptographic signing algorithm
func main() {
	for i := 1; i <= 64; i++ {
		key = append(key, byte(i))
	}

}

func signMessage(msg []byte) ([]byte, error) {
	//below func takes fist param hash.Hash made by func for that can use sha512 func
	//second param is key which is private key you generate yourself and use with all hmac func
	//key is used for cryptographic signature and for validation
	//hmac is good for messaging yourself bcause key is with you to validate
	//key size should match with size of algorithm
	//sha512 is 64 bytes(checked from constants)
	//h is hash and it is io writer
	h := hmac.New(sha512.New, key) //sha 512 is 64byte*8(each byte) = 512

	_, err := h.Write(msg)
	if err != nil {
		return nil, fmt.Errorf("Error in signMessage while hashing: %w", err)
	}
	//
	signature := h.Sum(nil)
	return signature, nil

}

//func return you signature and msg together which can use as bearer token
//user send you bearer token which need to compare
func checkSig(msg, sig []byte) (bool, error) {
	newSig, err := signMessage(msg)

	if err != nil {
		return false, fmt.Errorf("Error in checksig while getting signature message: %w", err)

	}

	same := hmac.Equal(newSig, sig)
	return same, nil
}
