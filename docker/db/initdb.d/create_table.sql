USE `test`;

CREATE TABLE  IF NOT EXISTS `test`.`todo` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT "" COMMENT 'タイトル',
  `content` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT "" COMMENT '内容',
  `deadline` timestamp NOT NULL COMMENT '期限',
  `state` int(11) NOT NULL DEFAULT '0' COMMENT '状態',
  `created_at` timestamp NULL DEFAULT now() COMMENT '作成日時',
  `updated_at` timestamp NULL DEFAULT now() ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT 'TODO管理テーブル';