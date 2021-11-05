CREATE DATABASE IF NOT EXISTS recorder DEFAULT CHARACTER SET utf8mb4;

use recorder;

CREATE TABLE if NOT EXISTS `users`
(
    `id` INT NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(255) DEFAULT '' NOT NULL COMMENT '用户名',
    `password` VARCHAR(255) DEFAULT '' NOT NULL COMMENT '密码',
    PRIMARY KEY (`id`),
    UNIQUE KEY `username_unique_index` (`username`)
) ENGINE=InnoDB;