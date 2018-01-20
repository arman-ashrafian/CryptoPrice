$(function() {
    fetch("https://min-api.cryptocompare.com/data/pricemulti?fsyms=ETH&tsyms=BTC,USD")
    .then(function(resp) {
        return resp.json()
    })
    .then(function(data) {
        console.log(data)
    })
})