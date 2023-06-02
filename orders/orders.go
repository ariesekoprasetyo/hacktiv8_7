package orders

import (
	"github.com/ariesekoprasetyo/hacktiv8_7/db"
	"log"
	"time"
)

type OrdersReq struct {
	CustomerName string     `json:"customer_name" binding:"required"`
	Items        []ItemsReq `json:"items" binding:"required"`
}

type ItemsReq struct {
	ItemCode    string `json:"item_code" binding:"required"`
	Description string `json:"description" binding:"required"`
	Quantity    int    `json:"quantity" binding:"required"`
}

type UpdateReq struct {
	CustomerName string     `json:"customer_name" binding:"required"`
	Items        []ItemsReq `json:"items" binding:"required"`
}

func CreateOrder(body OrdersReq) error {
	postOrder := db.Orders{
		CustomerName: body.CustomerName,
		Ordered_at:   time.Now(),
	}
	for _, itemReq := range body.Items {
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

func UpdateOrder(orderID uint, bodyUpdate UpdateReq) (db.Orders, error) {
	order := db.Orders{
		CustomerName: bodyUpdate.CustomerName,
	}
	cekById := db.Orders{OrderId: orderID}
	cekID := db.DB.Preload("Items").Take(&cekById)
	if cekID.Error != nil {
		return db.Orders{}, cekID.Error
	}

	if err := db.DB.Model(&order).Where("order_id = ?", orderID).Updates(order).Error; err != nil {
		return db.Orders{}, err
	}

	var items []db.Items
	if err := db.DB.Find(&items, "order_id = ?", orderID).Error; err != nil {
		return db.Orders{}, err
	}
	for i, bodyRequestItems := range bodyUpdate.Items {
		if err := db.DB.Model(&db.Items{}).Where("item_id = ?", items[i].ItemId).Updates(bodyRequestItems).Error; err != nil {
			log.Println(err)
			return db.Orders{}, err
		}
	}
	result, err := GetDataById(orderID)
	return result, err
}

func DeleteOrder(orderID uint) (err error) {

	_, err = GetDataById(orderID)
	if err != nil {
		return err
	}
	if err := db.DB.Model(&db.Items{}).Where("order_id = ?", orderID).Delete(db.Items{}).Error; err != nil {
		return err
	}
	if err := db.DB.Model(&db.Orders{}).Where("order_id = ?", orderID).Delete(db.Orders{}).Error; err != nil {
		return err
	}
	return nil
}
