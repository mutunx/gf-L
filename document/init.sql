DROP DATABASE  IF EXISTS  L;
CREATE DATABASE L;
USE L;
DROP TABLE IF EXISTS `USER`;
CREATE TABLE `USER`  (
    ID         INT UNSIGNED AUTO_INCREMENT COMMENT 'UID'
        PRIMARY KEY,
    PASSPORT   VARCHAR(45)          NOT NULL COMMENT '账号',
    PASSWORD   CHAR(32)             NOT NULL COMMENT 'MD5密码',
    NICKNAME   VARCHAR(45)          NULL COMMENT '昵称',
    AVATAR     VARCHAR(200)         NULL COMMENT '头像地址',
    STATUS     TINYINT    DEFAULT 1 NULL COMMENT '状态 0:启用 1:禁用',
    GENDER     TINYINT(1) DEFAULT 0 NULL COMMENT '性别 0: 未设置 1: 男 2: 女',
    CREATED_AT DATETIME             NULL COMMENT '注册时间',
    UPDATED_AT DATETIME             NULL COMMENT '更新时间',
    ID_CARD    VARCHAR(20)          NULL COMMENT '身份证号码',
    BIRTHDAY   VARCHAR(20)          NULL COMMENT '生日',
    ADDRESS    VARCHAR(50)          NULL COMMENT '具体地址',
    PROVINCE   VARCHAR(50)          NULL COMMENT '省',
    CITY       VARCHAR(50)          NULL COMMENT '市',
    AREA       VARCHAR(50)          NULL COMMENT '地区',
  UNIQUE INDEX `UNI_USER_PASSPORT`(`PASSPORT`) USING BTREE,
  UNIQUE INDEX `UNI_USER_NICKNAME`(`NICKNAME`) USING BTREE
) ENGINE = INNODB AUTO_INCREMENT = 18  COMMENT = '用户基础表' ROW_FORMAT = DYNAMIC;

DROP TABLE IF EXISTS `PROVINCE_CITY_CODE`;
CREATE TABLE  PROVINCE_CITY_CODE
(
    CODE          VARCHAR(100) NULL COMMENT '六位地区码',
    NAME          VARCHAR(100)  NULL COMMENT '六位地区码含义',
    PROVINCE_CODE VARCHAR(100) NULL COMMENT '前两位省市码',
    PROVINCE_NAME VARCHAR(100)  NULL COMMENT '前两位省市码含义',
    CITY_CODE     VARCHAR(100) NULL COMMENT '前四位市县码',
    CITY_NAME     VARCHAR(100)  NULL COMMENT '前四位市县码含义',
    AREA_CODE     VARCHAR(100) NULL COMMENT '前六位地区码',
    AREA_NAME     VARCHAR(100)  NULL COMMENT '前六位地区码含义'
)
    COMMENT '身份证前六位含义';

