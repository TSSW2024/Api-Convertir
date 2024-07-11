package currency

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func ConvertCurrency(amount float64, from string, to string) (float64, error) {
	apiKey := "664041a8639bd5d3bdcd" // Replace with your CurrencyConverterAPI key

	from = url.QueryEscape(from)
	to = url.QueryEscape(to)

	apiURL := fmt.Sprintf("https://free.currconv.com/api/v7/convert?q=%s_%s&compact=ultra&apiKey=%s", from, to, apiKey)

	resp, err := http.Get(apiURL)
	if err != nil {
		return 0, fmt.Errorf("error making API request: %v", err)
	}
	defer resp.Body.Close()

	var result map[string]float64
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return 0, fmt.Errorf("error decoding JSON response: %v", err)
	}

	query := fmt.Sprintf("%s_%s", from, to)
	conversionRate, ok := result[query]
	if !ok {
		return 0, fmt.Errorf("conversion rate not found for %s", query)
	}

	convertedAmount := amount * conversionRate

	return convertedAmount, nil
}
