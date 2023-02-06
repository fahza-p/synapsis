package model

import (
	"html"
	"time"
)

var CategoryFilter = []string{"name", "slug"}

/* Request */
type CategoryCreateReq struct {
	Name  string `json:"name" validate:"required,ascii,max=100"`
	Image string `json:"image,omitempty" validate:"omitempty,url"`
}

type CategoryUpdateReq struct {
	Name  string `json:"name" validate:"omitempty,ascii,max=100"`
	Image string `json:"image,omitempty" validate:"omitempty,url"`
}

/* Model Data */
type Category struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	Image     string `json:"image"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	CreatedBy string `json:"created_by,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
}

/* Functions */
func (u *Category) SetCategoryCreateData(userEmail string) {
	now := time.Now().Format("2006-01-02 15:04:05")

	u.Name = html.EscapeString(u.Name)
	u.Slug = html.EscapeString(u.Slug)
	u.Image = html.EscapeString(u.Image)
	u.CreatedAt = now
	u.UpdatedAt = now
	u.CreatedBy = userEmail
	u.UpdatedBy = userEmail
}
