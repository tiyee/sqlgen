package v5

import (
	"fmt"
)

func Parse() {
	s := "CREATE TABLE `cat_mood_user_like_relation` (\n  `id` bigint unsigned NOT NULL AUTO_INCREMENT,\n  `mood_id` bigint unsigned NOT NULL,\n  `user_id` bigint unsigned NOT NULL,\n  `cat_id` bigint unsigned NOT NULL,\n  `created_ts` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,\n  `updated_ts` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,\n  `arti_like` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '用户点赞',\n  `auto_like` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '自动点赞',\n  PRIMARY KEY (`id`),\n  UNIQUE KEY `uk_userId_moodId` (`user_id`,`mood_id`),\n  KEY `idx_userId_moodId` (`user_id`,`mood_id`),\n  KEY `idx_mood_id` (`mood_id`),\n  KEY `idx_cat_id` (`cat_id`)\n) ENGINE=InnoDB AUTO_INCREMENT=4990 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	bs := []rune(s)
	l := lex{
		input: bs,
		size:  len(bs),
		pos:   0,
	}
	yyParse(&l)
	fmt.Println(l.Stmt.String())
	fmt.Println("-------------")
	for _, x := range *l.Stmt.TableOptsStmt {
		fmt.Println(x.pos, x.s)
	}
	fmt.Printf("%+v\n", l.Stmt.TableOptsStmt)

}
func ParseSql(sql string) *Creator {
	bs := []rune(sql)
	l := lex{
		input: bs,
		size:  len(bs),
		pos:   0,
	}
	yyParse(&l)
	c := Creator(l.Stmt)
	return &c
}
