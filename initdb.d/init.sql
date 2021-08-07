CREATE TABLE IF NOT EXISTS `go_database`.`users` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(40) NOT NULL COMMENT 'ユーザーの名前',
  `token` VARCHAR(255) NOT NULL COMMENT '認証トークン',
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`)
);