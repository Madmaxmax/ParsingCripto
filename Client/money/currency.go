package money

import (
	"encoding/json"
	"os"
	"strings"
	"time"
)

type Currency struct {
	Ath                          float64     `json:"ath"`
	AthChangePercentage          float64     `json:"ath_change_percentage"`
	AthDate                      time.Time   `json:"ath_date"`
	Atl                          float64     `json:"atl"`
	AtlChangePercentage          float64     `json:"atl_change_percentage"`
	AtlDate                      time.Time   `json:"atl_date"`
	CirculatingSupply            float64     `json:"circulating_supply"`
	CurrentPrice                 float64     `json:"current_price"`
	FullyDilutedValuation        float64     `json:"fully_diluted_valuation"`
	High24H                      float64     `json:"high_24h"`
	Id                           string      `json:"id"`
	Image                        string      `json:"image"`
	LastUpdated                  time.Time   `json:"last_updated"`
	Low24H                       float64     `json:"low_24h"`
	MarketCap                    float64     `json:"market_cap"`
	MarketCapChange24H           float64     `json:"market_cap_change_24h"`
	MarketCapChangePercentage24H float64     `json:"market_cap_change_percentage_24h"`
	MarketCapRank                float64     `json:"market_cap_rank"`
	MaxSupply                    interface{} `json:"max_supply"`
	Name                         string      `json:"name"`
	PriceChange24H               float64     `json:"price_change_24h"`
	PriceChangePercentage24H     float64     `json:"price_change_percentage_24h"`
	Roi                          interface{} `json:"roi"`
	Symbol                       string      `json:"symbol"`
	TotalSupply                  float64     `json:"total_supply"`
	TotalVolume                  float64     `json:"total_volume"`
}
type Search struct {
	SearchKey string
	Results   map[string]Currency
}

func UpdateMoneyInfo() {
	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			GetMoneyInfo()
		}
	}
}

func ReadCurrency() map[string]Currency {
	file, err := os.ReadFile("currency.json")
	if err != nil {
		panic(err)
	}
	data := make(map[string]Currency)
	err = json.Unmarshal(file, &data)
	if err != nil {
		panic(err)
	}
	return data
}

func FindMatches(currency string, dict map[string]Currency) map[string]Currency {
	find := strings.ToLower(currency)
	if currency == "" {
		return dict
	}
	findDict := make(map[string]Currency)
	for key, value := range dict {
		if strings.Contains(strings.ToLower(key), find) {
			findDict[key] = value
		}
	}
	return findDict
}
