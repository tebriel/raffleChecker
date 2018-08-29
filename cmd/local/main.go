package main

import (
	"encoding/json"
	"fmt"
	"github.com/tebriel/raffleChecker/pkg/storage"
)

// RaffleResponse is the json response for the ticket search
type RaffleResponse struct {
	TicketNumbers []string `json:"ticket_numbers"`
}

func main() {
	resp := RaffleResponse{TicketNumbers: storage.FindMatches("013")}
	respBytes, _ := json.Marshal(resp)
	fmt.Printf("%v", string(respBytes))
}
