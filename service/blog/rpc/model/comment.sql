CREATE TABLE `comment`
(
    `id`           bigint(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `type`         tinyint(4) unsigned NOT NULL DEFAULT 0 COMMENT '评论类型：1. 对博客的评论 2. 对评论的评论 3. 对用户的评论',
    `content`      varchar(255)        not null DEFAULT '' COMMENT '评论的内容',
    `to_id`        bigint(11) unsigned NOT NULL DEFAULT 0 COMMENT '此评论所归属的id,若type是博客，此id是评论的博客',
    `from_user_id` bigint(11) unsigned NOT NULL DEFAULT 0 COMMENT '评论发起者',
    `to_user_id`   bigint(11) unsigned NOT NULL DEFAULT 0 COMMENT '评论所回应的人，若是博客则是写博客的人ID',
    `like_count`   bigint(11)          NOT NULL DEFAULT 0 COMMENT '喜欢的人数',
    `extra`        text                         DEFAULT NULL COMMENT '一些额外的json数据',
    `created_at`   timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `deleted_at`   bigint(11)                   DEFAULT 0 COMMENT '删除时间',
    primary key (`id`),
    key idx_belong_id_type_deleted (`to_id`, `type`, `deleted_at`),
    key idx_from_id (`from_user_id`),
    key idx_to_id (`to_user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='评论表';

# 常用的 sql
# 一条博客下的评论：select * from comment where belong_id = x and type = 1 and deleted_at = 0
# 一条评论下的评论: select * from comment where belong_id = x and type = 2 and deleted_at = 0
# 一个用户的发出的评论: select * from comment where from_id = x and deleted_at = 0
# 一个用户收到的评论：select * from comment where to_id = x and deleted_at = 0