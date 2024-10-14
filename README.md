# Bitcoin Price Exporter

A simple Prometheus exporter written in Go that fetches the current price of Bitcoin in USD from the CoinGecko API.

## Features

- Fetches Bitcoin price every 60 seconds.
- Exposes metrics in a format compatible with Prometheus.

## Getting Started

### Prerequisites

- Go (1.18+)
### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/pefman/prometheus-crypto-exporter.git
   cd bitcoin-price-exporter
2. Build the application:
   ```bash
   go build
3. Run the exporter:
   ```bash
   ./bitcoin-price-exporter
4. Access Metrics
   ```bash
   http://localhost:8080/metrics
5. Prometheus Integration
To scrape this exporter, add the following configuration to your prometheus.yml:
   ```bash
   scrape_configs:
   - job_name: 'bitcoin_price_exporter'
       static_configs:
       - targets: ['localhost:8080']
