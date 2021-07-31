CREATE TABLE `blog`
(
    `id`         bigint(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`    bigint(11)          NOT NULL DEFAULT 0 COMMENT '博客的作者',
    `title`      varchar(100)        NOT NULL DEFAULT '' COMMENT '博客标题',
    `is_folder`  tinyint(1)          NOT NULL DEFAULT 0 COMMENT '0: 正常博客，1：博客类别/文件夹',
    `content`    text                         DEFAULT NULL COMMENT '博客的内容',
    `status`     tinyint(4) unsigned NOT NULL DEFAULT 1 COMMENT '博客状态 1:所有人可见 2. 仅自己可见 3. 删除',
    `folder_id`  bigint(11) unsigned NOT NULL DEFAULT 0 COMMENT '博客所属类别/文件夹，0 则无类别',
    `extra`      text                         DEFAULT NULL COMMENT '一些额外的json数据',
    `like_count` bigint(11)          NOT NULL DEFAULT 0 COMMENT '喜欢的人数',
    `created_at` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    primary key (`id`),
    key idx_user_id_status (`user_id`, `status`),
    unique key uniq_created_at (`created_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='博客表';

# 可用常用到的一些sql
# 查看所有博客: select * from blog where is_folder = false order by created_at limit 10
# 查询用户所有博客: select * from blog where user_id = x and status = 1 and is_category = 0
# 查询用户所有一级博客与文件夹：select * from where user_id = x and category_id = 0 and status = 1
# 查询用户某个种类（目录下）a 的博客：select * from blog where user_id = x and category = a and status = 1