package model

import (
	"html"
	"time"

	"github.com/fahza-p/synapsis/lib/utils"
)

/* Request */
type AuthSigninReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=50"`
}

type AuthSignupReq struct {
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=6,max=50,eqfield=ConfirmPassword"`
	Name            string `json:"name" validate:"omitempty,max=120"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,min=6,max=50,eqfield=Password"`
}

/* Model Data */
type AuthUserData struct {
	Id        int64  `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	Role      int64  `json:"role"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	CreatedBy string `json:"created_by,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
}

/* Function */
func (u *AuthUserData) SetUserSignupData() error {
	hash, err := utils.HashBcrypt(u.Password)
	if err != nil {
		return err
	}

	now := time.Now().Format("2006-01-02 15:04:05")
	u.Email = html.EscapeString(u.Email)
	u.Name = html.EscapeString(u.Name)
	u.Role = 2
	u.CreatedAt = now
	u.UpdatedAt = now
	u.CreatedBy = "sys_registration"
	u.UpdatedBy = "sys_registration"
	u.Password = hash

	return nil
}
