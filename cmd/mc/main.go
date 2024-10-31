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
	payload := `{
    "responseHost": "site1.Mastercard.com ",
    "requestId": "demo-iwo-ntu-test-4",
    "paymentAppInstanceId": "123456789",
    "tokens": [
        {
            "correlationId": "ghts-yriti-1244",
            "tokenUniqueReference": "DM4MMC00001441360240f670d8be4a248296cd04379e6a0b",
            "status": "ACTIVE",
            "error": null,
            "eventReasonCode": "REDIGITIZATION_COMPLETE",
            "statusTimestamp": "2023-09-06T04:56:23.345-07:00",
            "suspendedBy": null,
            "tdsRegistrationUrl": "tds.Mastercard.com",
            "tokenRequestorId": "50297129100",
            "productConfig": {
                "brandLogoAssetId": "810f1c0b-3330-42c1-99bf-7ce817b6f458",
                "issuerLogoAssetId": "e2ac9194-a205-482e-9cdb-7ec1bc00f5c6",
                "isCoBranded": "false",
                "coBrandName": null,
                "coBrandLogoAssetId": null,
                "cardBackgroundCombinedAssetId": "4b3caae5-c671-4d81-ab2e-d52dcc8f58f3",
                "cardBackgroundAssetId": "184139f3-780a-4239-b3f7-b779cb52351d",
                "iconAssetId": "b5f1a458-f92f-475b-8540-d75e419737cf",
                "foregroundColor": "c00f00",
                "issuerName": "ProfName Brazil",
                "shortDescription": "Brazil Combo",
                "longDescription": null,
                "customerServiceUrl": null,
                "customerServiceEmail": null,
                "customerServicePhoneNumber": "555-555-4101",
                "issuerMobileApp": null,
                "onlineBankingLoginUrl": null,
                "termsAndConditionsUrl": null,
                "privacyPolicyUrl": null,
                "issuerProductConfigCode": "ProfName"
            },
            "tokenInfo": {
                "tokenPanSuffix": "9436",
                "accountPanSuffix": "9145",
                "accountPanPrefix": "529805",
                "financialAccountSuffix": null,
                "alternateAccountIdentifierSuffix": null,
                "tokenExpiry": "1026",
                "accountPanExpiry": "0335",
                "dsrpCapable": true,
                "tokenAssuranceLevel": 10,
                "productCategory": "CREDIT"
            }
        }
    ]
}`
	encryptPayload := domain.Encrypt(payload, jsonPath, bothEncryptionDecryptionConfigBuilderConfig)
	apputils.PrettyPrintJSON(encryptPayload, "Encrypted Payload:")
	// encryption code logic ends

	// decryption code logic starts
	decryptPayload := domain.Decrypt(encryptPayload, bothEncryptionDecryptionConfigBuilderConfig)
	apputils.PrettyPrintJSON(decryptPayload, "DecryptPayload Payload:")
	// decryption code logic ends
}