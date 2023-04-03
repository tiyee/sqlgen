package order_repo

import (
	"github.com/tiyee/sqlgen/models"
	"github.com/tiyee/sqlgen/orm"
)

func ShopOrder(orderId int64) (*models.ShopOrder, error) {
	return orm.New(&models.ShopOrder{}).Row("id=?", []interface{}{orderId})
}
func ShopOrderLimit(where string, values []interface{}, offset, page int64) ([]*models.ShopOrder, error) {
	return orm.New(&models.ShopOrder{}).Limit(where, values, offset, page, "id desc")
}
func ShopOrderCount(where string, values []interface{}) (int64, error) {
	return orm.New(&models.ShopOrder{}).Count(where, values)
}
func ShopOrderUpdateByPk(cls *models.ShopOrder, data map[string]interface{}, pk int64) (int64, error) {
	return orm.New(cls).UpdateByPk(data, pk)
}
func ShopOrderSave(cls *models.ShopOrder) (int64, error) {
	return orm.New(cls).Save()
}
func ShopOrderUpdate(cls *models.ShopOrder) (int64, error) {
	return orm.New(cls).Update(int64(cls.Id))
}
