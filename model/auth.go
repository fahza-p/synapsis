package model

/* Request */
type AuthSigninReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=50"`
}

type AuthSignupReq struct {
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=6,max=50,eqfield=ConfirmPassword"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,min=6,max=50,eqfield=Password"`
}

/* Model Data */
type AuthUserData struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	CreatedBy string `json:"created_by,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
}
