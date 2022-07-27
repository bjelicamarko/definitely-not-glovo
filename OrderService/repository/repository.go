package repository

import (
	"OrderService/models"
	"errors"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		if page < 0 {
			page = 0
		}

		pageSize, _ := strconv.Atoi(r.URL.Query().Get("size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := page * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (repo *Repository) SearchOrders(r *http.Request) ([]models.OrderDTO, int64, error) {
	return nil, 0, nil
}

func (repo *Repository) SearchOrdersByRestaurantAndStatus(r *http.Request) ([]models.OrderDTO, int64, error) {
	var ordersDTO []models.OrderDTO
	var orders []*models.Order
	var totalElements int64

	restaurantId := r.URL.Query().Get("restaurantId")
	orderStatus := r.URL.Query().Get("orderStatus")

	result := repo.db.Scopes(Paginate(r)).Table("orders").
		Where("(deleted_at IS NULL and id_restaurant = ?) and "+
			"('' = ? or order_status = ?)",
			restaurantId, orderStatus, orderStatus).
		Find(&orders)

	repo.db.Table("orders").
		Where("(deleted_at IS NULL and id_restaurant = ?) and "+
			"('' = ? or order_status = ?)",
			restaurantId, orderStatus, orderStatus).
		Count(&totalElements)

	if result.Error != nil {
		return nil, totalElements, result.Error
	}

	for _, order := range orders {
		ordersDTO = append(ordersDTO, order.ToOrderDTO())
	}

	return ordersDTO, totalElements, nil
}

func (repo *Repository) FindOrderById(id uint) (*models.OrderDTO, error) {
	var order models.Order
	result := repo.db.Table("orders").Where("id = ?", id).First(&order)

	if result.Error != nil {
		return nil, errors.New("order cannot be found")
	}

	var orderDTO models.OrderDTO = order.ToOrderDTO()

	var orderItems []models.OrderItem
	result2 := repo.db.Table("order_items").Where("id_order = ?", id).Find(&orderItems)

	if result2.Error != nil {
		return nil, result.Error
	}

	for _, orderItem := range orderItems {
		orderDTO.OrderItemsDTO = append(orderDTO.OrderItemsDTO, orderItem.ToOrderItemDTO())
	}

	return &orderDTO, nil
}

func (repo *Repository) CreateOrder(orderDTO *models.OrderDTO) (*models.OrderDTO, error) {
	var order models.Order = orderDTO.ToOrder()
	result := repo.db.Table("orders").Create(&order)

	if result.Error != nil {
		return nil, errors.New("error while creating order")
	}

	var retValue models.OrderDTO = order.ToOrderDTO()

	for _, orderItemDTO := range orderDTO.OrderItemsDTO {
		orderItemDTO.IdOrder = order.ID
		var orderItem = orderItemDTO.ToOrderItem()
		repo.db.Table("order_items").Create(&orderItem)
		retValue.OrderItemsDTO = append(retValue.OrderItemsDTO, orderItem.ToOrderItemDTO())
	}

	return &retValue, nil
}

func (repo *Repository) ChangeStatusOfOrder(orderStatusDTO *models.OrderStatusDTO) (*models.OrderDTO, error) {
	var order models.Order
	result := repo.db.Table("orders").Where("id = ?", orderStatusDTO.IdOrder).First(&order)

	if result.Error != nil {
		return nil, errors.New("order cannot be found")
	}

	order.OrderStatus = models.OrderStatus(orderStatusDTO.OrderStatus)

	result2 := repo.db.Table("orders").Save(&order)

	if result2.Error != nil {
		return nil, errors.New("error while changing status of order")
	}

	var retValue models.OrderDTO = order.ToOrderDTO()
	return &retValue, nil
}
