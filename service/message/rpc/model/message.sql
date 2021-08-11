CREATE TABLE `message`
(
    `id`           bigint(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `from_id`      bigint(11) unsigned NOT NULL DEFAULT 0 COMMENT '发送者 ID',
    `to_id`        bigint(11) unsigned NOT NULL DEFAULT 0 COMMENT '接受者 ID',
    `session_id`   varchar(50)         NOT NULL DEFAULT '' COMMENT '会话 ID， 由发送者接受者ID组合，小 ID：大 ID',
    `message_type` tinyint(4) unsigned NOT NULL DEFAULT 0 COMMENT '消息类型 1: 普通私信， 2：群聊',
    `content`      text COMMENT '私信内容',
    `status`       tinyint(4) unsigned NOT NULL DEFAULT 0 COMMENT '1：未读 2：已读 3：撤回',
    `extra`        text                         DEFAULT NULL COMMENT '一些额外的json数据',
    `created_at`   timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`   timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    primary key (`id`),
    key idx_to_id_status (`to_id`, `status`),
    key idx_from_to (`from_id`, `to_id`),
    Key idx_session_id (`session_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='私信表';