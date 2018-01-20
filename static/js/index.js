$(function() {
    const ethAdd = "0xaddfc1233fe9909e159715ac179a6ba4a470a451"
    const ethPriceAPI = "https://min-api.cryptocompare.com/data/pricemulti?fsyms=ETH&tsyms=BTC,USD"
    const ethWalletAPI = `https://api.ethplorer.io/getAddressInfo/${ethAdd}?apiKey=freekey`

    // ETH Price
    fetch(ethPriceAPI)
    .then(function(resp) {
        return resp.json()
    })
    .then(function(data) {
        console.log(data)
    })

    // ETH Wallet
    fetch(ethWalletAPI)
    .then((resp) => {
        return resp.json()
    })
    .then((data) => {
        console.log(data)
    })
})