package storage

import "strings"

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
)

func findLocalMatches(prefix string) []string {
	var matches []string
	for _, ticket := range Tickets {
		if strings.HasPrefix(ticket, prefix) {
			matches = append(matches, ticket)
		}
	}
	return matches
}
