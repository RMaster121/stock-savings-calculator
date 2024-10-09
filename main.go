package main

import (
	"flag"
	"fmt"
	"stock-savings-calculator/api"
	"time"
)

func main() {
	stock_query := flag.String("stock", "NVDA", "Stock query")
	forex_query_raw := flag.String("forex-query", "PLN", "Domesic currency (e.g. PLN)")
	forex_query := *forex_query_raw + "USD"
	forex_query_reverse := "USD" + *forex_query_raw
	// This number means how many times we want to put our money in stock
	number_of_weeks_back := flag.Int("weeks-ago", 5, "Number of weeks back (There will be +1 week, because we want current week as reference for comparision)")
	local_money := flag.Int("money", 100, "Amount of money to put into stock every week")
	api_key := flag.String("api-key", "", "API key for Polygon.io")
	flag.Parse()

	from_date := time.Now().AddDate(0, 0, -7*(*number_of_weeks_back))
	to_date := time.Now()

	from_date_str := from_date.Format("2006-01-02")
	to_date_str := to_date.Format("2006-01-02")

	fmt.Println("Stock:", *stock_query)
	fmt.Println("From date:", from_date_str)
	fmt.Println("To date:", to_date_str)
	fmt.Println("Money every week:", *local_money, "PLN")
	fmt.Println()

	price_data := api.SendRequest(*stock_query, false, from_date_str, to_date_str, *api_key)
	forex_data := api.SendRequest(forex_query, true, from_date_str, to_date_str, *api_key)
	forex_data_reverse := api.SendRequest(forex_query_reverse, true, from_date_str, to_date_str, *api_key)

	var stock_count float64
	stock_count = 0

	for i := 0; i < *number_of_weeks_back; i++ {
		stock_count += float64(*local_money) * forex_data[i].Price / price_data[i].Price
		fmt.Println("Week:", i+1)
		fmt.Println("Date:", price_data[i].Date)
		fmt.Println("Stock count:", stock_count)
		fmt.Println("Stock price:", price_data[i].Price, "USD")
		fmt.Println("PLN -> USD price:", forex_data[i].Price)
		fmt.Println("USD -> PLN price:", forex_data_reverse[i].Price)
		fmt.Println("Value:", stock_count*price_data[i].Price*forex_data_reverse[i].Price, "PLN")
		fmt.Println("Invested:", float64(i+1)*float64(*local_money), "PLN")
		if i != 0 {
			earnings := stock_count*price_data[i].Price*forex_data_reverse[i].Price - float64(i+1)*float64(*local_money)
			fmt.Println("Earnings:", earnings, "PLN")
		}
		fmt.Println()
	}

	total_invested := float64(*number_of_weeks_back) * float64(*local_money)
	total_value := stock_count * price_data[*number_of_weeks_back].Price * forex_data_reverse[*number_of_weeks_back].Price
	earnings := total_value - total_invested
	fmt.Println("Total invested:", total_invested, "PLN")
	fmt.Println("Total value:", total_value, "PLN")
	fmt.Println("Earnings:", earnings, "PLN")
}
