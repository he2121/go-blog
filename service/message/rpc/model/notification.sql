CREATE TABLE `notification`
(
    `id`          bigint(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`     bigint(11) unsigned NOT NULL DEFAULT 0 COMMENT '通知用户 ID',
    `event_type`  tinyint(4) unsigned NOT NULL DEFAULT 0 COMMENT '触发该通知的事件类型：1：点赞 2: 评论 3: 发帖 4：关注',
    `entity_type` tinyint(4) unsigned NOT NULL DEFAULT 0 COMMENT '触发事件的实体类型',
    `entity_id`   bigint(11) unsigned NOT NULL DEFAULT 0 COMMENT '触发事件的实体 ID',
    `trigger_id`  bigint(11) unsigned NOT NULL DEFAULT 0 COMMENT '触发事件的用户 ID',
    `status`      tinyint(4) unsigned NOT NULL DEFAULT 0 COMMENT '1：未读 2：已读',
    `extra`       text                         DEFAULT NULL COMMENT '一些额外的json数据',
    `created_at`  timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    primary key (`id`),
    key idx_user_id_status (`user_id`, `status`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='系统通知';