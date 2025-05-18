package aws

import (
	"context"
	"encoding/json"
	"log"
    "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
)

type Resource struct {
	AccountId    string `json:"accountId"`
	AwsRegion    string `json:"awsRegion"`
	ResourceName string `json:"resourceName"`
	Arn          string `json:"arn"`
}

func GettingResourceFromAwsConfigMapInventory(finderQuery string)[][]string{
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config: %v", err)
	}

	client := configservice.NewFromConfig(cfg)
	response, err := client.SelectResourceConfig(context.TODO(), &configservice.SelectResourceConfigInput{
		Expression: aws.String(finderQuery),
	})
	if err != nil {
		log.Fatalf("unable to query AWS Config Recorder: %v", err)
	}
	var data [][]string
	for _, result := range response.Results {
		var resource Resource
		err := json.Unmarshal([]byte(result), &resource)
		if err != nil {
			log.Printf("error unmarshaling resource: %v", err)
			continue
		}
		// Append new entry to the slice
		data = append(data, []string{resource.AccountId, resource.AwsRegion, resource.ResourceName, resource.Arn})
	}
	return data
}
