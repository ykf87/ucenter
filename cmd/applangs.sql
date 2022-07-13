/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 50737
 Source Host           : localhost:3306
 Source Schema         : ucenter

 Target Server Type    : MySQL
 Target Server Version : 50737
 File Encoding         : 65001

 Date: 12/07/2022 22:21:27
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for en_applangs
-- ----------------------------
DROP TABLE IF EXISTS `en_applangs`;
CREATE TABLE `en_applangs`  (
  `key` varchar(160) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '如果key过长,则使用缩写代替,首字母大写',
  `path` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '页面标识,指明这个key是哪个页面特有的,首字母小写',
  `val` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`key`, `path`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'app端多语言表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of en_applangs
-- ----------------------------
INSERT INTO `en_applangs` VALUES ('forgot', 'login', 'Forgot password?');
INSERT INTO `en_applangs` VALUES ('Tips', 'login', 'Thank you for your trust in us, we believe you can have fun here and get what you want!');
INSERT INTO `en_applangs` VALUES ('Tips', 'sign', 'We believe you won\'t regret signing up, but if you do, you can cancel your account to remove all traces of you!');
INSERT INTO `en_applangs` VALUES ('Welcome', 'login', 'Welcome to {{$1}}!');

-- ----------------------------
-- Table structure for zh-cn_applangs
-- ----------------------------
DROP TABLE IF EXISTS `zh-cn_applangs`;
CREATE TABLE `zh-cn_applangs`  (
  `key` varchar(160) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '如果key过长,则使用缩写代替,首字母大写',
  `path` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '页面标识,指明这个key是哪个页面特有的,首字母小写',
  `val` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`key`, `path`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'app端多语言表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of zh-cn_applangs
-- ----------------------------
INSERT INTO `zh-cn_applangs` VALUES ('Already have an account?', '', '已有账号?');
INSERT INTO `zh-cn_applangs` VALUES ('and acknowledge that you have read our', '', '同时确认您已经阅读我们的');
INSERT INTO `zh-cn_applangs` VALUES ('By continuing, you agree to our', '', '继续即表示您同意我们的');
INSERT INTO `zh-cn_applangs` VALUES ('Don\'t have an account yet?', '', '还没有账号?');
INSERT INTO `zh-cn_applangs` VALUES ('E-mail Captcha', '', '邮箱验证码');
INSERT INTO `zh-cn_applangs` VALUES ('Email address', '', '邮箱地址');
INSERT INTO `zh-cn_applangs` VALUES ('Forgot', 'login', '忘记密码?');
INSERT INTO `zh-cn_applangs` VALUES ('Get Code', '', '获取验证码');
INSERT INTO `zh-cn_applangs` VALUES ('Login', '', '登录');
INSERT INTO `zh-cn_applangs` VALUES ('Privacy Policy', '', '隐私政策');
INSERT INTO `zh-cn_applangs` VALUES ('Retrieve password', '', '找回密码');
INSERT INTO `zh-cn_applangs` VALUES ('Sign up', '', '注册');
INSERT INTO `zh-cn_applangs` VALUES ('Terms of Service', '', '服务条款');
INSERT INTO `zh-cn_applangs` VALUES ('Tips', 'login', '谢谢您的使用，祝您有一个愉快的经历!');
INSERT INTO `zh-cn_applangs` VALUES ('Tips', 'sign', '很高兴您选择我们,祝您有个愉快的体验!');
INSERT INTO `zh-cn_applangs` VALUES ('Welcome', 'login', '欢迎使用！');
INSERT INTO `zh-cn_applangs` VALUES ('Your Password', '', '你的密码');

SET FOREIGN_KEY_CHECKS = 1;
