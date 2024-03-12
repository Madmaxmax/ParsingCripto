package money

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func GetMoneyInfo() {
	resp, err := http.Get("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var currDict []Currency
	if err := json.Unmarshal(data, &currDict); err != nil {
		panic(err)
	}
	totalDict := make(map[string]Currency)
	for _, value := range currDict {
		totalDict[value.Name] = value
	}
	file, err := os.Create("currency.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	jsonData, err := json.Marshal(totalDict)
	if err != nil {
		panic(err)
	}

	_, err = file.Write(jsonData)
	if err != nil {
		panic(err)
	}
}
