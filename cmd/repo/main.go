package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"log"
	"os"
	"pyrch-go/internal/apigwp"
	"pyrch-go/pkg/model"
	"strconv"
	"time"
)

var db *dynamodb.DynamoDB

func init() {
	if sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	}); err != nil {
		log.Fatalf("Failed to connect to AWS: %s", err.Error())
	} else {
		db = dynamodb.New(sess)
	}
}

func save(table *string, i *interface{}) error {

	item, err := dynamodbattribute.MarshalMap(&i)
	if err != nil {
		return err
	}

	fmt.Println(item)

	nowStr := strconv.FormatInt(time.Now().Unix(), 10)
	nowVal := &dynamodb.AttributeValue{N: &nowStr}
	item["unix"] = nowVal
	utcStr := time.Now().UTC().Format(time.RFC3339)
	utcVal := &dynamodb.AttributeValue{S: &utcStr}
	item["utc"] = utcVal

	if id, ok := item["pk"]; !ok || id.S == nil {
		s, _ := uuid.NewUUID()
		item["pk"] = &dynamodb.AttributeValue{S: aws.String(s.String())}
	}

	if _, err := db.PutItem(&dynamodb.PutItemInput{
		Item:      item,
		TableName: table,
	}); err != nil {
		return err
	} else {
		err = dynamodbattribute.UnmarshalMap(item, &i)
		fmt.Println(&i)
		return err
	}
}

func findOne(table, id *string, i interface{}) error {
	if out, err := db.GetItem(&dynamodb.GetItemInput{
		TableName: table,
		Key:       map[string]*dynamodb.AttributeValue{"pk": {S: id}},
	}); err != nil {
		return err
	} else {
		return dynamodbattribute.UnmarshalMap(out.Item, &i)
	}
}

func Handle(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	apigwp.LogRequest(r)

	token := r.Headers["Authorize"]

	table := r.PathParameters["table"]
	i := model.Registry[table]
	_ = json.Unmarshal([]byte(r.Body), &i)

	if r.Path == "save" {
		if err := save(&table, &i); err != nil {
			return apigwp.Bad(err)
		}
		return apigwp.OkInterface(token, i)
	}

	if r.Path == "find-one" {
		id := r.PathParameters["pk"]
		if err := findOne(&table, &id, &i); err != nil {
			return apigwp.Bad(err)
		}
		return apigwp.OkInterface(token, i)
	}

	return apigwp.Bad(fmt.Errorf("no path [%s]\n", r.Path))
}

func main() {
	lambda.Start(Handle)
}
