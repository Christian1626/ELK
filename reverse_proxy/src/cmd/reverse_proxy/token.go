package main

//
//import (
//	"crypto/aes"
//	"crypto/cipher"
//	"encoding/base64"
//	"log"
//	"strings"
//)
//
//func decrypt_token() {
//	//replace spaces with '+'
//	token = strings.Replace(token, " ", "+", -1)
//
//	//get decrypted token
//	signature = string(cbcDecrypt(token))
//	//replace null characters
//	signature = strings.Replace(signature, "\x00", "", -1)
//	log.Println("Decrypted String :", signature)
//}
//
//func cbcDecrypt(text1 string) []byte {
//	key := []byte(config.DecryptKey)
//	ciphertext, _ := base64.StdEncoding.DecodeString(string(text1))
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		panic(err)
//	}
//
//	// include it at the beginning of the ciphertext.
//	if len(ciphertext) < aes.BlockSize {
//		panic("ciphertext too short")
//	}
//	iv := ciphertext[:aes.BlockSize]
//	ciphertext = ciphertext[aes.BlockSize:]
//
//	// CBC mode always works in whole blocks.
//	if len(ciphertext)%aes.BlockSize != 0 {
//		panic("ciphertext is not a multiple of the block size")
//	}
//
//	mode := cipher.NewCBCDecrypter(block, iv)
//
//	// CryptBlocks can work in-place if the two arguments are the same.
//	mode.CryptBlocks(ciphertext, ciphertext)
//	ciphertext = PKCS5UnPadding(ciphertext)
//	return ciphertext
//}
//
//func PKCS5UnPadding(src []byte) []byte {
//	length := len(src)
//	//unpadding := int(src[length-1])
//	//log.Println("8: len:", src, "  unpadding:", unpadding)
//	return src[:(length)]
//}
