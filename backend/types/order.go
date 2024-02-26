package types

import "errors"

type Order struct {
	Base
	UserID    uint64
	User      User
	ProductID uint64 `gorm:"primaryKey; not null"`
	Product   Product
	Count     uint32
	SessionID uint64
	Session   Session
}

type OrderCreateDto struct {
	UserID  uint64 `json:"user_id"`
	Product []struct {
		ID    uint64 `json:"product_id"`
		Count uint32 `json:"count"`
	}
}

type OrderInfoDto struct {
	OrderId uint64 `json:"order_id"`
	UserID  uint64 `json:"user_id"`
	Product []struct {
		ID    uint64 `json:"product_id"`
		Count uint32 `json:"count"`
	}
}

func (o *OrderCreateDto) ToOrders(OrderId uint64) []Order {
	size := len(o.Product)
	orders := make([]Order, size)

	for i := 0; i < size; i++ {
		orders = append(orders,
			Order{
				Base:      Base{ID: OrderId},
				UserID:    o.UserID,
				ProductID: o.Product[i].ID,
				Count:     o.Product[i].Count,
			})
	}

	return orders
}

func ToOrderInfo(orders []Order) (*OrderInfoDto, error) {
	size := len(orders)
	if size < 1 {
		return nil, errors.New("no orders")
	}

	order_info := OrderInfoDto{OrderId: orders[0].ID, UserID: orders[0].UserID}

	products := make([]struct {
		ID    uint64
		Count uint32
	}, size)

	for i := 0; i < size; i++ {
		products = append(products,
			struct {
				ID    uint64
				Count uint32
			}{
				ID:    orders[i].ProductID,
				Count: orders[i].Count,
			})
	}

	order_info.Product = []struct {
		ID    uint64 "json:\"product_id\""
		Count uint32 "json:\"count\""
	}(products)

	return &order_info, nil
}
