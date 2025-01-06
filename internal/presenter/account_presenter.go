package presenter

import "github.com/golang-jwt/jwt"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Access  string            `json:"access"`
	Refresh string            `json:"refresh"`
	User    UserLoginResponse `json:"user"`
}

type NewUser struct {
	Username    string `json:"username" `
	Email       string `json:"email"`
	Password    string `json:"password"`
	Role        string `json:"role"`
	PhoneNumber int    `json:"phonenumber"`
}

type JwtCustomClaims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

type UpdateUsersRequest struct {
	ID          uint   `json:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Role        string `json:"role"`
	PhoneNumber string `json:"phonenumber"`
}

type UserLoginResponse struct {
	ID          uint   `json:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Role        string `json:"role"`
	PhoneNumber int    `json:"phonenumber"`
}

type UserCreateRequest struct {
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	PhoneNumber int    `json:"phonenumber"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Role        string `json:"role"`
}
type UserResponse struct {
	ID         uint   `json:"id"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Role       string `json:"role"`
	IsVerified bool   `json:"isverified"`
}

type UserListPresenter struct {
	ID       uint   `json:"id"`
	UserName string `json:"username"`
}

type UserListReq struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}
