CREATE TABLE `shop_order` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `uuid` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'uuid，为兼容旧逻辑',
  `user_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '购买用户id',
  `payment` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '0-no,1-wepay,2-alipay',
  `fee_amount` int unsigned NOT NULL DEFAULT '0' COMMENT '总共需要支付费用',
  `prepare_time` int unsigned NOT NULL DEFAULT '0' COMMENT '发起支付时间,10位时间戳',
  `paid_time` int unsigned NOT NULL DEFAULT '0' COMMENT '回调确认支付时间',
  `goods_detail` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '订单商品详情，json数组',
  `status` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '1-待支付,3-待确认,4-已经支付成功,5-支付失败,6-物流中，7-关闭,8-完成',
  `created_ts` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_ts` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `logistics_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '物流id',
  `flag` int unsigned NOT NULL DEFAULT '0' COMMENT '01-已计算sold_amount,10-已经评论',
  PRIMARY KEY (`id`),
  KEY `idx_uuid` (`uuid`)
) ENGINE=InnoDB AUTO_INCREMENT=33830785918 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;