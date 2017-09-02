package utility

import (
	"encoding/hex"
	"fmt"

	rncryptor "github.com/RNCryptor/RNCryptor-go"
)

const (
	encrykey = "$freelancelancerziyoupingsixguys"
)

func Encrypt(message []byte) string {

	return EncryptByKey(message, encrykey)
}

func EncryptByKey(message []byte, salt string) string {
	encrypted, _ := rncryptor.Encrypt(salt, message)
	return hex.EncodeToString(encrypted)
}

func Decrypt(message []byte) string {
	return DecryptByKey(message, encrykey)
}

func DecryptByKey(message []byte, salt string) string {
	hexDecodeMsg, _ := hex.DecodeString(string(message))
	decrypted, err := rncryptor.Decrypt(salt, hexDecodeMsg)
	if err != nil {
		fmt.Printf(err.Error())
	}
	return string(decrypted)
}
