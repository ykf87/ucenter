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

 Date: 24/08/2022 14:40:11
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for languages
-- ----------------------------
DROP TABLE IF EXISTS `languages`;
CREATE TABLE `languages`  (
  `id` smallint(1) UNSIGNED NOT NULL AUTO_INCREMENT,
  `iso` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `sort` smallint(1) UNSIGNED NULL DEFAULT 0,
  `status` tinyint(1) UNSIGNED NULL DEFAULT 1,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `iso`(`iso`) USING BTREE,
  INDEX `status`(`status`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统语言表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of languages
-- ----------------------------
INSERT INTO `languages` VALUES (1, 'en', 'English', 999, 1);
INSERT INTO `languages` VALUES (2, 'zh-CN', '简体中文', 998, 1);
INSERT INTO `languages` VALUES (3, 'zh-TW', '繁体中文', 0, 0);
INSERT INTO `languages` VALUES (4, 'ja', 'やまと', 0, 0);
INSERT INTO `languages` VALUES (5, 'it', 'Italiano', 0, 0);
INSERT INTO `languages` VALUES (6, 'fr', 'Français', 0, 0);
INSERT INTO `languages` VALUES (7, 'de', 'Deutsch', 0, 0);
INSERT INTO `languages` VALUES (8, 'ru', 'Русский', 0, 0);

SET FOREIGN_KEY_CHECKS = 1;
