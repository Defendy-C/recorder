CREATE DATABASE IF NOT EXISTS recorder DEFAULT CHARACTER SET utf8mb4;

use recorder;

create table video
(
    id          INT NOT NULL AUTO_INCREMENT,
    title       VARCHAR(255) DEFAULT '' COMMENT '视频名',
    user_id     INT          DEFAULT 0 COMMENT '用户Id',
    file_id     INT          DEFAULT 0 COMMENT '文件Id',
    created_at  DATETIME     DEFAULT 0 COMMENT '创建时间',
    description TEXT COMMENT '描述',
    PRIMARY KEY (`id`),
    UNIQUE KEY (`user_id`, `created_at`, `title`)
)