package orders

import (
	"github.com/ariesekoprasetyo/hacktiv8_7/datatransfers"
	"github.com/ariesekoprasetyo/hacktiv8_7/db"
)

func CreateOrder(req datatransfers.Orders) error {
	postOrder := db.Orders{
		CustomerName: req.CustomerName,
		Ordered_at:   req.OrderedAt,
		Items:        []db.Items{},
	}
	for _, itemReq := range req.Items {
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

func UpdateOrder(orderID uint, req datatransfers.OrdersUpdate) (db.Orders, error) {
	// Update data orders
	order := db.Orders{
		CustomerName: req.CustomerName,
	}

	if err := db.DB.Model(&order).Where("order_id = ?", orderID).Updates(order).Error; err != nil {
		return db.Orders{}, err
	}

	for _, item := range req.Items {
		itemData := db.Items{
			ItemCode:    item.ItemCode,
			Description: item.Description,
			Quantity:    item.Quantity,
		}

		if err := db.DB.Model(&itemData).Where("order_id = ?", orderID).Updates(itemData).Error; err != nil {
			return db.Orders{}, nil
		}
	}
	result, err := GetDataById(orderID)
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
