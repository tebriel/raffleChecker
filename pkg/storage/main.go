package storage

import (
	"os"
)

type RaffleTicket struct {
	TicketNumber string
}

var (
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

// FindMatches uses the specified storage type to find matches
func FindMatches(prefix string) []string {
	switch StorageType {
	case USE_DYNAMO:
		return findDynamoMatches(prefix)
	default:
		return findLocalMatches(prefix)
	}
}
