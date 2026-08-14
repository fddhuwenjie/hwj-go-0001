# Expensebook

Expensebook is a small command-line expense ledger. It imports a CSV export, validates each expense, and prints a deterministic summary grouped by category.

## CSV format

Each record has four columns: date in YYYY-MM-DD format, category, amount in decimal currency units, and description. The first row is a header.

## Commands

Run the complete test suite with `go test ./...`.

Build the command with `go build ./...`.

Print a summary from a CSV file with `go run ./cmd/expensebook path/to/expenses.csv`.
