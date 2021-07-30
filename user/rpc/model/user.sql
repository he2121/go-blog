CREATE TABLE `user`
(
    `id`         bigint(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `email`      varchar(50) NOT NULL COMMENT '注册邮箱',
    `phone`      varchar(50) NOT NULL DEFAULT '' COMMENT '用户手机',
    `name`       varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
    `password`   char(64)    NOT NULL DEFAULT '' COMMENT '加 salt Hash密码',
    `status`     tinyint(4) unsigned NOT NULL DEFAULT 1 COMMENT '用户状态 1:正常，2:停用',
    `gender`     tinyint(1) unsigned NOT NULL DEFAULT 1 COMMENT '性别 1:男 2:女',
    `birth_date`  timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '出生日期',
    `extra`      text                 DEFAULT NULL COMMENT '一些额外的json数据',
    `created_at` timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    primary key (`id`),
    unique key `uniq_email` (`email`),
    key          `idx_phone` (`phone`)
)ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT ='用户表';