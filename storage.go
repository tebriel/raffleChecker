package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"os"
	"strings"
)

type RaffleTicket struct {
	TicketNumber string
}

var (
	Tickets = []string{
		"0136297",
		"0965302",
		"0223289",
		"0107291",
		"0911453",
		"0911452",
		"0911451",
		"0925406",
		"0222383",
		"0027490",
		"0224464",
		"0035561",
		"0924543",
		"0020730",
		"0119485",
		"0225621",
		"0032567",
		"0107296",
		"119436"}
	StorageType = USE_LOCAL
)

const (
	USE_DYNAMO = iota
	USE_LOCAL
)

func init() {
	if os.Getenv("STORAGE_TYPE") == "DYNAMO" {
		StorageType = USE_DYNAMO
	}
}

func findMatches(prefix string) []string {
	switch StorageType {
	case USE_DYNAMO:
		return findDynamoMatches(prefix)
	default:
		return findLocalMatches(prefix)
	}
}

func findLocalMatches(prefix string) []string {
	var matches []string
	for _, ticket := range Tickets {
		if strings.HasPrefix(ticket, prefix) {
			matches = append(matches, ticket)
		}
	}
	return matches
}

func findDynamoMatches(prefix string) []string {
	svc := dynamodb.New(session.New())
	input := &dynamodb.ScanInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":a": {
				S: aws.String(prefix),
			},
		},
		FilterExpression: aws.String("begins_with(TicketNumber, :a)"),
		TableName:        aws.String("raffle-tickets"),
	}

	result, err := svc.Scan(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return []string{}
	}
	var recs []RaffleTicket
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &recs)
	if err != nil {
		panic(fmt.Sprintf("failed to unmarshal Dynamodb Scan Items, %v", err))
	}
	var results []string
	for _, rec := range recs {
		results = append(results, rec.TicketNumber)
	}
	fmt.Println(result)
	return results
}
