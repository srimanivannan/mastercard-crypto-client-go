package domain

import (
	"github.com/Jeffail/gabs/v2"
	"github.com/mastercard/client-encryption-go/encryption"
	"github.com/mastercard/client-encryption-go/field_level_encryption"
	"github.com/mastercard/client-encryption-go/jwe"
	"github.com/mastercard/client-encryption-go/mastercard_encryption"
	"github.com/mastercard/client-encryption-go/utils"
)

func Encrypt(payload string, jsonPath string, bothEncryptionDecryptionConfigBuilderConfig *field_level_encryption.FieldLevelEncryptionConfig) string {
	jsonPayload, _ := gabs.ParseJSON([]byte(payload))
	payloadToEncrypt := utils.GetPayloadToEncrypt(jsonPayload, jsonPath)
	encryptPayload := mastercard_encryption.EncryptPayload(payloadToEncrypt, *bothEncryptionDecryptionConfigBuilderConfig)
	//fmt.Println("Encrypted Payload:", encryptPayload)
	return encryptPayload
}

func JWEEncrypt(payload string, jsonPath string, jweEncryptionDecryptionConfigBuilderConfig *jwe.JWEConfig) string {
	jsonPayload, _ := gabs.ParseJSON([]byte(payload))
	payloadToEncrypt := utils.GetPayloadToEncrypt(jsonPayload, jsonPath)
	encryptPayload := encryption.EncryptPayload(payloadToEncrypt, *jweEncryptionDecryptionConfigBuilderConfig)
	//fmt.Println("Encrypted Payload:", encryptPayload)
	return encryptPayload
}
