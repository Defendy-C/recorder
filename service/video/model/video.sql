CREATE DATABASE IF NOT EXISTS recorder DEFAULT CHARACTER SET utf8mb4;

use recorder;

create table video
(
    id          INT NOT NULL AUTO_INCREMENT,
    name        VARCHAR(255) DEFAULT '' COMMENT '视频名',
    user_id     INT          DEFAULT 0 COMMENT '用户Id',
    path        VARCHAR(255) DEFAULT '' COMMENT '存放路径',
    created_at  DATETIME     DEFAULT 0 COMMENT '创建时间',
    finished_at DATETIME     DEFAULT 0 COMMENT '完成时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY (`user_id`, `name`)
)

