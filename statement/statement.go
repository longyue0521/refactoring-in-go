package statement

import (
	"errors"
	"fmt"
	"math"
)

type Performance struct {
	PlayID   string
	Audience int
}

type Invoice struct {
	Customer     string
	Performances []Performance
}

type Play struct {
	Name string
	Type string
}

type Plays map[string]Play

func Statement(invoice *Invoice, plays Plays) (string, error) {
	totalAmount := 0
	volumeCredits := float64(0)

	result := fmt.Sprintf("Statement for %s\n", invoice.Customer)

	for _, performance := range invoice.Performances {

		play := plays[performance.PlayID]
		thisAmount, err := amountFor(performance, play)
		if err != nil {
			return "", err
		}

		// add volume credits
		volumeCredits += math.Max(float64(performance.Audience-30), 0)

		// add extra credit for every ten comedy attendees
		if "comedy" == play.Type {
			volumeCredits += math.Floor(float64(performance.Audience / 5))
		}

		// print line for this order
		result += fmt.Sprintf("  %s: $%.2f (%d seats)\n", play.Name, float64(thisAmount/100), performance.Audience)
		totalAmount += thisAmount
	}

	result += fmt.Sprintf("Amount owed is $%.2f\n", float64(totalAmount/100))
	result += fmt.Sprintf("You earned %d credits", int64(volumeCredits))

	return result, nil
}

func amountFor(performance Performance, play Play) (int, error) {
	result := 0
	switch play.Type {
	case "tragedy":
		result = 40000
		if performance.Audience > 30 {
			result += 1000 * (performance.Audience - 30)
		}
	case "comedy":
		result = 30000
		if performance.Audience > 20 {
			result += 10000 + 500*(performance.Audience-20)
		}
		result += 300 * performance.Audience
	default:
		return 0, errors.New(`unknown type: ` + play.Type)
	}
	return result, nil
}
