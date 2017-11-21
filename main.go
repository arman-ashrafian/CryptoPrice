package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Price struct {
	BTC float64 `json:"BTC"`
	USD float64 `json:"USD"`
}

type Ether struct {
	Price Price `json:"ETH"`
}

func main() {
	fmt.Println("Ethereum Exchange Rate")

	resp, _ := http.Get("https://min-api.cryptocompare.com/data/pricemulti?fsyms=ETH&tsyms=BTC,USD")

	body, _ := ioutil.ReadAll(resp.Body)

	eth := Ether{}
	err := json.Unmarshal(body, &eth)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("USD: %.2f\n", eth.Price.USD)
	fmt.Printf("BTC: %.5f\n", eth.Price.BTC)
}
