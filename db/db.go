package db

type Customer struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Contacted *bool  `json:"contacted"`
}

var T = true
var F = false

var Customers = map[int]Customer{
	1: {
		ID:        1,
		Name:      "John Doe",
		Role:      "Senior Manager",
		Email:     "jodoe@example.com",
		Phone:     "1234567890",
		Contacted: &T,
	},
	2: {
		ID:        2,
		Name:      "Jane Doe",
		Role:      "Sales Manager",
		Email:     "jadoe@example.com",
		Phone:     "0987654321",
		Contacted: &F,
	},
	3: {
		ID:        3,
		Name:      "Jack Doe",
		Role:      "VP",
		Email:     "jkdoe@example.com",
		Phone:     "0987654321",
		Contacted: &F,
	},
}
