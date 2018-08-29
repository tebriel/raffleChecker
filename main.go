package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"strings"
)

// RaffleEvent contains a ticket number for matching
type RaffleEvent struct {
	TicketNum string `json:"ticket_num"`
}

// RaffleResponse is the json response for the ticket search
type RaffleResponse struct {
	TicketNumbers []string `json:"ticket_numbers"`
}

var Tickets = []string{
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

// HandleRequest is called by the aws lambda
func HandleRequest(ctx context.Context, event RaffleEvent) (RaffleResponse, error) {
	response := RaffleResponse{TicketNumbers:findMatches(event.TicketNum)}

	return response, nil
}

func findMatches(prefix string) ([]string) {
	var matches []string
	for _, ticket := range Tickets {
		if strings.HasPrefix(ticket, prefix) {
			matches = append(matches, ticket)
		}
	}

	return matches
}

func main() {
	lambda.Start(HandleRequest)
}
