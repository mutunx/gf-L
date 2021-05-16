DROP DATABASE  IF EXISTS  L;
CREATE DATABASE L;
USE L;
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
    id         int unsigned auto_increment comment 'UID'
        primary key,
    passport   varchar(45)          not null comment '账号',
    password   char(32)             not null comment 'MD5密码',
    nickname   varchar(45)          null comment '昵称',
    avatar     varchar(200)         null comment '头像地址',
    status     tinyint    default 1 null comment '状态 0:启用 1:禁用',
    gender     tinyint(1) default 0 null comment '性别 0: 未设置 1: 男 2: 女',
    created_at datetime             null comment '注册时间',
    updated_at datetime             null comment '更新时间',
    id_card    varchar(20)          null comment '身份证号码',
    birthday   varchar(20)          null comment '生日',
    address    varchar(50)          null comment '具体地址',
    province   varchar(50)          null comment '省',
    city       varchar(50)          null comment '市',
    area       varchar(50)          null comment '地区',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_user_passport`(`passport`) USING BTREE,
  UNIQUE INDEX `uni_user_nickname`(`nickname`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 18  COMMENT = '用户基础表' ROW_FORMAT = Dynamic;

DROP TABLE IF EXISTS `l.province_city_code`;
create table  l.province_city_code
(
    code          varchar(100) null comment '六位地区码',
    name          varchar(10)  null comment '六位地区码含义',
    province_code varchar(100) null comment '前两位省市码',
    province_name varchar(10)  null comment '前两位省市码含义',
    city_code     varchar(100) null comment '前四位市县码',
    city_name     varchar(10)  null comment '前四位市县码含义',
    area_code     varchar(100) null comment '前六位地区码',
    area_name     varchar(10)  null comment '前六位地区码含义'
)
    comment '身份证前六位含义';

