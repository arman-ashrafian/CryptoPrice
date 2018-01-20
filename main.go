package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

const (
	ethExRate    = "https://min-api.cryptocompare.com/data/pricemulti?fsyms=ETH&tsyms=BTC,USD"
	ethAddress   = "0xaddfc1233fe9909e159715ac179a6ba4a470a451"
	ethplorerAPI = "freekey"
)

type price struct {
	BTC float64 `json:"BTC"`
	USD float64 `json:"USD"`
}

type ether struct {
	Price price `json:"ETH"`
}

type indexPage struct {
	Title string
	ETH   ether
}

func main() {
	PORT := ":8080"
	fmt.Printf("Starting server ... localhost%s\n\n", PORT)

	http.HandleFunc("/", indexHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static")))) // serve static files
	http.ListenAndServe(":8080", nil)
}

func indexHandler(rw http.ResponseWriter, req *http.Request) {
	// get Ethereum exchange rate
	resp, _ := http.Get(ethExRate)
	body, _ := ioutil.ReadAll(resp.Body)

	eth := ether{}
	json.Unmarshal(body, &eth)

	p := &indexPage{Title: "Ethereum Price", ETH: eth}

	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(rw, p)
}
