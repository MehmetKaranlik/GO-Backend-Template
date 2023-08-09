package User

import (
	"Backend/Core/Utilities/Base"
	JWT "Backend/Core/Utilities/Jwt"
	"Backend/Product/Enums"
	"github.com/google/uuid"
	"time"
)

type User struct {
	Base.ResponseItem `json:"-" bson:"-"`
	Id                string
	Username          string
	Email             string
	Phone             string
	ProfileImage      string
	Password          []byte `json:"-"`
	AccessToken       string
	RefreshToken      string
	CreateDate        time.Time
	UpdateDate        time.Time
	Role              Enums.UserRole
}

func FactoryUser(email string, password []byte, role Enums.UserRole) *User {
	id := uuid.NewString()
	accessToken, _ := JWT.GenerateAccessToken(id)
	refreshToken, _ := JWT.GenerateRefreshToken(id)
	return &User{
		Id:           id,
		Email:        email,
		Password:     password,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		CreateDate:   time.Now(),
		UpdateDate:   time.Now(),
		Role:         role,
	}
}
