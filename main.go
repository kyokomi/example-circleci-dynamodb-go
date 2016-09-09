package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/labstack/gommon/log"
)

const (
	region    = "us-west-2"
	endpoint  = "http://127.0.0.1:7777"
	tableName = "example-table"
)

var (
	dyn *dynamo.DB
)

type exampleDynamoTable struct {
	ID      int `dynamo:"ID,hash"`
	Count   int
	Message string
}

func main() {
	dyn = dynamo.New(session.New(), &aws.Config{
		Region:   aws.String(region),
		Endpoint: aws.String(endpoint),
	})

	if err := dyn.CreateTable(tableName, exampleDynamoTable{}).Run(); err != nil {
		fmt.Println()
	}

	id := 1
	if err := updateExampleData(id, 100, "test"); err != nil {
		log.Fatal(err)
	}

	exampleData, err := fetchExampleData(id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", exampleData)
}

func updateExampleData(id int, count int, message string) error {
	updater := dyn.Table(tableName).Update("ID", id)
	updater.Set("Count", count)
	updater.Set("Message", message)
	return updater.Run()
}

func fetchExampleData(id int) (exampleDynamoTable, error) {
	var e exampleDynamoTable
	return e, dyn.Table(tableName).Get("ID", id).One(&e)
}
