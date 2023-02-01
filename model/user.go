package model

type User struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	Password  string
	Name      string `json:"name"`
	Role      string `json:"role"`
	CreatedAt int64  `json:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
	CreatedBy string `json:"created_by,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
}
