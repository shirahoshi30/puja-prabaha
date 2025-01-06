package models

type User struct {
	BaseModel
	FirstName   string   `json:"firstName"`
	LastName    string   `json:"lastname"`
	Username    string   `json:"username"`
	Email       string   `json:"email"`
	PhoneNumber int      `json:"phonenumber"`
	Password    string   `json:"password"`
	IsVerified  string   `json:"isverified"`
	Role        UserRole `json:"role"`
}

type UserRole string

const (
	AdminRole UserRole = "admin"
	userRole  UserRole = "user"
	GuestRole UserRole = "guest"
)
