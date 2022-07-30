package amazon

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

type AwsSvc struct {
	Cfg  aws.Config
	Keys KmsKeys
}

type KmsKeys struct {
	EncryptionKey string
	JwtKey        string
}

var Service *AwsSvc

func BuildAWSClient(awsRegion string, kmsKeys KmsKeys) *AwsSvc {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(awsRegion),
	)
	if err != nil {
		// TODO: handle error
		log.Fatalln("Error Configurating AWS Client")
	}

	return &AwsSvc{
		Cfg:  cfg,
		Keys: kmsKeys,
	}
}

func (awsSvc *AwsSvc) Init() {
	Service = awsSvc
}
