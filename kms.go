package lib

import (
	"encoding/base64"

	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/kms"
)

// Decrypt decrypt kms token
func Decrypt(token string) string {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("KMS: Unable to load SDK config: " + err.Error())
	}
	cfg.Region = endpoints.UsEast1RegionID

	decoded, _ := base64.StdEncoding.DecodeString(token)

	svc := kms.New(cfg)
	input := &kms.DecryptInput{
		CiphertextBlob: []byte(decoded),
	}

	req := svc.DecryptRequest(input)
	res, err := req.Send()

	if err != nil {
		panic(err)
	}

	return string(res.Plaintext)
}
