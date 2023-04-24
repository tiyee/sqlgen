package v4

import (
	"fmt"
)

func Parse() {
	s := "CREATE TABLE `activity` (\n  `id` bigint unsigned NOT NULL AUTO_INCREMENT,\n  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,\n  `code` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,\n  `begin_date` timestamp NULL DEFAULT NULL,\n  `end_date` timestamp NULL DEFAULT NULL,\n  `required_num` bigint unsigned DEFAULT NULL,\n  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,\n  `created_ts` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,\n  `updated_ts` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,\n  `activity_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '',\n  `bonus` int DEFAULT NULL,\n  PRIMARY KEY (`id`),\n  UNIQUE KEY `uk_code` (`code`)\n) ENGINE=InnoDB AUTO_INCREMENT=42 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	bs := []rune(s)
	l := lex{
		input: bs,
		size:  len(bs),
		pos:   0,
	}
	n := yyParse(&l)
	fmt.Println(n)
	fmt.Printf("%+v\n", l.Stmt)

}
