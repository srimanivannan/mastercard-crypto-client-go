package utils

import (
	"crypto/aes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/TylerBrock/colorjson"
	"io"
)

func PrettyPrintJSON(jsonInput string, message string) string {
	var obj map[string]interface{}
	err := json.Unmarshal([]byte(jsonInput), &obj)
	if err != nil {
		return ""
	}

	formatter := colorjson.NewFormatter()
	formatter.Indent = 2

	coloredResult, _ := formatter.Marshal(obj)
	fmt.Println(message)
	fmt.Println(string(coloredResult))
	return string(coloredResult)
}

// GenerateCEK generates a random CEK (Content Encryption Key) of a specified length.
// The length is typically 16, 24, or 32 bytes for AES-128, AES-192, or AES-256 respectively.
func GenerateCEK(keySize int) ([]byte, error) {
	cek := make([]byte, keySize)
	_, err := io.ReadFull(rand.Reader, cek)
	if err != nil {
		return nil, fmt.Errorf("failed to generate CEK: %w", err)
	}
	fmt.Printf("Generated CEK: %x\n", cek)
	return cek, nil
}

// GenerateIV generates a random Initialization Vector (IV) for AES encryption
// The IV size should match the AES block size, which is 16 bytes.
func GenerateIV() ([]byte, error) {
	iv := make([]byte, aes.BlockSize) // AES block size is 16 bytes
	_, err := io.ReadFull(rand.Reader, iv)
	if err != nil {
		return nil, fmt.Errorf("failed to generate IV: %w", err)
	}
	fmt.Printf("Generated IV: %x\n", iv)
	return iv, nil
}
