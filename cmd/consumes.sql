/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 50734
 Source Host           : 127.0.0.1:3306
 Source Schema         : ucenter

 Target Server Type    : MySQL
 Target Server Version : 50734
 File Encoding         : 65001

 Date: 17/08/2022 17:46:30
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for consumes
-- ----------------------------
DROP TABLE IF EXISTS `consumes`;
CREATE TABLE `consumes`  (
  `id` int(1) UNSIGNED NOT NULL AUTO_INCREMENT,
  `uid` int(1) UNSIGNED NOT NULL COMMENT '扣费用户id',
  `connect_id` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '连通id',
  `voice` tinyint(1) UNSIGNED NOT NULL COMMENT '1-语音, 2-视频',
  `start` int(1) UNSIGNED NOT NULL COMMENT '开始时间',
  `uptime` int(1) UNSIGNED NOT NULL COMMENT '上一次更新时间',
  `end` int(1) UNSIGNED NULL DEFAULT NULL COMMENT '结束时间',
  `usetime` int(1) UNSIGNED NULL DEFAULT 1 COMMENT '消耗时长',
  `seccost` smallint(1) UNSIGNED NOT NULL COMMENT '每分钟消耗代币数',
  `cost` int(1) UNSIGNED NULL DEFAULT 0 COMMENT '本次消耗',
  `status` tinyint(1) UNSIGNED NULL DEFAULT 0 COMMENT '是否正常结束',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `uid`(`uid`) USING BTREE,
  INDEX `status`(`status`) USING BTREE,
  INDEX `cost`(`cost`) USING BTREE,
  INDEX `voice`(`voice`) USING BTREE,
  INDEX `connect_id`(`connect_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '语音视频记录表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for orders
-- ----------------------------
DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders`  (
  `id` int(1) UNSIGNED NOT NULL AUTO_INCREMENT,
  `uid` int(1) UNSIGNED NOT NULL COMMENT 'user表id',
  `pid` int(1) UNSIGNED NOT NULL COMMENT 'pay_programs表id',
  `orderid` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '第三方订单id',
  `amount` decimal(10, 2) UNSIGNED NOT NULL COMMENT '支付金额',
  `bi` int(1) UNSIGNED NULL DEFAULT 0 COMMENT '代币数量',
  `addtime` int(1) UNSIGNED NULL DEFAULT NULL,
  `status` tinyint(1) NULL DEFAULT 0 COMMENT '支付状态,-1:订单无效 0-待支付 1-成功 2-已保存 3-需要付款人进一步动作 4-批准其他支付',
  `pay_way` tinyint(1) UNSIGNED NULL DEFAULT 1 COMMENT '1-paypal  2-apple pay',
  `paytime` int(1) UNSIGNED NULL DEFAULT NULL COMMENT '付款成功时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `uid`(`uid`) USING BTREE,
  INDEX `status`(`status`) USING BTREE,
  INDEX `pid`(`pid`) USING BTREE,
  INDEX `orderid`(`orderid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
