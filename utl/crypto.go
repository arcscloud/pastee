package utl

import (
    "bytes"
    "crypto/aes"
    "crypto/cipher"
    "encoding/base64"
)

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
    padding := blockSize - len(ciphertext)%blockSize
    padText := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(ciphertext, padText...)
}

func PKCS7UnPadding(origData []byte) []byte {
    length := len(origData)
    unpadding := int(origData[length-1])
    return origData[:(length - unpadding)]
}

func AesEncryptCBC(orig string, key string) (string, error) {
    origData := []byte(orig)
    k := []byte(key)

    block, err := aes.NewCipher(k)
    if err != nil {
        return "", err
    }

    blockSize := block.BlockSize()

    origData = PKCS7Padding(origData, blockSize)

    blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])

    cryted := make([]byte, len(origData))

    blockMode.CryptBlocks(cryted, origData)
    return base64.StdEncoding.EncodeToString(cryted), nil
}

func AesDecryptCBC(cryted string, key string) (string, error) {
    crytedByte, err := base64.StdEncoding.DecodeString(cryted)
    if err != nil {
        return "", err
    }
    k := []byte(key)

    block, err := aes.NewCipher(k)
    if err != nil {
        return "", err
    }

    blockSize := block.BlockSize()

    blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])

    orig := make([]byte, len(crytedByte))

    blockMode.CryptBlocks(orig, crytedByte)

    orig = PKCS7UnPadding(orig)
    return string(orig), nil
}
