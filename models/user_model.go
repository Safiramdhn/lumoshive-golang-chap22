package models

type User struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Email      string `json:"email,omitempty"`
	Password   string `json:"password,omitempty"`
	Token      string `json:"token,omitempty"`
	UserStatus string `json:"user_status,omitempty"`
}
