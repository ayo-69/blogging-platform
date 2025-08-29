package models

type Profile struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserID    string `json:"user_id"`
	Bio       string `json:"bio"`
}
