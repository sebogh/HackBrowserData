//go:build linux

package crypto

import "errors"

func DecryptWithChromium(key, encryptPass []byte) ([]byte, error) {
	if len(encryptPass) < 3 {
		return nil, ErrCiphertextLengthIsInvalid
	}
	iv := []byte{32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32}
	return AES128CBCDecrypt(key, iv, encryptPass[3:])
}

func DecryptWithDPAPI(_ []byte) ([]byte, error) {
	return nil, errors.New("not windows")
}
