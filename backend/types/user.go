package types

import "strconv"

type User struct {
	Base
	FirstName   string
	LastName    string
	Address     string
	Salt        string
	Hash        string
	AdminRights bool
}

type UserLoginDto struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

type UserCreateDto struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Address     string `json:"address"`
	Password    string `json:"password"`
	AdminRights bool   `json:"adminrights"`
}

type UserDetailsDto struct {
	Id          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Address     string `json:"address"`
	AdminRights bool   `json:"adminrights"`
}

func (u *User) ToDetailsDtop() *UserDetailsDto {
	return &UserDetailsDto{strconv.FormatUint(u.ID, 10), u.FirstName,
		u.LastName, u.Address, u.AdminRights}
}
