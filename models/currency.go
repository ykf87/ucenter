// 币种
package models

type CurrencyModel struct {
	Id      int64   `json:"id"`
	Code    string  `json:"code"`
	Symbol  string  `json:"symbol"`
	Default int64   `json:"default"`
	Rate    float64 `json:"rate"`
}
