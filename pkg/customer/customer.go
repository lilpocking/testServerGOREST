package customer

type Customer struct {
	Id           int    `json:"CustomerId"`
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	Company      string `json:"Company"`
	Address      string `json:"Address"`
	City         string `json:"City"`
	State        string `json:"State"`
	Country      string `json:"Country"`
	PostalCode   string `json:"PostalCode"`
	Phone        string `json:"Phone"`
	Fax          string `json:"Fax"`
	Email        string `json:"Email"`
	SupportRepId int    `json:"SupportRepId"`
}
