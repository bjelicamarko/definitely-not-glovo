package models

type OrderItemDTO struct {
	Id           uint    `json:"Id"`
	IdOrder      uint    `json:"IdOrder"`
	IdArticle    uint    `json:"IdArticle"`
	ArticleName  string  `json:"ArticleName"`
	CurrentPrice float32 `json:"CurrentPrice"`
	Quantity     uint    `json:"Quantity"`
	TotalPrice   float32 `json:"TotalPrice"`
}

type OrderDTO struct {
	Id            uint           `json:"Id"`
	IdRestaurant  uint           `json:"IdRestaurant"`
	IdAppUser     uint           `json:"IdAppUser"`
	IdEmployee    uint           `json:"IdEmployee"`
	IdDeliverer   uint           `json:"IdDeliverer"`
	OrderStatus   string         `json:"OrderStatus"`
	TotalPrice    float32        `json:"TotalPrice"`
	Tip           float32        `json:"Tip"`
	Note          string         `json:"Note"`
	DateTime      string         `json:"DateTime"`
	OrderItemsDTO []OrderItemDTO `json:"OrderItemsDTO"`
}

type OrderDTOMessage struct {
	OrderDTO OrderDTO `json:"OrderDTO"`
	Message  string   `json:"Message"`
}

type OrdersPageable struct {
	Elements      []OrderDTO `json:"Elements"`
	TotalElements int64      `json:"TotalElements"`
}

type OrderStatusDTO struct {
	IdOrder     uint   `json:"IdOrder"`
	OrderStatus string `json:"OrderStatus"`
	IdEmployee  uint   `json:"IdEmployee"`
	IdDeliverer uint   `json:"IdDeliverer"`
}
