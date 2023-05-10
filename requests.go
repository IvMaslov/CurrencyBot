package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetCurrency(from string) (CurrencyPrice, error) {
	url := "https://cdn.jsdelivr.net/gh/fawazahmed0/currency-api@1/latest/currencies/" + from + "/rub.json"

	curPrice := CurrencyPrice{}
	resp, err := http.Get(url)
	if err != nil {
		return curPrice, fmt.Errorf("something wrong with API")
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&curPrice)
	if err != nil {
		return curPrice, fmt.Errorf("something wrong with API")
	}

	return curPrice, nil
}

func GetCrypto(val string) (CryptoPrice, error) {
	url := "https://api.binance.us/api/v3/trades?symbol=" + val + "USDT&limit=1"

	data := []CryptoPrice{}
	resp, err := http.Get(url)
	if err != nil {
		return CryptoPrice{}, fmt.Errorf("something wrong with API")
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return CryptoPrice{}, fmt.Errorf("something wrong with API")
	}

	return data[0], nil
}
