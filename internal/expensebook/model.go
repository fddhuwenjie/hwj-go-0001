package expensebook

import (
	"fmt"
	"math"
	"time"
)

type Money int64

func (m Money) String() string {
	return fmt.Sprintf("%d.%02d", m/100, int64(math.Abs(float64(m%100))))
}

type Expense struct {
	Date        time.Time
	Category    string
	Amount      Money
	Description string
}

type CategoryTotal struct {
	Category string
	Total    Money
}

type Summary struct {
	Total      Money
	Categories []CategoryTotal
}
