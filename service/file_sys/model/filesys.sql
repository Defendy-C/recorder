CREATE DATABASE IF NOT EXISTS recorder DEFAULT CHARACTER SET utf8mb4;

use recorder;

create table file_sys
(
    id          INT NOT NULL AUTO_INCREMENT,
    path        VARCHAR(255) DEFAULT '' COMMENT '存放路径',
    created_at  DATETIME     DEFAULT 0 COMMENT '创建时间',
    finished_at DATETIME     DEFAULT 0 COMMENT '完成时间',
    total_chunk INT          DEFAULT 0 COMMENT '文件分段数',
    PRIMARY KEY (`id`),
    UNIQUE KEY (`path`)
)
