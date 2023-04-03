// Code generated by log-gen. DO NOT EDIT.
package models

import "time"

type ShopOrder struct {
	Id          uint64    `json:"id"`// bigint(20) UNSIGNED2
	Uuid        string    `json:"uuid"`// varchar(50) CHARACTER SET utf8mb43
	UserId      uint64    `json:"user_id"`// bigint(20) UNSIGNED3
	Payment     uint8     `json:"payment"`// tinyint(4) UNSIGNED3
	FeeAmount   uint      `json:"fee_amount"`// int(11) UNSIGNED3
	PrepareTime uint      `json:"prepare_time"`// int(11) UNSIGNED3
	PaidTime    uint      `json:"paid_time"`// int(11) UNSIGNED3
	GoodsDetail string    `json:"goods_detail"`// text CHARACTER SET utf8mb43
	Status      uint8     `json:"status"`// tinyint(4) UNSIGNED3
	CreatedTs   time.Time `json:"created_ts"`// timestamp2
	UpdatedTs   time.Time `json:"updated_ts"`// timestamp3
	LogisticsId uint64    `json:"logistics_id"`// bigint(20) UNSIGNED3
	Flag        uint      `json:"flag"`// int(11) UNSIGNED3
}

func (s *ShopOrder) Pk() string {
	return "id"
}
func (s *ShopOrder) TableName() string {
	return "shop_order"
}
