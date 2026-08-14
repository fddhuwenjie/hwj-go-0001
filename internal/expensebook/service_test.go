package expensebook

import (
	"strings"
	"testing"
)

func TestImportCSVAndSummary(t *testing.T) {
	entries, err := ImportCSV(strings.NewReader("date,category,amount,description\n2026-08-01,food,12.50,lunch\n2026-08-02,travel,8.25,bus\n2026-08-03,food,3.25,snack\n"))
	if err != nil {
		t.Fatalf("ImportCSV() error = %v", err)
	}
	if len(entries) != 3 {
		t.Fatalf("ImportCSV() returned %d entries, want 3", len(entries))
	}
	summary := Summarize(entries)
	if summary.Total != 2400 {
		t.Fatalf("total = %d, want 2400", summary.Total)
	}
	if got := summary.Categories[0]; got.Category != "food" || got.Total != 1575 {
		t.Fatalf("food total = %#v, want 1575", got)
	}
}

func TestImportCSVAcceptsQuotedCommaInDescription(t *testing.T) {
	entries, err := ImportCSV(strings.NewReader("date,category,amount,description\n2026-08-04,office,15.00,\"paper, pens\"\n"))
	if err != nil {
		t.Fatalf("ImportCSV() error = %v", err)
	}
	if len(entries) != 1 || entries[0].Description != "paper, pens" {
		t.Fatalf("quoted description was not preserved: %#v", entries)
	}
}

func TestImportCSVRejectsInvalidRows(t *testing.T) {
	_, err := ImportCSV(strings.NewReader("date,category,amount,description\n2026-08-04,,15.00,missing category\n"))
	if err == nil {
		t.Fatal("ImportCSV() accepted an empty category")
	}
}
