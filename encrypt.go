package fiputil

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"fmt"
)

const (
	keyLength = 16
	codeSize  = 52
)

var (
	defaultKey = []byte{54, 25, 62, 43, 59, 13, 6, 42, 9, 53, 37, 32, 76, 25, 5, 45}
)

func Encrypt(text, key []byte) (code []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	hash := sha1.Sum(text)

	//fmt.Println("hash:",hash,",len:",len(hash),sha1.Size)
	cfb := cipher.NewCFBEncrypter(block, key)
	code = make([]byte, sha1.Size)
	cfb.XORKeyStream(code, hash[:])

	//fmt.Printf("key:%d,cipher: %d\n",len(key),len(code))
	//pw = make([]byte, 52)
	//base64.StdEncoding.Encode(pw[:28],code)
	//base64.StdEncoding.Encode(pw[28:],key)

	return
}

func HashPwd(pwd []byte) (hashed []byte, err error) {
	key := genKey()
	code, err := Encrypt(pwd, key)
	if err != nil {
		fmt.Println("hash password error:", err)
		return
	}
	hashed = append(code, key...)
	return
}

func CheckPwd(pwdInput, pwdStore []byte) (ok bool) {
	if len(pwdStore) < 36 {
		return
	}
	key := pwdStore[20:]

	code, err := Encrypt(pwdInput, key)
	if err != nil {
		fmt.Println("Encrypt failed:", err)
		return
	}

	if bytes.Equal(pwdStore[:20], code) {
		return true
	}
	return
}

func genKey() (key []byte) {
	key = make([]byte, keyLength)
	_, err := rand.Read(key)
	if err != nil {
		fmt.Println("gen key error:", err)
		return defaultKey
	}
	return
}
