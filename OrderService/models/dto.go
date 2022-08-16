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
	Id             uint           `json:"Id"`
	RestaurantName string         `json:"RestaurantName"`
	IdRestaurant   uint           `json:"IdRestaurant"`
	IdAppUser      uint           `json:"IdAppUser"`
	IdEmployee     uint           `json:"IdEmployee"`
	IdDeliverer    uint           `json:"IdDeliverer"`
	OrderStatus    string         `json:"OrderStatus"`
	TotalPrice     float32        `json:"TotalPrice"`
	Tip            float32        `json:"Tip"`
	Note           string         `json:"Note"`
	DateTime       string         `json:"DateTime"`
	Country        string         `json:"Country"`
	City           string         `json:"City"`
	Street         string         `json:"Street"`
	StreetNumber   string         `json:"StreetNumber"`
	Ptt            uint           `json:"Ptt"`
	DisplayName    string         `json:"DisplayName"`
	Longitude      float32        `json:"Longitude"`
	Latitude       float32        `json:"Latitude"`
	OrderItemsDTO  []OrderItemDTO `json:"OrderItemsDTO"`
	Reviewed       bool           `json:"Reviewed"`
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

type OrderForReportDTO struct {
	IdOrder                uint                    `json:"id_order"`
	RestaurantName         string                  `json:"restaurant_name"`
	IdRestaurant           uint                    `json:"id_restaurant"`
	IdAppUser              uint                    `json:"id_app_user"`
	IdEmployee             uint                    `json:"id_employee"`
	IdDeliverer            uint                    `json:"id_deliverer"`
	OrderStatus            string                  `json:"order_status"`
	TotalPrice             float32                 `json:"total_price"`
	Tip                    float32                 `json:"tip"`
	DateTime               string                  `json:"date_time"`
	OrderItemsForReportDTO []OrderItemForReportDTO `json:"order_items_for_report_dto"`
}

type OrderItemForReportDTO struct {
	Id           uint    `json:"id"`
	IdOrder      uint    `json:"id_order"`
	IdArticle    uint    `json:"id_article"`
	ArticleName  string  `json:"article_name"`
	CurrentPrice float32 `json:"current_price"`
	Quantity     uint    `json:"quantity"`
	TotalPrice   float32 `json:"total_price"`
}
