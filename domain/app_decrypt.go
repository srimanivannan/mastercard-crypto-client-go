package domain

import (
	"github.com/mastercard/client-encryption-go/encryption"
	"github.com/mastercard/client-encryption-go/field_level_encryption"
	"github.com/mastercard/client-encryption-go/jwe"
	"github.com/mastercard/client-encryption-go/mastercard_encryption"
)

func Decrypt(encryptPayload string, bothEncryptionDecryptionConfigBuilderConfig *field_level_encryption.FieldLevelEncryptionConfig) string {
	decryptPayload := mastercard_encryption.DecryptPayload(encryptPayload, *bothEncryptionDecryptionConfigBuilderConfig)
	//fmt.Println("DecryptPayload Payload:", decryptPayload)
	return decryptPayload
}

func JWEDecrypt(encryptPayload string, jweEncryptionDecryptionConfigBuilderConfig *jwe.JWEConfig) string {
	payload := encryption.DecryptPayload(encryptPayload, *jweEncryptionDecryptionConfigBuilderConfig)
	//fmt.Println("DecryptPayload Payload:", decryptPayload)
	return payload
}
