package priceParser

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// TODO: implement currency detection
// Amount is represented as cents (1 USD = 100 cents)
type Price struct {
	Amount   uint
	Currency string
}

// Returns the price with two decimal places and currency (if any)
func (p *Price) String() string {
	if p.Currency == "" {
		return fmt.Sprintf("%.2f", float32(p.Amount/100))
	}
	return fmt.Sprintf("%.2f %s", float32(p.Amount/100), p.Currency)
}

func (p *Price) Float() float64 {
	return float64(p.Amount) / 100
}

// Given a string, returns a price struct
// This parser is a direct Go port from this Python implementation: https://github.com/hayj/SystemTools/blob/master/systemtools/number.py
// TODO: Add tests
// TODO: Optimize code
func PriceFromString(text string) (*Price, error) {
	if text == "" {
		return nil, errors.New("empty string")
	}

	text = strings.TrimSpace(text)
	re := regexp.MustCompile(`-?[0-9]*([,. ]?[0-9]+)+`)
	n := re.Find([]byte(text))
	if n == nil {
		return nil, errors.New("no number found")
	}
	n = []byte(strings.TrimSpace(string(n)))
	re = regexp.MustCompile(`.*[0-9]+.*`)
	if !re.Match([]byte(text)) {
		return nil, errors.New("no number found")
	}
	for strings.Contains(string(n), ",") && strings.Contains(string(n), ".") && strings.Contains(string(n), " ") {
		index := max(strings.LastIndex(string(n), ","), strings.LastIndex(string(n), " "), strings.LastIndex(string(n), "."))
		n = n[0:index]
	}
	n = []byte(strings.TrimSpace(string(n)))

	symbolsCount := 0
	for _, c := range []string{" ", ",", "."} {
		if strings.Contains(string(n), c) {
			symbolsCount++
		}
	}

	if symbolsCount == 1 {
		if strings.Contains(string(n), " ") {
			n = []byte(strings.ReplaceAll(string(n), " ", ""))
		} else {
			var theSymbol string
			if strings.Contains(string(n), ",") {
				theSymbol = ","
			} else {
				theSymbol = "."
			}
			if strings.Count(string(n), theSymbol) > 1 {
				n = []byte(strings.ReplaceAll(string(n), theSymbol, ""))
			} else {
				n = []byte(strings.ReplaceAll(string(n), theSymbol, "."))
			}
		}
	} else if symbolsCount > 1 {
		rightSymbolIndex := max(strings.LastIndex(string(n), ","), strings.LastIndex(string(n), " "), strings.LastIndex(string(n), "."))
		rightSymbol := string(n[rightSymbolIndex : rightSymbolIndex+1])
		if rightSymbol == " " {
			return PriceFromString(strings.ReplaceAll(string(n), " ", "_"))
		}
		n = []byte(strings.ReplaceAll(string(n), rightSymbol, "R"))
		leftSymbolIndex := max(strings.Index(string(n), ","), strings.Index(string(n), " "), strings.Index(string(n), "."))
		leftSymbol := string(n[leftSymbolIndex : leftSymbolIndex+1])
		n = []byte(strings.ReplaceAll(string(n), leftSymbol, "L"))
		n = []byte(strings.ReplaceAll(string(n), "L", ""))
		n = []byte(strings.ReplaceAll(string(n), "R", "."))
	}
	f, err := strconv.ParseFloat(string(n), 32)
	// rounds float to 2 decimal places
	f = math.Round(f*100) / 100
	if err != nil {
		return nil, err
	}
	// returns the value in cents
	amountInInt := uint(f * 100)
	price := &Price{
		Amount:   amountInInt,
		Currency: "", // TODO: implement currency detection
	}
	return price, nil
}

func max(args ...int) int {
	max := args[0]
	for _, arg := range args {
		if arg > max {
			max = arg
		}
	}
	return max
}
