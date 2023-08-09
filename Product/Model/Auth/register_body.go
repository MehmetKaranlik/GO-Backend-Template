package Auth

import (
	"Backend/Product/Enums"
)

type RegisterBody struct {
	Email    string         `json:"email" validate:"required"`
	Password string         `json:"password" validate:"required,min=8,max=32"`
	Role     Enums.UserRole `json:"role" validate:"required"`
}
