package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Define the metric
var (
	bitcoinPrice = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "bitcoin_price_usd",
		Help: "Current price of Bitcoin in USD",
	})
)

// Structure for JSON response from the API
type CoinGeckoResponse struct {
	Bitcoin struct {
		USD float64 `json:"usd"`
	} `json:"bitcoin"`
}

func fetchBitcoinPrice() {
	resp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd")
	if err != nil {
		log.Printf("Error fetching Bitcoin price: %v", err)
		return
	}
	defer resp.Body.Close()

	var priceResponse CoinGeckoResponse
	if err := json.NewDecoder(resp.Body).Decode(&priceResponse); err != nil {
		log.Printf("Error decoding response: %v", err)
		return
	}

	bitcoinPrice.Set(priceResponse.Bitcoin.USD)
	log.Printf("Bitcoin price updated: %f USD", priceResponse.Bitcoin.USD)
}

func main() {
	// Register the metric
	prometheus.MustRegister(bitcoinPrice)

	// register right away
	fetchBitcoinPrice()
	// Set up a ticker to fetch price every 60 seconds
	ticker := time.NewTicker(60 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				fetchBitcoinPrice()
			}
		}
	}()

	// Expose the registered metrics via HTTP
	http.Handle("/metrics", promhttp.Handler())
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}
