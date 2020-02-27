package models

//InputData - this model is used for request body
type InputData struct {
	A int `json:"a"`
	B int `json:"b"`
}

//Product - this is for output response body
type Product struct {
	A uint64 `json:"a!"`
	B uint64 `json:"b!"`
}
