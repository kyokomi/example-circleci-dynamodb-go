package main

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/stretchr/testify/assert"
)

func TestExampleData(t *testing.T) {
	as := assert.New(t)

	dyn = dynamo.New(session.New(), &aws.Config{
		Region:   aws.String("test-region"),
		Endpoint: aws.String(endpoint),
	})

	if err := dyn.Table(tableName).DeleteTable().Run(); err != nil {
		fmt.Println()
	}
	if err := dyn.CreateTable(tableName, exampleDynamoTable{}).Run(); err != nil {
		fmt.Println()
	}

	testCase := exampleDynamoTable{
		ID:      2,
		Count:   300,
		Message: "hoge",
	}

	as.NoError(updateExampleData(testCase.ID, testCase.Count, testCase.Message))

	exampleData, err := fetchExampleData(testCase.ID)
	as.NoError(err)
	as.EqualValues(exampleData, testCase)
}
