package models

type StockData struct {
	Date string
	Price float64 // Generally, it's better NOT to use float for money (rounding error), but for the sake of simplicity, we'll use it here
}