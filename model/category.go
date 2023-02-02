package model

type Category struct {
	Id      string     `json:"id"`
	Name    string     `json:"name"`
	Slug    string     `json:"slug"`
	Product []*Product `json:"products"`
}
