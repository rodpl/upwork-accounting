package nbp

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const nbpUrl = "http://api.nbp.pl/api/exchangerates/rates/"

type tableARate struct {
	Table    string `json:"table"`
	Currency string `json:"currency"`
	Code     string `json:"code"`
	Rates    []struct {
		No            string  `json:"no"`
		EffectiveDate string  `json:"effectiveDate"`
		Mid           float64 `json:"mid"`
	} `json:"rates"`
}

func GetExchangeRate(currencyCode string, at time.Time) float64 {
	url := nbpUrl + "A/" + currencyCode
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return -1
	}
	client := &http.Client{}
	resp, _ := client.Do(req)

	defer resp.Body.Close()

	var record tableARate

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	return record.Rates[0].Mid
}
