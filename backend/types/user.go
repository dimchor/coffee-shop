package types

import "strconv"

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
	AdminRights bool   `json:"adminrights"`
}

type UserDetailsDto struct {
	Username    string `json:"username"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Address     string `json:"address"`
	AdminRights bool   `json:"adminrights"`
}

type UserChangePasswordDto struct {
	Username    string `json:"username"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func (u *User) ToDetailsDto() *UserDetailsDto {
	return &UserDetailsDto{strconv.FormatUint(u.ID, 10), u.FirstName,
		u.LastName, u.Address, u.AdminRights}
}
