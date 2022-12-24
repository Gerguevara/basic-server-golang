package types

import "encoding/json"

//la notacion `json:"name"` le deice a este struct que cuando sea parseado (serialize) Json esa propiedad tomara ese nombre

type User struct {
	Name  string `json:"name"`
	Email string `json:"email_address"`
	Phone string `json:"phone_number"`
}

// pasea el usuario a json
func (u *User) ToJson() ([]byte, error) {
	// marshal permite enodear un Struct a Json
	return json.Marshal(u)
}
