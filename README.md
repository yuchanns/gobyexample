# gobyexample
personal practises of go usage
## Construct of Table
```sql
CREATE TABLE `order` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `order_no` varchar(32) NOT NULL DEFAULT '',
  `user_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'user_id',
  `total_price` int(10) unsigned NOT NULL DEFAULT '0',
  `postage` int(10) unsigned NOT NULL DEFAULT '0',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '0-pending 1-payed 2-transporting 3-received 4-confirmed 5-refunding 6-refunded 7-cancel',
  `is_deleted` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `created_at` int(10) unsigned NOT NULL DEFAULT '0',
  `updated_at` int(10) unsigned NOT NULL DEFAULT '0',
  `deleted_at` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `order_no` (`order_no`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='Main order';

CREATE TABLE `order_item` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `order_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'related to order primary key',
  `s_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'shop id',
  `user_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'user id',
  `g_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'related to goods primary key',
  `name` varchar(50) NOT NULL DEFAULT '',
  `num` int(10) unsigned NOT NULL DEFAULT '0',
  `price` int(10) unsigned NOT NULL DEFAULT '0',
  `is_deleted` int(10) unsigned NOT NULL DEFAULT '0',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '0-pending 1-payed 2-transporting 3-received 4-confirmed 5-refunding 6-refunded 7-cancel',
  `created_at` int(10) unsigned NOT NULL DEFAULT '0',
  `updated_at` int(10) unsigned NOT NULL DEFAULT '0',
  `deleted_at` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `order_id` (`order_id`),
  KEY `g_id` (`g_id`),
  KEY `s_id` (`s_id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='Order items';

```
