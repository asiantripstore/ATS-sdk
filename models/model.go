package models

type RoleType string

const (
	MEMBER     RoleType = "member"
	ADMIN      RoleType = "admin"
	UNVALIDATE RoleType = "unvalidate"
)

//User models for user
type User struct {
	ID     int      `json:"id" gorm:"-:primaryKey"`
	Email  string   `json:"email"`
	Secret string   `json:"secret"`
	Role   RoleType `json:"role"`
}
