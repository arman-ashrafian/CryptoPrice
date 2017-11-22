package main

import (
	"encoding/json"
	"fmt"
	"html/template"
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

type IndexPage struct {
	Title string
	Eth   Ether
}

func main() {
	fmt.Println("Starting...\n")

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(rw http.ResponseWriter, req *http.Request) {
	// get Ethereum exchange rate
	resp, _ := http.Get("https://min-api.cryptocompare.com/data/pricemulti?fsyms=ETH&tsyms=BTC,USD")
	body, _ := ioutil.ReadAll(resp.Body)

	eth := Ether{}
	json.Unmarshal(body, &eth)

	p := &IndexPage{Title: "Ethereum Price", Eth: eth}

	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(rw, p)
}
