package order

import "fmt"

type service struct {
	OrderRepository Repository
}

func NewService(repository Repository) Service {
	return &service{
		OrderRepository: repository,
	}
}

func (s *service) CreateOrder(req *CreateOrderRequest) (*CreateOrderResponse, error) {
	var items []Item

	for _, item := range req.Item {
		items = append(items, Item{
			Code:        item.Code,
			Description: item.Description,
			Quantity:    item.Quantity,
		})
	}

	orderReq := &Order{
		OrderedAt:    req.OrderedAt,
		CustomerName: req.CustomerName,
		Item:         items,
	}

	res, err := s.OrderRepository.Create(orderReq)
	if err != nil {
		return nil, err
	}
	response := &CreateOrderResponse{
		Id:           res.Id,
		OrderedAt:    res.OrderedAt,
		CustomerName: res.CustomerName,
		Item:         res.Item,
	}

	return response, nil
}

func (s *service) GetOrder() (*[]GetOrderResponse, error) {
	var response []GetOrderResponse

	results, err := s.OrderRepository.Get()
	if err != nil {
		return nil, err
	}

	if len(*results) == 0 {
		return nil, fmt.Errorf("empty")
	}

	for _, result := range *results {
		response = append(response, GetOrderResponse{
			Id:           result.Id,
			OrderedAt:    result.OrderedAt,
			CustomerName: result.CustomerName,
			Item:         result.Item,
		})
	}

	return &response, nil
}

func (s *service) DeleteOrder(req *DeleteOrderRequest) error {
	err := s.OrderRepository.Delete(req.Id)
	if err != nil {
		if err.Error() == "null" {
			return err
		}
		return fmt.Errorf("not found")
	}

	return err
}

func (s *service) UpdateOrder(orderId uint, req *UpdateOrderRequest) (*UpdateOrderResponse, error) {

	order := &Order{
		Id:           orderId,
		CustomerName: req.CustomerName,
		OrderedAt:    req.OrderedAt,
		Item:         req.Item,
	}

	res, err := s.OrderRepository.Update(order)
	if err != nil {
		return nil, err
	}

	result := &UpdateOrderResponse{
		Id:           res.Id,
		CustomerName: res.CustomerName,
		OrderedAt:    res.OrderedAt,
		Item:         res.Item,
	}

	return result, nil
}
