/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50738
 Source Host           : localhost:3306
 Source Schema         : ucenter

 Target Server Type    : MySQL
 Target Server Version : 50738
 File Encoding         : 65001

 Date: 11/08/2022 20:32:29
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user_devices
-- ----------------------------
DROP TABLE IF EXISTS `user_devices`;
CREATE TABLE `user_devices`  (
  `id` bigint(1) UNSIGNED NOT NULL AUTO_INCREMENT,
  `uid` bigint(1) UNSIGNED NULL DEFAULT NULL COMMENT '用户编号',
  `deviceid` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '设备唯一编号',
  `platform` tinyint(1) UNSIGNED NULL DEFAULT NULL COMMENT '平台,1-Android 2-IOS 3-web',
  `version` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '版本',
  `brand` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '品牌',
  `devicemodel` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '手机型号',
  `ip` int(1) UNSIGNED NULL DEFAULT NULL COMMENT 'ip地址',
  `addtime` bigint(1) UNSIGNED NULL DEFAULT NULL COMMENT '添加时间',
  `tp` tinyint(1) UNSIGNED NULL DEFAULT 0 COMMENT '0为登录,1为注册',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `uid`(`uid`) USING BTREE,
  INDEX `deviceid`(`deviceid`) USING BTREE,
  INDEX `tp`(`tp`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13599 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户设备信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` bigint(1) UNSIGNED NOT NULL AUTO_INCREMENT,
  `pid` bigint(1) UNSIGNED NULL DEFAULT 0 COMMENT '推荐人id',
  `invite` char(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '邀请码',
  `chain` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '用户关系链',
  `account` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '登录账号,有则唯一',
  `mail` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '邮箱号,有则唯一',
  `phone` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '手机号,有则唯一',
  `mailvery` tinyint(1) NULL DEFAULT 0 COMMENT '邮箱是否验证,1为已验证',
  `phonevery` tinyint(1) NULL DEFAULT 0 COMMENT '手机是否验证,1为已验证',
  `pwd` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '密码',
  `nickname` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '昵称',
  `avatar` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '头像地址',
  `background` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '主页背景图片',
  `signature` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '签名',
  `visits` int(11) NULL DEFAULT 0 COMMENT '访问量',
  `addtime` int(1) UNSIGNED NULL DEFAULT 0 COMMENT '注册时间',
  `status` tinyint(1) NULL DEFAULT 1 COMMENT '账号状态,1为正常,其他值均为不正常',
  `sex` tinyint(1) NULL DEFAULT 0 COMMENT '性别,0保密,1男，2女',
  `height` tinyint(1) UNSIGNED NULL DEFAULT 0 COMMENT '身高cm',
  `weight` float(5, 2) UNSIGNED NULL DEFAULT 0.00 COMMENT '体重kg',
  `birth` int(1) UNSIGNED NULL DEFAULT NULL COMMENT '生日',
  `age` tinyint(1) UNSIGNED NULL DEFAULT NULL COMMENT '年龄',
  `job` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '职业',
  `income` tinyint(1) NULL DEFAULT NULL COMMENT '收入',
  `emotion` tinyint(1) NULL DEFAULT 0 COMMENT '情感状态',
  `constellation` tinyint(1) UNSIGNED NULL DEFAULT NULL COMMENT '星座',
  `edu` tinyint(1) UNSIGNED NULL DEFAULT NULL COMMENT '学历',
  `temperament` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '性格,可设置多个',
  `ip` int(1) UNSIGNED NULL DEFAULT NULL COMMENT '注册时的ipv4地址',
  `country` smallint(1) UNSIGNED NULL DEFAULT 0 COMMENT '国家id',
  `province` smallint(1) UNSIGNED NULL DEFAULT NULL COMMENT '省份id',
  `city` smallint(1) UNSIGNED NULL DEFAULT NULL COMMENT '城市id',
  `singleid` int(1) NULL DEFAULT 0 COMMENT '单点登录token id',
  `lang` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户手动选择的语言',
  `currency` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户手动选择的币种',
  `timezone` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户时区',
  `platform` tinyint(1) UNSIGNED NULL DEFAULT 0 COMMENT '用户注册平台,0未知,1安卓,2苹果,3web',
  `md5` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户设备和ip变更验证',
  `private` tinyint(1) UNSIGNED NULL DEFAULT 0 COMMENT '是否是私密账号',
  `staff` tinyint(1) UNSIGNED NULL DEFAULT 0 COMMENT '是否是员工',
  `realuser` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '实拍用户本人图片地址',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `account`(`account`) USING BTREE,
  INDEX `mail`(`mail`) USING BTREE,
  INDEX `phone`(`phone`) USING BTREE,
  INDEX `status`(`status`) USING BTREE,
  INDEX `private`(`private`) USING BTREE,
  INDEX `invite`(`invite`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1484 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
