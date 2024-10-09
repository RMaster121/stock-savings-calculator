package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"stock-savings-calculator/models"
	"time"
)

const POLYGON_STOCK_URL = "https://api.polygon.io/v2/aggs/ticker/%s/range/1/week/%s/%s?adjusted=true&apiKey=%s"
const POLYGON_FOREX_URL = "https://api.polygon.io/v2/aggs/ticker/C:%s/range/1/week/%s/%s?apiKey=%s"

// SendRequest fetches stock or forex data from Polygon.io
func SendRequest(query string, forex bool, from string, to string, api_key string) []models.StockData {
	var apiUrl string
	if forex {
        apiUrl = fmt.Sprintf(POLYGON_FOREX_URL, query, from, to, api_key)
	} else {
        apiUrl = fmt.Sprintf(POLYGON_STOCK_URL, query, from, to, api_key)
	}

	resp, err := http.Get(apiUrl)
	if err != nil {
		log.Fatalf("Failed to send HTTP GET request: %v", err)
	}
	defer resp.Body.Close()

	rawData := checkAndExtractRawData(resp)
	processedData := ProcessData(rawData, forex)

	return processedData
}

func checkAndExtractRawData(resp *http.Response) []byte {
	if resp.StatusCode != 200 {
		log.Fatalf("HTTP GET Status: %v", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if len(body) == 0 {
		log.Fatal("HTTP GET body empty")
	}

	return body
}

func ProcessData(raw []byte, forex bool) []models.StockData {
	if forex {
		return processForexData(raw)
	}
	return processStockData(raw)
}

func processStockData(raw []byte) []models.StockData {
	var result struct {
		Results []struct {
			Ticker       string    `json:"T"`
			Close        float64   `json:"c"`
			Timestamp    int64     `json:"t"`
		} `json:"results"`
	}
	err := json.Unmarshal(raw, &result)
	if err != nil {
		log.Fatal(err)
	}

	stockData := make([]models.StockData, 0)
	for _, res := range result.Results {
		date := time.Unix(res.Timestamp/1000, 0).Format("2006-01-02")
		stockData = append(stockData, models.StockData{Date: date, Price: res.Close})
	}

	sort.Slice(stockData, func(i, j int) bool {
		return stockData[i].Date < stockData[j].Date
	})

	return stockData
}

func processForexData(raw []byte) []models.StockData {
    var result struct {
        Results []struct {
            Close     float64 `json:"c"`
            Timestamp int64   `json:"t"`
        } `json:"results"`
    }
    err := json.Unmarshal(raw, &result)
    if err != nil {
        log.Fatal(err)
    }

    forexData := make([]models.StockData, 0)
    for _, res := range result.Results {
        date := time.Unix(res.Timestamp/1000, 0).Format("2006-01-02")
        forexData = append(forexData, models.StockData{Date: date, Price: res.Close})
    }

    sort.Slice(forexData, func(i, j int) bool {
        return forexData[i].Date < forexData[j].Date
    })

    return forexData
}

