package storage

import (
	"testing"
)

func TestHandleRequest(t *testing.T) {
	scenarios := []struct {
		Prefix        string
		ExpectedCount int
	}{{
		"013",
		1,
	}, {
		Prefix:        "999",
		ExpectedCount: 0,
	},
	}

	for _, scenario := range scenarios {
		resp := findLocalMatches(scenario.Prefix)

		numMatches := len(resp)
		if numMatches != scenario.ExpectedCount {
			t.Errorf("Expected %d match but got %d", scenario.ExpectedCount, numMatches)
		}
	}
}
