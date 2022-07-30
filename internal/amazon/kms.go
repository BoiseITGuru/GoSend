package amazon

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
)

type KMSClient struct {
	*kms.Client
}

func (awsSvc *AwsSvc) KMS() *KMSClient {
	return &KMSClient{
		kms.NewFromConfig(awsSvc.Cfg),
	}
}

func (svc *KMSClient) EncryptString(data string, keyId string) []byte {
	// Encrypt the data
	result, err := svc.Encrypt(context.TODO(), &kms.EncryptInput{
		KeyId:     aws.String(keyId),
		Plaintext: []byte(data),
	})

	if err != nil {
		fmt.Println("Got error encrypting data: ", err)
		os.Exit(1)
	}

	return result.CiphertextBlob
}

func (svc *KMSClient) DecryptString(data []byte) string {
	// Decrypt the data
	result, err := svc.Decrypt(context.TODO(), &kms.DecryptInput{CiphertextBlob: data})

	if err != nil {
		fmt.Println("Got error decrypting data: ", err)
	}

	return string(result.Plaintext)
}
