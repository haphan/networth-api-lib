package lib

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbattribute"
)

var accountTable = aws.String(GetEnv("ACCOUNT_TABLE", ""))

// DBInterface db interface
type DBInterface interface {
	DocumentClient() *dynamodb.DynamoDB
	GetTokens() map[string]interface{}
	GetAccounts() map[string]interface{}
}

// DB db struct
type DB struct {
	table *dynamodb.DynamoDB
}

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

// GetTokens return tokens from db
func (d DB) GetTokens(username string) []string {
	req := d.table.GetItemRequest(&dynamodb.GetItemInput{
		TableName: accountTable,
		Key: map[string]dynamodb.AttributeValue{
			"username": {S: aws.String(fmt.Sprintf("%s:tokens", username))},
		},
	})

	res, err := req.Send()
	if err != nil {
		panic(err.Error())
	}

	tokens := make(map[string]interface{})
	if err := dynamodbattribute.UnmarshalMap(res.Item, &tokens); err != nil {
		panic(err)
	}

	payload := []string{""}
	for k, v := range tokens {
		if strings.HasPrefix(k, "ins_") {
			decrypted := Decrypt(v.([]string)[0])
			payload = append(payload, decrypted)
		}
	}

	return payload
}

// GetAccounts return accounts from db
func (d DB) GetAccounts(table *dynamodb.DynamoDB, username string) map[string]interface{} {
	req := d.table.GetItemRequest(&dynamodb.GetItemInput{
		TableName: accountTable,
		Key: map[string]dynamodb.AttributeValue{
			"username": {S: aws.String(fmt.Sprintf("%s:accounts", username))},
		},
	})

	res, err := req.Send()
	if err != nil {
		panic(err.Error())
	}

	account := make(map[string]interface{})
	if err := dynamodbattribute.UnmarshalMap(res.Item, &account); err != nil {
		panic(err)
	}

	return account
}
