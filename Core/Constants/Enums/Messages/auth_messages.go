package Messages

type AuthResponseMessages string

const (
	LoginSuccess                AuthResponseMessages = "Login success"
	LoginFailed                 AuthResponseMessages = "Login failed"
	LoginFailedUserNotFound     AuthResponseMessages = "Login failed, user not found"
	LoginFailedPasswordNotMatch AuthResponseMessages = "Login failed, password not match"
	LoginFailedEmailNotMatch    AuthResponseMessages = "Login failed, email not match"
	UserNotFound                AuthResponseMessages = "User not found or password is wrong."
	RegisterSuccess             AuthResponseMessages = "Register success"
	RegisterFailed              AuthResponseMessages = "Register failed"
	UserAlreadyExists           AuthResponseMessages = "User already exists"
	EmailAlreadyInUse           AuthResponseMessages = "Email already in use!"
)

func (m AuthResponseMessages) ToString() string {
	return string(m)
}
