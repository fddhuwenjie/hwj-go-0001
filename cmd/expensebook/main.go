package main

import (
	"fmt"
	"os"

	"github.com/fddhuwenjie/hwj-go-0001/internal/expensebook"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: expensebook FILE.csv")
		os.Exit(2)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	entries, err := expensebook.ImportCSV(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	summary := expensebook.Summarize(entries)
	fmt.Printf("total=%s\n", summary.Total.String())
	for _, item := range summary.Categories {
		fmt.Printf("%s=%s\n", item.Category, item.Total.String())
	}
}
