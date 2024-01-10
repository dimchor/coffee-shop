package types

type Order struct {
	Base
	UserID    uint64
	User      User
	ProductID uint64
	Products  []Product
}
