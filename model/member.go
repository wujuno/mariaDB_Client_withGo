package model

type Member struct {
	Id        string `json:"memberId"`
	Name      string `json:"memberName"`
	Number    int8   `json:"memberNumber"`
	Addr      string `json:"memberAddr"`
	Phone1    string `json:"memberPhone1"`
	Phone2    string `json:"memberPhone2"`
	Height    uint8  `json:"memberHeight"`
	DebutDate string `json:"debutDate"`
}