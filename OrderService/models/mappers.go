package models

import (
	"gorm.io/gorm"
)

func (order *Order) ToOrderDTO() OrderDTO {
	return OrderDTO{
		Id:           order.ID,
		IdRestaurant: order.IdRestaurant,
		IdAppUser:    order.IdAppUser,
		IdEmployee:   order.IdEmployee,
		IdDeliverer:  order.IdDeliverer,
		OrderStatus:  string(order.OrderStatus),
		TotalPrice:   order.TotalPrice,
		Tip:          order.Tip,
		Note:         order.Note,
		DateTime:     order.DateTime,
		Country:      order.Country,
		City:         order.City,
		Street:       order.Street,
		StreetNumber: order.StreetNumber,
		Ptt:          order.Ptt,
		DisplayName:  order.DisplayName,
		Longitude:    order.Longitude,
		Latitude:     order.Latitude,
	}
}

func (orderItem *OrderItem) ToOrderItemDTO() OrderItemDTO {
	return OrderItemDTO{
		Id:           orderItem.ID,
		IdOrder:      orderItem.IdOrder,
		IdArticle:    orderItem.IdArticle,
		ArticleName:  orderItem.ArticleName,
		CurrentPrice: orderItem.CurrentPrice,
		Quantity:     orderItem.Quantity,
		TotalPrice:   orderItem.TotalPrice,
	}
}

func (orderDTO *OrderDTO) ToOrder() Order {
	return Order{
		Model:        gorm.Model{},
		IdRestaurant: orderDTO.IdRestaurant,
		IdAppUser:    orderDTO.IdAppUser,
		OrderStatus:  OrderStatus(orderDTO.OrderStatus),
		TotalPrice:   orderDTO.TotalPrice,
		Tip:          orderDTO.Tip,
		Note:         orderDTO.Note,
		DateTime:     orderDTO.DateTime,
		Country:      orderDTO.Country,
		City:         orderDTO.City,
		Street:       orderDTO.Street,
		StreetNumber: orderDTO.StreetNumber,
		Ptt:          orderDTO.Ptt,
		DisplayName:  orderDTO.DisplayName,
		Longitude:    orderDTO.Longitude,
		Latitude:     orderDTO.Latitude,
	}
}

func (orderItemDTO *OrderItemDTO) ToOrderItem() OrderItem {
	return OrderItem{
		Model:        gorm.Model{},
		IdOrder:      orderItemDTO.IdOrder,
		IdArticle:    orderItemDTO.IdArticle,
		ArticleName:  orderItemDTO.ArticleName,
		CurrentPrice: orderItemDTO.CurrentPrice,
		Quantity:     orderItemDTO.Quantity,
		TotalPrice:   orderItemDTO.TotalPrice,
	}
}
