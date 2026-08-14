package expensebook

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
	"time"
)

var ErrHeader = errors.New("invalid CSV header")

func ImportCSV(source io.Reader) ([]Expense, error) {
	reader := csv.NewReader(source)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("read expenses: %w", err)
	}
	if len(records) == 0 || len(records[0]) != 4 || strings.Join(records[0], ",") != "date,category,amount,description" {
		return nil, ErrHeader
	}
	entries := make([]Expense, 0, len(records)-1)
	for rowNumber, record := range records[1:] {
		if len(record) != 4 {
			return nil, fmt.Errorf("row %d: expected four columns", rowNumber+2)
		}
		date, err := time.Parse("2006-01-02", strings.TrimSpace(record[0]))
		if err != nil {
			return nil, fmt.Errorf("row %d date: %w", rowNumber+2, err)
		}
		amount, err := parseMoney(record[2])
		if err != nil {
			return nil, fmt.Errorf("row %d amount: %w", rowNumber+2, err)
		}
		category := strings.TrimSpace(record[1])
		if category == "" {
			return nil, fmt.Errorf("row %d: category is required", rowNumber+2)
		}
		entries = append(entries, Expense{Date: date, Category: category, Amount: amount, Description: strings.TrimSpace(record[3])})
	}
	return entries, nil
}

func parseMoney(value string) (Money, error) {
	value = strings.TrimSpace(value)
	amount, err := strconv.ParseFloat(value, 64)
	if err != nil || amount < 0 {
		return 0, errors.New("amount must be a non-negative decimal")
	}
	return Money(amount*100 + 0.5), nil
}

func Summarize(entries []Expense) Summary {
	totals := make(map[string]Money)
	var total Money
	for _, entry := range entries {
		total += entry.Amount
		totals[entry.Category] += entry.Amount
	}
	categories := make([]CategoryTotal, 0, len(totals))
	for category, amount := range totals {
		categories = append(categories, CategoryTotal{Category: category, Total: amount})
	}
	sort.Slice(categories, func(i, j int) bool { return categories[i].Category < categories[j].Category })
	return Summary{Total: total, Categories: categories}
}
