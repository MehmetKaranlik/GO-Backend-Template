package Auth

import (
	"Backend/Core/Constants/Enums/Messages"
	"Backend/Core/Utilities/Hasher"
	JWT "Backend/Core/Utilities/Jwt"
	"Backend/Core/Utilities/Validation"
	"Backend/Product/Model/Auth"
	"Backend/Product/Model/User"
	"Backend/Product/Repositories"
	"errors"
	"io"
	"time"
)

type AuthService struct {
	UserRepository Repositories.IUserRepository
}

func (self *AuthService) Login(reader io.Reader) (*User.User, error) {
	var model Auth.LoginBody

	if err := Validation.NewCustomJsonDecoder(reader).Decode(&model); err != nil {
		return nil, errors.New(err.Error())
	}

	user := self.UserRepository.SecureAccess(model)
	if user == nil {
		return nil, errors.New(Messages.UserNotFound.ToString())
	}
	user.AccessToken, _ = JWT.GenerateAccessToken(user.Id)
	user.RefreshToken, _ = JWT.GenerateRefreshToken(user.Id)
	user.UpdateDate = time.Now()

	err := self.UserRepository.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (self *AuthService) Register(reader io.Reader) (*User.User, error) {
	var model Auth.RegisterBody
	if err := Validation.NewCustomJsonDecoder(reader).Decode(&model); err != nil {
		return nil, errors.New(err.Error())
	}
	user := self.UserRepository.ReadByEntry(model.Email)
	if user != nil {
		return nil, errors.New(Messages.EmailAlreadyInUse.ToString())
	}
	user = User.FactoryUser(model.Email, Hasher.EncryptPassword(model.Password), model.Role)
	err := self.UserRepository.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil

}
