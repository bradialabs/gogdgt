package tpl

func UserModelTemplate() []byte {
	return []byte(`package {{ .PkgName }}

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User struct for database storage
type User struct {
	ID    primitive.ObjectID ` + "`" + `json:"id,omitempty" bson:"_id"` + "`" + `
	First string             ` + "`" + `json:"first,omitempty" bson:"first,omitempty"` + "`" + `
	Last  string             ` + "`" + `json:"last,omitempty" bson:"last,omitempty"` + "`" + `
	Email string             ` + "`" + `json:"email,omitempty" bson:"email,omitempty"` + "`" + `

	PasswordDigest string ` + "`" + `json:"-" bson:"password_digest"` + "`" + `

	CreatedAt time.Time ` + "`" + `json:"createdAt" bson:"created_at"` + "`" + `
}

// CreateUser is a struct for passing new user information
type CreateUser struct {
	First           string ` + "`" + `json:"first"` + "`" + `
	Last            string ` + "`" + `json:"last"` + "`" + `
	Email           string ` + "`" + `json:"email"` + "`" + `
	Password        string ` + "`" + `json:"password"` + "`" + `
	PasswordConfirm string ` + "`" + `json:"passwordConfirm"` + "`" + `
}

// UserService interface for working with user data
type UserService interface {
	// List() ([]*User, error)
	View(id primitive.ObjectID) (*User, error)
	Create(user *CreateUser) (*User, error)
	// Update(user *User) error
	// Delete(id int) error
}
`)
}

func AuthModelTemplate() []byte {
	return []byte(`package {{ .PkgName }}

// LoginData for receiver user login details
type LoginData struct {
	Email    string ` + "`" + `json:"email"` + "`" + `
	Password string ` + "`" + `json:"password"` + "`" + `
}

// LoginResponse data for returning a new jwt token
type LoginResponse struct {
	Token string ` + "`" + `json:"token"` + "`" + `
	User  *User  ` + "`" + `json:"user"` + "`" + `
}

// AuthService for working with authentication of users
type AuthService interface {
	Login(loginData *LoginData) (bool, *User)
}`)
}
