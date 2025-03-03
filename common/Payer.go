package common

type Payer struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Identity string `json:"identity,omitempty"`
}
