package types

type Benefit struct {
	Id         string `json:"id"`
	Name       string `json:"name,omitempty"`
	Expiration string `json:"expiration,omitempty"`
}
