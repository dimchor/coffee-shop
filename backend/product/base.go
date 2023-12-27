package product

type ProductBase struct {
	Id    uint64  `json:"id" gorm:"primaryKey;not null;autoIncrement"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
