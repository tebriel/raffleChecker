package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tebriel/raffleChecker/pkg/storage"
)

// RaffleEvent contains a ticket number for matching
type RaffleEvent struct {
	TicketNum string `json:"ticket_num"`
}

// RaffleResponse is the json response for the ticket search
type RaffleResponse struct {
	TicketNumbers []string `json:"ticket_numbers"`
}

// HandleRequest is called by the aws lambda
func HandleRequest(ctx context.Context, event RaffleEvent) (RaffleResponse, error) {
	response := RaffleResponse{TicketNumbers: storage.FindMatches(event.TicketNum)}

	return response, nil
}

func main() {
	lambda.Start(HandleRequest)
}
