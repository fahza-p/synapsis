package model

type User struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	Password  string
	Name      string `json:"name"`
	Role      int64  `json:"role"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	CreatedBy string `json:"created_by,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
}
