package amazon

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type SecretsClient struct {
	*secretsmanager.Client
}

func (awsSvc *AwsSvc) SecretsManager() *SecretsClient {
	return &SecretsClient{
		secretsmanager.NewFromConfig(awsSvc.Cfg),
	}
}

func (svc *SecretsClient) StoreSecret(name string, data string) string {
	secretData := &secretsmanager.CreateSecretInput{
		Name:         &name,
		SecretString: &data,
	}

	result, err := svc.CreateSecret(context.TODO(), secretData)

	if err != nil {
		// TODO: handle error
		fmt.Println("Got error encrypting data: ", err)
	}

	return *result.ARN
}

func (svc *SecretsClient) GetSecretData(data string) string {
	secretData := &secretsmanager.GetSecretValueInput{
		SecretId: &data,
	}

	result, err := svc.GetSecretValue(context.TODO(), secretData)

	if err != nil {
		// TODO: handle error
		fmt.Println("Got error encrypting data: ", err)
	}

	return *result.SecretString
}
