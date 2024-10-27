package main

import (
	"mastercard_payload_encryption_client_go/config"
	"mastercard_payload_encryption_client_go/domain"
	apputils "mastercard_payload_encryption_client_go/utils"
)

func main() {

	// load configuration starts
	bothEncryptionDecryptionConfigBuilderConfig, err := config.LoadEncryptionDecryptionConfig()
	if err != nil {
		return
	}
	// load configuration ends

	// encryption code logic starts
	jsonPath := "$"
	payload := `{"tokeStatus":"ACTIVE"}`
	encryptPayload := domain.Encrypt(payload, jsonPath, bothEncryptionDecryptionConfigBuilderConfig)
	apputils.PrettyPrintJSON(encryptPayload, "Encrypted Payload:")
	// encryption code logic ends

	// decryption code logic starts
	decryptPayload := domain.Decrypt(encryptPayload, bothEncryptionDecryptionConfigBuilderConfig)
	apputils.PrettyPrintJSON(decryptPayload, "DecryptPayload Payload:")
	// decryption code logic ends
}
