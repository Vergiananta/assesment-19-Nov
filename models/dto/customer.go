package dto

type CustomerRequest struct {
	Name 		string
	Address 	string
	Email 		string
	Password 	string
}

type CustomerResponse struct {
	Name 		string
	Address 	string
	Email 		string
}