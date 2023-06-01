package orders

import (
	"github.com/ariesekoprasetyo/hacktiv8_7/db"
	"time"
)

type Orders struct {
	CustomerName string  `json:"customer_name" binding:"required"`
	Items        []Items `json:"items" binding:"required"`
}

type Items struct {
	ItemCode    string `json:"item_code" binding:"required"`
	Description string `json:"description" binding:"required"`
	Quantity    int    `json:"quantity" binding:"required"`
}

type Update struct {
	CustomerName string  `json:"customer_name" binding:"required"`
	Items        []Items `json:"items" binding:"required"`
}

func CreateOrder(orderReq Orders) error {
	postOrder := db.Orders{
		CustomerName: orderReq.CustomerName,
		Ordered_at:   time.Now(),
	}
	for _, itemReq := range orderReq.Items {
		item := db.Items{
			ItemCode:    itemReq.ItemCode,
			Description: itemReq.Description,
			Quantity:    itemReq.Quantity,
		}
		postOrder.Items = append(postOrder.Items, item)
	}
	if result := db.DB.Create(&postOrder); result.Error != nil {
		return result.Error
	}
	return nil
}

func GetAllData() ([]db.Orders, error) {
	var orders []db.Orders
	result := db.DB.Preload("Items").Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}

func GetDataById(id uint) (db.Orders, error) {
	orderById := db.Orders{OrderId: id}
	result := db.DB.Preload("Items").Take(&orderById)
	if result.Error != nil {
		return db.Orders{}, result.Error
	}
	return orderById, nil
}

func UpdateOrder(orderID uint, orderReq Update) (db.Orders, error) {
	order := db.Orders{
		CustomerName: orderReq.CustomerName,
	}
	result, err := GetDataById(orderID)
	if err := db.DB.Debug().Model(&order).Where("order_id = ?", orderID).Updates(order).Error; err != nil {
		return db.Orders{}, err
	}

	for _, item := range orderReq.Items {
		itemData := db.Items{
			ItemCode:    item.ItemCode,
			Description: item.Description,
			Quantity:    item.Quantity,
		}

		if err := db.DB.Model(&itemData).Where("order_id = ?", orderID).Updates(itemData).Error; err != nil {
			return db.Orders{}, nil
		}
	}
	return result, err
}

func DeleteOrder(orderID uint) error {

	if err := db.DB.Model(&db.Items{}).Where("order_id = ?", orderID).Delete(db.Items{}).Error; err != nil {
		return err
	}
	if err := db.DB.Model(&db.Orders{}).Where("order_id = ?", orderID).Delete(db.Orders{}).Error; err != nil {
		return err
	}
	return nil
}
