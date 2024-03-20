package order

import (
	"time"
)

type Order struct {
	Id           uint      `gorm:"type:bigint;primary_key;" json:"id"`
	CustomerName string    `gorm:"type:varchar(50);" json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Item         []Item    `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;" json:"items"`
}

type Item struct {
	Id          uint   `gorm:"type:bigint;primary_key;" json:"-"`
	Code        string `gorm:"varchar(10)" json:"itemCode"`
	Description string `gorm:"varchar(50)" json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderID     uint   `gorm:"type:bigint;" json:"order_id"`
}

type CreateOrderRequest struct {
	OrderedAt    time.Time `json:"orderedAt"`
	CustomerName string    `json:"customerName"`
	Item         []Item    `json:"items"`
}

type CreateOrderResponse struct {
	Id           uint      `json:"id"`
	OrderedAt    time.Time `json:"orderedAt"`
	CustomerName string    `json:"customerName"`
	Item         []Item    `json:"items"`
}

type GetOrderResponse struct {
	Id           uint      `json:"id"`
	OrderedAt    time.Time `json:"orderedAt"`
	CustomerName string    `json:"customerName"`
	Item         []Item    `json:"items"`
}

type UpdateOrderRequest struct {
	OrderedAt    time.Time `json:"orderedAt"`
	CustomerName string    `json:"customerName"`
	Item         []Item    `json:"items"`
}

type UpdateOrderResponse struct {
	Id           uint      `json:"id"`
	OrderedAt    time.Time `json:"orderedAt"`
	CustomerName string    `json:"customerName"`
	Item         []Item    `json:"items"`
}

type DeleteOrderRequest struct {
	Id uint `json:"id"`
}

type Repository interface {
	Create(order *Order) (*Order, error)
	Get() (*[]Order, error)
	Update(order *Order) (*Order, error)
	Delete(orderId uint) error
}

type Service interface {
	CreateOrder(req *CreateOrderRequest) (*CreateOrderResponse, error)
	GetOrder() (*[]GetOrderResponse, error)
	UpdateOrder(orderId uint, req *UpdateOrderRequest) (*UpdateOrderResponse, error)
	DeleteOrder(req *DeleteOrderRequest) error
}
