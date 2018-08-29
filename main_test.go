package main

import (
	"context"
	"testing"
)

func TestHandleRequest(t *testing.T) {
	ctx := context.Background()
	scenarios := []struct {
		Event RaffleEvent
		ExpectedCount int
	}{{
		RaffleEvent{TicketNum: "013"},
		1,
	},{
		Event: RaffleEvent{TicketNum: "999"},
		ExpectedCount: 0,
	},
	}

	for _, scenario := range scenarios {
		resp, err := HandleRequest(ctx, scenario.Event)
		if err != nil {
			t.Errorf("Error #%v", err)
		}


		numMatches := len(resp.TicketNumbers)
		if numMatches != scenario.ExpectedCount {
			t.Errorf("Expected %d match but got %d", scenario.ExpectedCount, numMatches)
		}
	}
}
