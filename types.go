package main

type CurrencyPrice struct {
	Date  string  `json:"date"`
	Price float32 `json:"rub"`
}

type CryptoPrice struct {
	Id           int    `json:"id"`
	Price        string `json:"price"`
	Qty          string `json:"qty"`
	QuoteQty     string `json:"quoteQty"`
	Time         int    `json:"time"`
	IsBuyerMaker bool   `json:"isBuyerMaker"`
	IsBestMatch  bool   `json:"isBestMatch"`
}
