package b64

import (
	"crypto/aes"
	"crypto/cipher"
	//	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	//	"io"
)

func EncryptText(text string) (string, error) {
	// key := []byte(setting.GetSystemSetting("thaiwater30.util.b64.key"))
	key := []byte("d73befaf0b46d870")
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	//	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
	//		return "", err
	//	}
	for i, _ := range iv {
		iv[i] = byte(i)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}
func DecryptText(cryptoText string) (string, error) {
	// key := []byte(setting.GetSystemSetting("thaiwater30.util.b64.key"))
	key := []byte("d73befaf0b46d870")
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext), nil
}
