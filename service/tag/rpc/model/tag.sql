CREATE TABLE `tag`
(
    `id`          bigint(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `entity_type` tinyint(4) unsigned NOT NULL DEFAULT 0 COMMENT '标签的所属实体: 1: blog 2: comment 3: user',
    `entity_id`   bigint(11) unsigned NOT NULL DEFAULT 0 COMMENT '标签所属实体ID',
    `content`     varchar(10)                  DEFAULT NULL COMMENT '标签内容',
    `count`       int(11) unsigned    NOT NULL DEFAULT 0 COMMENT '标签认同数量',
    `created_at`  timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    primary key (`id`),
    key idx_user_id_status (`entity_id`, `entity_type`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='标签表';