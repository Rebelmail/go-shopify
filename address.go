package goshopify

type Address struct {
	ID           int     `json:"id"`
	Address1     string  `json:"address1"`
	Address2     string  `json:"address2"`
	City         string  `json:"city"`
	Company      string  `json:"company"`
	Country      string  `json:"country"`
	CountryCode  string  `json:"country_code"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Name         string  `json:"name"`
	Phone        string  `json:"phone"`
	Province     string  `json:"province"`
	ProvinceCode string  `json:"province_code"`
	Zip          string  `json:"zip"`
	Default      bool    `json:"default"`
}
