package types

type User struct {
	Base
	Username    string `gorm:"unique"`
	FirstName   string
	LastName    string
	Address     string
	Salt        string
	Hash        string
	AdminRights bool
}

type UserLoginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserCreateDto struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Address     string `json:"address"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	AdminRights bool   `json:"admin_rights"`
	Token       string `json:"token"`
}

type UserDetailsDto struct {
	Username    string `json:"username"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Address     string `json:"address"`
	AdminRights bool   `json:"admin_rights"`
}

type UserChangePasswordDto struct {
	Token       string `json:"token"`
	NewPassword string `json:"new_password"`
}

func (u *User) ToDetailsDto() *UserDetailsDto {
	return &UserDetailsDto{u.Username, u.FirstName, u.LastName, u.Address,
		u.AdminRights}
}

func (u *UserCreateDto) ToUser(salt string, hash string) *User {
	return &User{Base{}, u.Username, u.FirstName, u.LastName, u.Address, salt,
		hash, u.AdminRights}
}
