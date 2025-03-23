package encryption

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/hex"
    "errors"
    "bytes"
)

// PKCS7 padding uygulaması
func applyPadding(data []byte, blockSize int) []byte {
    padding := blockSize - len(data)%blockSize
    padText := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(data, padText...)
}

// Padding'i kaldırma işlemi
func removePadding(data []byte, blockSize int) ([]byte, error) {
    length := len(data)
    padding := int(data[length-1])
    if padding > length || padding > blockSize {
        return nil, errors.New("invalid padding")
    }
    return data[:length-padding], nil
}

// EncryptData şifreleme işlevi
func EncryptData(key []byte, data string) (string, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }

    // Veriyi AES için uygun uzunluğa getirecek padding uygulaması
    dataBytes := []byte(data)
    blockSize := block.BlockSize()
    dataBytes = applyPadding(dataBytes, blockSize)

    ciphertext := make([]byte, len(dataBytes))
    iv := make([]byte, blockSize)
    if _, err := rand.Read(iv); err != nil {
        return "", err
    }

    mode := cipher.NewCBCEncrypter(block, iv)
    mode.CryptBlocks(ciphertext, dataBytes)

    // IV'yi başta ekleyip ardından şifreli veriyi hexadecimal formatında döndürüyoruz
    ivAndCiphertext := append(iv, ciphertext...)
    return hex.EncodeToString(ivAndCiphertext), nil
}

// DecryptData şifre çözme işlevi
func DecryptData(key []byte, encryptedData string) (string, error) {
    data, err := hex.DecodeString(encryptedData)
    if err != nil {
        return "", err
    }

    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }

    blockSize := block.BlockSize()
    if len(data) < blockSize {
        return "", errors.New("ciphertext too short")
    }

    iv, ciphertext := data[:blockSize], data[blockSize:]
    mode := cipher.NewCBCDecrypter(block, iv)
    mode.CryptBlocks(ciphertext, ciphertext)

    // Padding'i kaldırmak
    decrypted, err := removePadding(ciphertext, blockSize)
    if err != nil {
        return "", err
    }

    return string(decrypted), nil
}
