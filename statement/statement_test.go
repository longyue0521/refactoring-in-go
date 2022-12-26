package statement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatement(t *testing.T) {
	testCases := []struct {
		name     string
		invoice  *Invoice
		plays    Plays
		wantStat string
		wantErr  error
	}{
		{
			name: "normal case",
			invoice: &Invoice{
				Customer: "BigCo",
				Performances: []Performance{
					{
						PlayID:   "hamlet",
						Audience: 55,
					},
					{
						PlayID:   "as-like",
						Audience: 35,
					},
					{
						PlayID:   "othello",
						Audience: 40,
					},
				},
			},
			plays: map[string]Play{
				"hamlet":  {Name: "Hamlet", Type: "tragedy"},
				"as-like": {Name: "As You Like It", Type: "comedy"},
				"othello": {Name: "Othello", Type: "tragedy"},
			},
			wantStat: `Statement for BigCo
  Hamlet: $650.00 (55 seats)
  As You Like It: $580.00 (35 seats)
  Othello: $500.00 (40 seats)
Amount owed is $1730.00
You earned 47 credits`,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			statement, err := Statement(tc.invoice, tc.plays)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantStat, statement)
		})
	}
}
