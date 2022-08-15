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

 Date: 15/08/2022 19:43:23
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for orders
-- ----------------------------
DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders`  (
  `id` int(1) UNSIGNED NOT NULL AUTO_INCREMENT,
  `uid` int(1) UNSIGNED NOT NULL COMMENT 'user表id',
  `pid` int(1) UNSIGNED NOT NULL COMMENT 'pay_programs表id',
  `addtime` int(1) UNSIGNED NULL DEFAULT NULL,
  `status` tinyint(1) NULL DEFAULT 0 COMMENT '支付状态',
  `pay_way` tinyint(1) UNSIGNED NULL DEFAULT 1 COMMENT '1-paypal  2-apple pay',
  `paytime` int(1) UNSIGNED NULL DEFAULT NULL COMMENT '付款成功时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `uid`(`uid`) USING BTREE,
  INDEX `status`(`status`) USING BTREE,
  INDEX `pid`(`pid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pay_programs
-- ----------------------------
DROP TABLE IF EXISTS `pay_programs`;
CREATE TABLE `pay_programs`  (
  `id` int(1) UNSIGNED NOT NULL AUTO_INCREMENT,
  `price` decimal(10, 2) NOT NULL COMMENT '金额,美金',
  `bi` int(1) UNSIGNED NULL DEFAULT 0 COMMENT '获得的代币,平台货币',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `status` tinyint(1) NULL DEFAULT 1 COMMENT '状态',
  `used` int(1) UNSIGNED NULL DEFAULT 0 COMMENT '充值此方案次数',
  `pin` tinyint(1) UNSIGNED NULL DEFAULT 0 COMMENT '是否置顶',
  `pin_time` int(1) UNSIGNED NULL DEFAULT 0 COMMENT '置顶时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of pay_programs
-- ----------------------------
INSERT INTO `pay_programs` VALUES (1, 10.00, 100, NULL, 1, 0, 0, 0);
INSERT INTO `pay_programs` VALUES (2, 50.00, 550, NULL, 1, 0, 0, 0);
INSERT INTO `pay_programs` VALUES (3, 100.00, 1200, NULL, 1, 0, 0, 0);
INSERT INTO `pay_programs` VALUES (4, 200.00, 3000, NULL, 1, 0, 0, 0);

SET FOREIGN_KEY_CHECKS = 1;
