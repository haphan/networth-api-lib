package lib

import (
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// DocumentClient new instance of doc client
func DocumentClient() *dynamodb.DynamoDB {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("DynamoDB: Unable to load SDK config: " + err.Error())
	}

	cfg.Region = endpoints.UsEast1RegionID
	table := dynamodb.New(cfg)

	return table
}
