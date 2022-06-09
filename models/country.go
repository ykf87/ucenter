package models

type CountryModel struct {
	Id        int64   `json:"id"`
	Iso3      string  `json:"iso3"`
	Iso2      string  `json:"iso2"`
	Phonecode string  `json:"phonecode"`
	Currency  string  `json:"currency"`
	Region    int64   `json:"region"`
	Subregion int64   `json:"subregion"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Emoji     string  `json:"emoji"`
}
