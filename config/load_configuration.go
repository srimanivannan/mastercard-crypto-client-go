package config

import (
	"fmt"
	"github.com/mastercard/client-encryption-go/field_level_encryption"
	"github.com/mastercard/client-encryption-go/jwe"
	"github.com/mastercard/client-encryption-go/utils"
)

const (
	EncryptedPayloadJsonPathFieldName     = "encryptedPayload"
	EncryptedDataJsonPathFieldName        = "encryptedData"
	EncryptedKeyJsonPathFieldName         = "encryptedKey"
	PublicKeyFingerprintJsonPathFieldName = "publicKeyFingerprint"
	OaepHashingAlgorithmJsonPathFieldName = "oaepHashingAlgorithm"
	JsonWildCardSymbol                    = "$"
	IvJsonPathFieldName                   = "iv"
	PublicKeyFingerprintDevPortalValue    = "<<fingerPrint>>"
)

func LoadEncryptionDecryptionConfig() (*field_level_encryption.FieldLevelEncryptionConfig, error) {
	// Load decryption config file starts
	decryptionKey, decryptionP12Error := utils.LoadDecryptionKey("credentials/cert.p12", "<<password>>")

	if decryptionP12Error != nil {
		fmt.Println("Error loading p12 file:", decryptionP12Error)
		return nil, decryptionP12Error
	}
	// Load decryption config file ends

	// Load encryption config file starts
	encryptionCertificate, certificateErr := utils.LoadEncryptionCertificate("credentials/cert.pem")
	if certificateErr != nil {
		fmt.Println("Error loading pem file:", certificateErr)
		return nil, certificateErr
	}
	// Load encryption config file ends

	// BothEncryptionDecryptionConfigBuilder starts
	bothEncryptionDecryptionConfigBuilder := field_level_encryption.NewFieldLevelEncryptionConfigBuilder()
	bothEncryptionDecryptionConfigBuilderConfig, bothEncryptionDecryptionConfigError := bothEncryptionDecryptionConfigBuilder.
		WithEncryptionCertificate(encryptionCertificate).
		WithEncryptionPath(JsonWildCardSymbol, EncryptedPayloadJsonPathFieldName).
		WithEncryptedValueFieldName(EncryptedDataJsonPathFieldName).
		WithEncryptedKeyFieldName(EncryptedKeyJsonPathFieldName).
		WithIvFieldName(IvJsonPathFieldName).
		WithFieldValueEncoding(field_level_encryption.HEX).
		WithOaepPaddingDigestAlgorithm(field_level_encryption.SHA256).
		WithOaepPaddingDigestAlgorithmFieldName(OaepHashingAlgorithmJsonPathFieldName).
		WithEncryptionCertificateFingerprintFieldName(PublicKeyFingerprintJsonPathFieldName).
		WithEncryptionCertificateFingerprint(PublicKeyFingerprintDevPortalValue).
		WithDecryptionKey(decryptionKey).                                          //decrypt
		WithDecryptionPath(EncryptedPayloadJsonPathFieldName, JsonWildCardSymbol). //decrypt
		Build()

	if bothEncryptionDecryptionConfigError != nil {
		fmt.Println("Error when building bothEncryptionDecryptionConfigBuilderConfig:", bothEncryptionDecryptionConfigError)
		return nil, bothEncryptionDecryptionConfigError
	}
	return bothEncryptionDecryptionConfigBuilderConfig, nil
	// BothEncryptionDecryptionConfigBuilder ends
}

func LoadJWEEncryptionDecryptionConfig(iv []byte, cek []byte) (*jwe.JWEConfig, error) {
	// Load decryption config file starts
	decryptionKey, decryptionP12Error := utils.LoadDecryptionKey("credentials/cert.p12", "<<password>>")

	if decryptionP12Error != nil {
		fmt.Println("Error loading p12 file:", decryptionP12Error)
		return nil, decryptionP12Error
	}
	// Load decryption config file ends

	// Load encryption config file starts
	encryptionCertificate, certificateErr := utils.LoadEncryptionCertificate("credentials/cert.pem")
	if certificateErr != nil {
		fmt.Println("Error loading pem file:", certificateErr)
		return nil, certificateErr
	}
	// Load encryption config file ends
	jweConfigBuilder := jwe.NewJWEConfigBuilder()
	jweConfig := jweConfigBuilder.
		WithCertificate(encryptionCertificate).
		WithEncryptionPath(JsonWildCardSymbol, JsonWildCardSymbol).
		WithEncryptedValueFieldName(EncryptedDataJsonPathFieldName).
		WithIv(iv).
		WithCek(cek).
		WithDecryptionKey(decryptionKey).                           //decrypt
		WithDecryptionPath(JsonWildCardSymbol, JsonWildCardSymbol). //decrypt
		Build()
	return jweConfig, nil
}
