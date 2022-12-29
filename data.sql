CREATE TABLE `cron` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `uniue_code` char(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '唯一编号',
  `cron_type` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '1 脚本 2 请求链接',
  `cron_time` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'linux定时串',
  `command` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '脚本或请求链接',
  `run_status` tinyint(1) unsigned DEFAULT '0' COMMENT '0 停止 1 运行',
  `status` tinyint(1) unsigned DEFAULT '1' COMMENT '0 禁用 1 启用',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `last_runtime` int(11) DEFAULT '0' COMMENT '最近一次执行时间',
  `next_runtime` int(11) DEFAULT '0' COMMENT '下次执行时间',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci



CREATE TABLE `admin` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) DEFAULT '',
  `email` varchar(45) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '邮箱',
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '密码',
  `email_verify` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '1-已绑定 0-未绑定',
  `status` tinyint(11) NOT NULL DEFAULT '1' COMMENT '1-正常 0-禁用 -1 删除',
  `rank` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '账号等级 1超管 2业务方',
  `group_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '所属用户组',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email_UNIQUE` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
