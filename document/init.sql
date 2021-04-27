DROP DATABASE  IF EXISTS  L;
CREATE DATABASE L;
USE L;
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'UID',
  `passport` varchar(45)  NOT NULL COMMENT '账号',
  `password` char(32)  NOT NULL COMMENT 'MD5密码',
  `nickname` varchar(45)  NULL DEFAULT NULL COMMENT '昵称',
  `avatar` varchar(200)  NULL DEFAULT NULL COMMENT '头像地址',
  `status` tinyint(4) NULL DEFAULT 1 COMMENT '状态 0:启用 1:禁用',
  `gender` tinyint(1) NULL DEFAULT 0 COMMENT '性别 0: 未设置 1: 男 2: 女',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '注册时间',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_user_passport`(`passport`) USING BTREE,
  UNIQUE INDEX `uni_user_nickname`(`nickname`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 18  COMMENT = '用户基础表' ROW_FORMAT = Dynamic;
