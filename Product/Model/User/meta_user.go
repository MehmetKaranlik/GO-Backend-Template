package User

type MetaUser struct {
	Id           string
	Username     string
	ProfileImage string
}

func NewMetaUser(id string, username string, profileImage string) *MetaUser {
	return &MetaUser{
		Id:           id,
		Username:     username,
		ProfileImage: profileImage,
	}
}

func NewMetaUserFromUser(user User) *MetaUser {
	return &MetaUser{
		Id:           user.Id,
		Username:     user.Username,
		ProfileImage: user.ProfileImage,
	}
}
