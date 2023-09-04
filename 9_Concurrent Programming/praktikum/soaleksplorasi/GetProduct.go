package soaleksplorasi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Product struct {
	Id          int     `json:"id"`
	Title       string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
}

func Running() {
	channel := make(chan Product)
	go GetProduct(channel)

	for prodct := range channel {
		fmt.Println(prodct.Title, prodct.Price, prodct.Category)
	}
}

func GetProduct(chanel chan Product) {
	req, err := http.NewRequest("GET", "https://fakestoreapi.com/products", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	products := []Product{}
	err = json.Unmarshal(data, &products)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, prdct := range products {
		chanel <- prdct
	}
}
