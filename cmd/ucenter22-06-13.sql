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

 Date: 13/06/2022 12:05:49
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for countries
-- ----------------------------
DROP TABLE IF EXISTS `countries`;
CREATE TABLE `countries`  (
  `id` smallint(1) UNSIGNED NOT NULL AUTO_INCREMENT,
  `iso` char(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `iso3` char(3) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `phonecode` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `timezone` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '时区 ',
  `lat` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `lon` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `emoji` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `currency` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '默认货币',
  `flags` smallint(1) UNSIGNED NULL DEFAULT NULL COMMENT '国旗',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `iso`(`iso`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 237 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of countries
-- ----------------------------
INSERT INTO `countries` VALUES (1, 'CN', 'CHN', '86', 'Asia/Shanghai', '35.00000000', '105.00000000', '🇨🇳', 'CNY', NULL);
INSERT INTO `countries` VALUES (2, 'AL', 'ALB', '355', 'Europe/Tirane', '41.00000000', '20.00000000', '🇦🇱', 'ALL', NULL);
INSERT INTO `countries` VALUES (3, 'DZ', 'DZA', '213', 'Africa/Algiers', '28.00000000', '3.00000000', '🇩🇿', 'DZD', NULL);
INSERT INTO `countries` VALUES (4, 'AF', 'AFG', '93', 'Asia/Kabul', '33.00000000', '65.00000000', '🇦🇫', 'AFN', NULL);
INSERT INTO `countries` VALUES (5, 'AR', 'ARG', '54', 'America/Argentina/Buenos_Aires', '-34.00000000', '-64.00000000', '🇦🇷', 'ARS', NULL);
INSERT INTO `countries` VALUES (6, 'AE', 'ARE', '971', 'Asia/Dubai', '24.00000000', '54.00000000', '🇦🇪', 'AED', NULL);
INSERT INTO `countries` VALUES (7, 'AW', 'ABW', '297', 'America/Aruba', '12.50000000', '-69.96666666', '🇦🇼', 'AWG', NULL);
INSERT INTO `countries` VALUES (8, 'OM', 'OMN', '968', 'Asia/Muscat', '21.00000000', '57.00000000', '🇴🇲', 'OMR', NULL);
INSERT INTO `countries` VALUES (9, 'AZ', 'AZE', '994', 'Asia/Baku', '40.50000000', '47.50000000', '🇦🇿', 'AZN', NULL);
INSERT INTO `countries` VALUES (10, 'EG', 'EGY', '20', 'Africa/Cairo', '27.00000000', '30.00000000', '🇪🇬', 'EGP', NULL);
INSERT INTO `countries` VALUES (11, 'ET', 'ETH', '251', 'Africa/Addis_Ababa', '8.00000000', '38.00000000', '🇪🇹', 'ETB', NULL);
INSERT INTO `countries` VALUES (12, 'IE', 'IRL', '353', 'Europe/Dublin', '53.00000000', '-8.00000000', '🇮🇪', 'EUR', NULL);
INSERT INTO `countries` VALUES (13, 'EE', 'EST', '372', 'Europe/Tallinn', '59.00000000', '26.00000000', '🇪🇪', 'EUR', NULL);
INSERT INTO `countries` VALUES (14, 'AD', 'AND', '376', 'Europe/Andorra', '42.50000000', '1.50000000', '🇦🇩', 'EUR', NULL);
INSERT INTO `countries` VALUES (15, 'AO', 'AGO', '244', 'Africa/Luanda', '-12.50000000', '18.50000000', '🇦🇴', 'AOA', NULL);
INSERT INTO `countries` VALUES (16, 'AI', 'AIA', '1-264', 'America/Anguilla', '18.25000000', '-63.16666666', '🇦🇮', 'XCD', NULL);
INSERT INTO `countries` VALUES (17, 'AG', 'ATG', '1-268', 'America/Antigua', '17.05000000', '-61.80000000', '🇦🇬', 'XCD', NULL);
INSERT INTO `countries` VALUES (18, 'AU', 'AUS', '61', 'Antarctica/Macquarie', '-27.00000000', '133.00000000', '🇦🇺', 'AUD', NULL);
INSERT INTO `countries` VALUES (19, 'AT', 'AUT', '43', 'Europe/Vienna', '47.33333333', '13.33333333', '🇦🇹', 'EUR', NULL);
INSERT INTO `countries` VALUES (20, 'AX', 'ALA', '358-18', 'Europe/Mariehamn', '60.11666700', '19.90000000', '🇦🇽', 'EUR', NULL);
INSERT INTO `countries` VALUES (21, 'BB', 'BRB', '1-246', 'America/Barbados', '13.16666666', '-59.53333333', '🇧🇧', 'BBD', NULL);
INSERT INTO `countries` VALUES (22, 'PG', 'PNG', '675', 'Pacific/Bougainville', '-6.00000000', '147.00000000', '🇵🇬', 'PGK', NULL);
INSERT INTO `countries` VALUES (23, 'BS', 'BHS', '1-242', 'America/Nassau', '24.25000000', '-76.00000000', '🇧🇸', 'BSD', NULL);
INSERT INTO `countries` VALUES (24, 'PK', 'PAK', '92', 'Asia/Karachi', '30.00000000', '70.00000000', '🇵🇰', 'PKR', NULL);
INSERT INTO `countries` VALUES (25, 'PY', 'PRY', '595', 'America/Asuncion', '-23.00000000', '-58.00000000', '🇵🇾', 'PYG', NULL);
INSERT INTO `countries` VALUES (26, 'PS', 'PSE', '970', 'Asia/Gaza', '31.90000000', '35.20000000', '🇵🇸', 'ILS', NULL);
INSERT INTO `countries` VALUES (27, 'BH', 'BHR', '973', 'Asia/Bahrain', '26.00000000', '50.55000000', '🇧🇭', 'BHD', NULL);
INSERT INTO `countries` VALUES (28, 'PA', 'PAN', '507', 'America/Panama', '9.00000000', '-80.00000000', '🇵🇦', 'PAB', NULL);
INSERT INTO `countries` VALUES (29, 'BR', 'BRA', '55', 'America/Araguaina', '-10.00000000', '-55.00000000', '🇧🇷', 'BRL', NULL);
INSERT INTO `countries` VALUES (30, 'BY', 'BLR', '375', 'Europe/Minsk', '53.00000000', '28.00000000', '🇧🇾', 'BYN', NULL);
INSERT INTO `countries` VALUES (31, 'BM', 'BMU', '1-441', 'Atlantic/Bermuda', '32.33333333', '-64.75000000', '🇧🇲', 'BMD', NULL);
INSERT INTO `countries` VALUES (32, 'BG', 'BGR', '359', 'Europe/Sofia', '43.00000000', '25.00000000', '🇧🇬', 'BGN', NULL);
INSERT INTO `countries` VALUES (33, 'MP', 'MNP', '1-670', 'Pacific/Saipan', '15.20000000', '145.75000000', '🇲🇵', 'USD', NULL);
INSERT INTO `countries` VALUES (34, 'BJ', 'BEN', '229', 'Africa/Porto-Novo', '9.50000000', '2.25000000', '🇧🇯', 'XOF', NULL);
INSERT INTO `countries` VALUES (35, 'BE', 'BEL', '32', 'Europe/Brussels', '50.83333333', '4.00000000', '🇧🇪', 'EUR', NULL);
INSERT INTO `countries` VALUES (36, 'IS', 'ISL', '354', 'Atlantic/Reykjavik', '65.00000000', '-18.00000000', '🇮🇸', 'ISK', NULL);
INSERT INTO `countries` VALUES (37, 'PR', 'PRI', '1-787 and 1-939', 'America/Puerto_Rico', '18.25000000', '-66.50000000', '🇵🇷', 'USD', NULL);
INSERT INTO `countries` VALUES (38, 'PL', 'POL', '48', 'Europe/Warsaw', '52.00000000', '20.00000000', '🇵🇱', 'PLN', NULL);
INSERT INTO `countries` VALUES (39, 'BO', 'BOL', '591', 'America/La_Paz', '-17.00000000', '-65.00000000', '🇧🇴', 'BOB', NULL);
INSERT INTO `countries` VALUES (40, 'BA', 'BIH', '387', 'Europe/Sarajevo', '44.00000000', '18.00000000', '🇧🇦', 'BAM', NULL);
INSERT INTO `countries` VALUES (41, 'BW', 'BWA', '267', 'Africa/Gaborone', '-22.00000000', '24.00000000', '🇧🇼', 'BWP', NULL);
INSERT INTO `countries` VALUES (42, 'BZ', 'BLZ', '501', 'America/Belize', '17.25000000', '-88.75000000', '🇧🇿', 'BZD', NULL);
INSERT INTO `countries` VALUES (43, 'BT', 'BTN', '975', 'Asia/Thimphu', '27.50000000', '90.50000000', '🇧🇹', 'BTN', NULL);
INSERT INTO `countries` VALUES (44, 'BF', 'BFA', '226', 'Africa/Ouagadougou', '13.00000000', '-2.00000000', '🇧🇫', 'XOF', NULL);
INSERT INTO `countries` VALUES (45, 'BI', 'BDI', '257', 'Africa/Bujumbura', '-3.50000000', '30.00000000', '🇧🇮', 'BIF', NULL);
INSERT INTO `countries` VALUES (46, 'BV', 'BVT', '0055', 'Europe/Oslo', '-54.43333333', '3.40000000', '🇧🇻', 'NOK', NULL);
INSERT INTO `countries` VALUES (47, 'KP', 'PRK', '850', 'Asia/Pyongyang', '40.00000000', '127.00000000', '🇰🇵', 'KPW', NULL);
INSERT INTO `countries` VALUES (48, 'DK', 'DNK', '45', 'Europe/Copenhagen', '56.00000000', '10.00000000', '🇩🇰', 'DKK', NULL);
INSERT INTO `countries` VALUES (49, 'DE', 'DEU', '49', 'Europe/Berlin', '51.00000000', '9.00000000', '🇩🇪', 'EUR', NULL);
INSERT INTO `countries` VALUES (50, 'TL', 'TLS', '670', 'Asia/Dili', '-8.83333333', '125.91666666', '🇹🇱', 'USD', NULL);
INSERT INTO `countries` VALUES (51, 'TG', 'TGO', '228', 'Africa/Lome', '8.00000000', '1.16666666', '🇹🇬', 'XOF', NULL);
INSERT INTO `countries` VALUES (52, 'DM', 'DMA', '1-767', 'America/Dominica', '15.41666666', '-61.33333333', '🇩🇲', 'XCD', NULL);
INSERT INTO `countries` VALUES (53, 'DO', 'DOM', '1-809 and 1-829', 'America/Santo_Domingo', '19.00000000', '-70.66666666', '🇩🇴', 'DOP', NULL);
INSERT INTO `countries` VALUES (54, 'RU', 'RUS', '7', 'Asia/Anadyr', '60.00000000', '100.00000000', '🇷🇺', 'RUB', NULL);
INSERT INTO `countries` VALUES (55, 'EC', 'ECU', '593', 'America/Guayaquil', '-2.00000000', '-77.50000000', '🇪🇨', 'USD', NULL);
INSERT INTO `countries` VALUES (56, 'ER', 'ERI', '291', 'Africa/Asmara', '15.00000000', '39.00000000', '🇪🇷', 'ERN', NULL);
INSERT INTO `countries` VALUES (57, 'FR', 'FRA', '33', 'Europe/Paris', '46.00000000', '2.00000000', '🇫🇷', 'EUR', NULL);
INSERT INTO `countries` VALUES (58, 'FO', 'FRO', '298', 'Atlantic/Faroe', '62.00000000', '-7.00000000', '🇫🇴', 'DKK', NULL);
INSERT INTO `countries` VALUES (59, 'PF', 'PYF', '689', 'Pacific/Gambier', '-15.00000000', '-140.00000000', '🇵🇫', 'XPF', NULL);
INSERT INTO `countries` VALUES (60, 'GF', 'GUF', '594', 'America/Cayenne', '4.00000000', '-53.00000000', '🇬🇫', 'EUR', NULL);
INSERT INTO `countries` VALUES (61, 'TF', 'ATF', '262', 'Indian/Kerguelen', '-49.25000000', '69.16700000', '🇹🇫', 'EUR', NULL);
INSERT INTO `countries` VALUES (62, 'VA', 'VAT', '379', 'Europe/Vatican', '41.90000000', '12.45000000', '🇻🇦', 'EUR', NULL);
INSERT INTO `countries` VALUES (63, 'PH', 'PHL', '63', 'Asia/Manila', '13.00000000', '122.00000000', '🇵🇭', 'PHP', NULL);
INSERT INTO `countries` VALUES (64, 'FJ', 'FJI', '679', 'Pacific/Fiji', '-18.00000000', '175.00000000', '🇫🇯', 'FJD', NULL);
INSERT INTO `countries` VALUES (65, 'FI', 'FIN', '358', 'Europe/Helsinki', '64.00000000', '26.00000000', '🇫🇮', 'EUR', NULL);
INSERT INTO `countries` VALUES (66, 'CV', 'CPV', '238', 'Atlantic/Cape_Verde', '16.00000000', '-24.00000000', '🇨🇻', 'CVE', NULL);
INSERT INTO `countries` VALUES (67, 'FK', 'FLK', '500', 'Atlantic/Stanley', '-51.75000000', '-59.00000000', '🇫🇰', 'FKP', NULL);
INSERT INTO `countries` VALUES (68, 'GM', 'GMB', '220', 'Africa/Banjul', '13.46666666', '-16.56666666', '🇬🇲', 'GMD', NULL);
INSERT INTO `countries` VALUES (69, 'CG', 'COG', '242', 'Africa/Brazzaville', '-1.00000000', '15.00000000', '🇨🇬', 'XAF', NULL);
INSERT INTO `countries` VALUES (70, 'CD', 'COD', '243', 'Africa/Kinshasa', '0.00000000', '25.00000000', '🇨🇩', 'CDF', NULL);
INSERT INTO `countries` VALUES (71, 'CO', 'COL', '57', 'America/Bogota', '4.00000000', '-72.00000000', '🇨🇴', 'COP', NULL);
INSERT INTO `countries` VALUES (72, 'CR', 'CRI', '506', 'America/Costa_Rica', '10.00000000', '-84.00000000', '🇨🇷', 'CRC', NULL);
INSERT INTO `countries` VALUES (73, 'GG', 'GGY', '44-1481', 'Europe/Guernsey', '49.46666666', '-2.58333333', '🇬🇬', 'GBP', NULL);
INSERT INTO `countries` VALUES (74, 'GD', 'GRD', '1-473', 'America/Grenada', '12.11666666', '-61.66666666', '🇬🇩', 'XCD', NULL);
INSERT INTO `countries` VALUES (75, 'GL', 'GRL', '299', 'America/Danmarkshavn', '72.00000000', '-40.00000000', '🇬🇱', 'DKK', NULL);
INSERT INTO `countries` VALUES (76, 'CU', 'CUB', '53', 'America/Havana', '21.50000000', '-80.00000000', '🇨🇺', 'CUP', NULL);
INSERT INTO `countries` VALUES (77, 'GP', 'GLP', '590', 'America/Guadeloupe', '16.25000000', '-61.58333300', '🇬🇵', 'EUR', NULL);
INSERT INTO `countries` VALUES (78, 'GU', 'GUM', '1-671', 'Pacific/Guam', '13.46666666', '144.78333333', '🇬🇺', 'USD', NULL);
INSERT INTO `countries` VALUES (79, 'GY', 'GUY', '592', 'America/Guyana', '5.00000000', '-59.00000000', '🇬🇾', 'GYD', NULL);
INSERT INTO `countries` VALUES (80, 'KZ', 'KAZ', '7', 'Asia/Almaty', '48.00000000', '68.00000000', '🇰🇿', 'KZT', NULL);
INSERT INTO `countries` VALUES (81, 'HT', 'HTI', '509', 'America/Port-au-Prince', '19.00000000', '-72.41666666', '🇭🇹', 'HTG', NULL);
INSERT INTO `countries` VALUES (82, 'KR', 'KOR', '82', 'Asia/Seoul', '37.00000000', '127.50000000', '🇰🇷', 'KRW', NULL);
INSERT INTO `countries` VALUES (83, 'NL', 'NLD', '31', 'Europe/Amsterdam', '52.50000000', '5.75000000', '🇳🇱', 'EUR', NULL);
INSERT INTO `countries` VALUES (84, 'HM', 'HMD', '672', 'Indian/Kerguelen', '-53.10000000', '72.51666666', '🇭🇲', 'AUD', NULL);
INSERT INTO `countries` VALUES (85, 'HN', 'HND', '504', 'America/Tegucigalpa', '15.00000000', '-86.50000000', '🇭🇳', 'HNL', NULL);
INSERT INTO `countries` VALUES (86, 'KI', 'KIR', '686', 'Pacific/Enderbury', '1.41666666', '173.00000000', '🇰🇮', 'AUD', NULL);
INSERT INTO `countries` VALUES (87, 'DJ', 'DJI', '253', 'Africa/Djibouti', '11.50000000', '43.00000000', '🇩🇯', 'DJF', NULL);
INSERT INTO `countries` VALUES (88, 'KG', 'KGZ', '996', 'Asia/Bishkek', '41.00000000', '75.00000000', '🇰🇬', 'KGS', NULL);
INSERT INTO `countries` VALUES (89, 'GN', 'GIN', '224', 'Africa/Conakry', '11.00000000', '-10.00000000', '🇬🇳', 'GNF', NULL);
INSERT INTO `countries` VALUES (90, 'GW', 'GNB', '245', 'Africa/Bissau', '12.00000000', '-15.00000000', '🇬🇼', 'XOF', NULL);
INSERT INTO `countries` VALUES (91, 'CA', 'CAN', '1', 'America/Atikokan', '60.00000000', '-95.00000000', '🇨🇦', 'CAD', NULL);
INSERT INTO `countries` VALUES (92, 'GH', 'GHA', '233', 'Africa/Accra', '8.00000000', '-2.00000000', '🇬🇭', 'GHS', NULL);
INSERT INTO `countries` VALUES (93, 'GA', 'GAB', '241', 'Africa/Libreville', '-1.00000000', '11.75000000', '🇬🇦', 'XAF', NULL);
INSERT INTO `countries` VALUES (94, 'KH', 'KHM', '855', 'Asia/Phnom_Penh', '13.00000000', '105.00000000', '🇰🇭', 'KHR', NULL);
INSERT INTO `countries` VALUES (95, 'CZ', 'CZE', '420', 'Europe/Prague', '49.75000000', '15.50000000', '🇨🇿', 'CZK', NULL);
INSERT INTO `countries` VALUES (96, 'ZW', 'ZWE', '263', 'Africa/Harare', '-20.00000000', '30.00000000', '🇿🇼', 'ZWL', NULL);
INSERT INTO `countries` VALUES (97, 'CM', 'CMR', '237', 'Africa/Douala', '6.00000000', '12.00000000', '🇨🇲', 'XAF', NULL);
INSERT INTO `countries` VALUES (98, 'QA', 'QAT', '974', 'Asia/Qatar', '25.50000000', '51.25000000', '🇶🇦', 'QAR', NULL);
INSERT INTO `countries` VALUES (99, 'KY', 'CYM', '1-345', 'America/Cayman', '19.50000000', '-80.50000000', '🇰🇾', 'KYD', NULL);
INSERT INTO `countries` VALUES (100, 'CC', 'CCK', '61', 'Indian/Cocos', '-12.50000000', '96.83333333', '🇨🇨', 'AUD', NULL);
INSERT INTO `countries` VALUES (101, 'KM', 'COM', '269', 'Indian/Comoro', '-12.16666666', '44.25000000', '🇰🇲', 'KMF', NULL);
INSERT INTO `countries` VALUES (102, 'CI', 'CIV', '225', 'Africa/Abidjan', '8.00000000', '-5.00000000', '🇨🇮', 'XOF', NULL);
INSERT INTO `countries` VALUES (103, 'KW', 'KWT', '965', 'Asia/Kuwait', '29.50000000', '45.75000000', '🇰🇼', 'KWD', NULL);
INSERT INTO `countries` VALUES (104, 'HR', 'HRV', '385', 'Europe/Zagreb', '45.16666666', '15.50000000', '🇭🇷', 'HRK', NULL);
INSERT INTO `countries` VALUES (105, 'KE', 'KEN', '254', 'Africa/Nairobi', '1.00000000', '38.00000000', '🇰🇪', 'KES', NULL);
INSERT INTO `countries` VALUES (106, 'CK', 'COK', '682', 'Pacific/Rarotonga', '-21.23333333', '-159.76666666', '🇨🇰', 'NZD', NULL);
INSERT INTO `countries` VALUES (107, 'LV', 'LVA', '371', 'Europe/Riga', '57.00000000', '25.00000000', '🇱🇻', 'EUR', NULL);
INSERT INTO `countries` VALUES (108, 'LS', 'LSO', '266', 'Africa/Maseru', '-29.50000000', '28.50000000', '🇱🇸', 'LSL', NULL);
INSERT INTO `countries` VALUES (109, 'LA', 'LAO', '856', 'Asia/Vientiane', '18.00000000', '105.00000000', '🇱🇦', 'LAK', NULL);
INSERT INTO `countries` VALUES (110, 'LB', 'LBN', '961', 'Asia/Beirut', '33.83333333', '35.83333333', '🇱🇧', 'LBP', NULL);
INSERT INTO `countries` VALUES (111, 'LR', 'LBR', '231', 'Africa/Monrovia', '6.50000000', '-9.50000000', '🇱🇷', 'LRD', NULL);
INSERT INTO `countries` VALUES (112, 'LY', 'LBY', '218', 'Africa/Tripoli', '25.00000000', '17.00000000', '🇱🇾', 'LYD', NULL);
INSERT INTO `countries` VALUES (113, 'LT', 'LTU', '370', 'Europe/Vilnius', '56.00000000', '24.00000000', '🇱🇹', 'EUR', NULL);
INSERT INTO `countries` VALUES (114, 'LI', 'LIE', '423', 'Europe/Vaduz', '47.26666666', '9.53333333', '🇱🇮', 'CHF', NULL);
INSERT INTO `countries` VALUES (115, 'RE', 'REU', '262', 'Indian/Reunion', '-21.15000000', '55.50000000', '🇷🇪', 'EUR', NULL);
INSERT INTO `countries` VALUES (116, 'LU', 'LUX', '352', 'Europe/Luxembourg', '49.75000000', '6.16666666', '🇱🇺', 'EUR', NULL);
INSERT INTO `countries` VALUES (117, 'RW', 'RWA', '250', 'Africa/Kigali', '-2.00000000', '30.00000000', '🇷🇼', 'RWF', NULL);
INSERT INTO `countries` VALUES (118, 'RO', 'ROU', '40', 'Europe/Bucharest', '46.00000000', '25.00000000', '🇷🇴', 'RON', NULL);
INSERT INTO `countries` VALUES (119, 'MG', 'MDG', '261', 'Indian/Antananarivo', '-20.00000000', '47.00000000', '🇲🇬', 'MGA', NULL);
INSERT INTO `countries` VALUES (120, 'MV', 'MDV', '960', 'Indian/Maldives', '3.25000000', '73.00000000', '🇲🇻', 'MVR', NULL);
INSERT INTO `countries` VALUES (121, 'MT', 'MLT', '356', 'Europe/Malta', '35.83333333', '14.58333333', '🇲🇹', 'EUR', NULL);
INSERT INTO `countries` VALUES (122, 'MW', 'MWI', '265', 'Africa/Blantyre', '-13.50000000', '34.00000000', '🇲🇼', 'MWK', NULL);
INSERT INTO `countries` VALUES (123, 'MY', 'MYS', '60', 'Asia/Kuala_Lumpur', '2.50000000', '112.50000000', '🇲🇾', 'MYR', NULL);
INSERT INTO `countries` VALUES (124, 'ML', 'MLI', '223', 'Africa/Bamako', '17.00000000', '-4.00000000', '🇲🇱', 'XOF', NULL);
INSERT INTO `countries` VALUES (125, 'MK', 'MKD', '389', 'Europe/Skopje', '41.83333333', '22.00000000', '🇲🇰', 'MKD', NULL);
INSERT INTO `countries` VALUES (126, 'MH', 'MHL', '692', 'Pacific/Kwajalein', '9.00000000', '168.00000000', '🇲🇭', 'USD', NULL);
INSERT INTO `countries` VALUES (127, 'MQ', 'MTQ', '596', 'America/Martinique', '14.66666700', '-61.00000000', '🇲🇶', 'EUR', NULL);
INSERT INTO `countries` VALUES (128, 'YT', 'MYT', '262', 'Indian/Mayotte', '-12.83333333', '45.16666666', '🇾🇹', 'EUR', NULL);
INSERT INTO `countries` VALUES (129, 'IM', 'IMN', '44-1624', 'Europe/Isle_of_Man', '54.25000000', '-4.50000000', '🇮🇲', 'GBP', NULL);
INSERT INTO `countries` VALUES (130, 'MU', 'MUS', '230', 'Indian/Mauritius', '-20.28333333', '57.55000000', '🇲🇺', 'MUR', NULL);
INSERT INTO `countries` VALUES (131, 'MR', 'MRT', '222', 'Africa/Nouakchott', '20.00000000', '-12.00000000', '🇲🇷', 'MRO', NULL);
INSERT INTO `countries` VALUES (132, 'US', 'USA', '1', 'America/Adak', '38.00000000', '-97.00000000', '🇺🇸', 'USD', NULL);
INSERT INTO `countries` VALUES (133, 'AS', 'ASM', '1-684', 'Pacific/Pago_Pago', '-14.33333333', '-170.00000000', '🇦🇸', 'USD', NULL);
INSERT INTO `countries` VALUES (134, 'UM', 'UMI', '1', 'Pacific/Midway', '0.00000000', '0.00000000', '🇺🇲', 'USD', NULL);
INSERT INTO `countries` VALUES (135, 'MN', 'MNG', '976', 'Asia/Choibalsan', '46.00000000', '105.00000000', '🇲🇳', 'MNT', NULL);
INSERT INTO `countries` VALUES (136, 'MS', 'MSR', '1-664', 'America/Montserrat', '16.75000000', '-62.20000000', '🇲🇸', 'XCD', NULL);
INSERT INTO `countries` VALUES (137, 'BD', 'BGD', '880', 'Asia/Dhaka', '24.00000000', '90.00000000', '🇧🇩', 'BDT', NULL);
INSERT INTO `countries` VALUES (138, 'FM', 'FSM', '691', 'Pacific/Chuuk', '6.91666666', '158.25000000', '🇫🇲', 'USD', NULL);
INSERT INTO `countries` VALUES (139, 'PE', 'PER', '51', 'America/Lima', '-10.00000000', '-76.00000000', '🇵🇪', 'PEN', NULL);
INSERT INTO `countries` VALUES (140, 'MM', 'MMR', '95', 'Asia/Yangon', '22.00000000', '98.00000000', '🇲🇲', 'MMK', NULL);
INSERT INTO `countries` VALUES (141, 'MD', 'MDA', '373', 'Europe/Chisinau', '47.00000000', '29.00000000', '🇲🇩', 'MDL', NULL);
INSERT INTO `countries` VALUES (142, 'MA', 'MAR', '212', 'Africa/Casablanca', '32.00000000', '-5.00000000', '🇲🇦', 'MAD', NULL);
INSERT INTO `countries` VALUES (143, 'MC', 'MCO', '377', 'Europe/Monaco', '43.73333333', '7.40000000', '🇲🇨', 'EUR', NULL);
INSERT INTO `countries` VALUES (144, 'MZ', 'MOZ', '258', 'Africa/Maputo', '-18.25000000', '35.00000000', '🇲🇿', 'MZN', NULL);
INSERT INTO `countries` VALUES (145, 'MX', 'MEX', '52', 'America/Bahia_Banderas', '23.00000000', '-102.00000000', '🇲🇽', 'MXN', NULL);
INSERT INTO `countries` VALUES (146, 'NA', 'NAM', '264', 'Africa/Windhoek', '-22.00000000', '17.00000000', '🇳🇦', 'NAD', NULL);
INSERT INTO `countries` VALUES (147, 'ZA', 'ZAF', '27', 'Africa/Johannesburg', '-29.00000000', '24.00000000', '🇿🇦', 'ZAR', NULL);
INSERT INTO `countries` VALUES (148, 'AQ', 'ATA', '672', 'Antarctica/Casey', '-74.65000000', '4.48000000', '🇦🇶', 'AAD', NULL);
INSERT INTO `countries` VALUES (149, 'GS', 'SGS', '500', 'Atlantic/South_Georgia', '-54.50000000', '-37.00000000', '🇬🇸', 'GBP', NULL);
INSERT INTO `countries` VALUES (150, 'NR', 'NRU', '674', 'Pacific/Nauru', '-0.53333333', '166.91666666', '🇳🇷', 'AUD', NULL);
INSERT INTO `countries` VALUES (151, 'NP', 'NPL', '977', 'Asia/Kathmandu', '28.00000000', '84.00000000', '🇳🇵', 'NPR', NULL);
INSERT INTO `countries` VALUES (152, 'NI', 'NIC', '505', 'America/Managua', '13.00000000', '-85.00000000', '🇳🇮', 'NIO', NULL);
INSERT INTO `countries` VALUES (153, 'NE', 'NER', '227', 'Africa/Niamey', '16.00000000', '8.00000000', '🇳🇪', 'XOF', NULL);
INSERT INTO `countries` VALUES (154, 'NG', 'NGA', '234', 'Africa/Lagos', '10.00000000', '8.00000000', '🇳🇬', 'NGN', NULL);
INSERT INTO `countries` VALUES (155, 'NU', 'NIU', '683', 'Pacific/Niue', '-19.03333333', '-169.86666666', '🇳🇺', 'NZD', NULL);
INSERT INTO `countries` VALUES (156, 'NO', 'NOR', '47', 'Europe/Oslo', '62.00000000', '10.00000000', '🇳🇴', 'NOK', NULL);
INSERT INTO `countries` VALUES (157, 'NF', 'NFK', '672', 'Pacific/Norfolk', '-29.03333333', '167.95000000', '🇳🇫', 'AUD', NULL);
INSERT INTO `countries` VALUES (158, 'PW', 'PLW', '680', 'Pacific/Palau', '7.50000000', '134.50000000', '🇵🇼', 'USD', NULL);
INSERT INTO `countries` VALUES (159, 'PN', 'PCN', '870', 'Pacific/Pitcairn', '-25.06666666', '-130.10000000', '🇵🇳', 'NZD', NULL);
INSERT INTO `countries` VALUES (160, 'PT', 'PRT', '351', 'Atlantic/Azores', '39.50000000', '-8.00000000', '🇵🇹', 'EUR', NULL);
INSERT INTO `countries` VALUES (161, 'GE', 'GEO', '995', 'Asia/Tbilisi', '42.00000000', '43.50000000', '🇬🇪', 'GEL', NULL);
INSERT INTO `countries` VALUES (162, 'JP', 'JPN', '81', 'Asia/Tokyo', '36.00000000', '138.00000000', '🇯🇵', 'JPY', NULL);
INSERT INTO `countries` VALUES (163, 'SE', 'SWE', '46', 'Europe/Stockholm', '62.00000000', '15.00000000', '🇸🇪', 'SEK', NULL);
INSERT INTO `countries` VALUES (164, 'CH', 'CHE', '41', 'Europe/Zurich', '47.00000000', '8.00000000', '🇨🇭', 'CHF', NULL);
INSERT INTO `countries` VALUES (165, 'SV', 'SLV', '503', 'America/El_Salvador', '13.83333333', '-88.91666666', '🇸🇻', 'USD', NULL);
INSERT INTO `countries` VALUES (166, 'WS', 'WSM', '685', 'Pacific/Apia', '-13.58333333', '-172.33333333', '🇼🇸', 'WST', NULL);
INSERT INTO `countries` VALUES (167, 'SL', 'SLE', '232', 'Africa/Freetown', '8.50000000', '-11.50000000', '🇸🇱', 'SLL', NULL);
INSERT INTO `countries` VALUES (168, 'SN', 'SEN', '221', 'Africa/Dakar', '14.00000000', '-14.00000000', '🇸🇳', 'XOF', NULL);
INSERT INTO `countries` VALUES (169, 'CY', 'CYP', '357', 'Asia/Famagusta', '35.00000000', '33.00000000', '🇨🇾', 'EUR', NULL);
INSERT INTO `countries` VALUES (170, 'SC', 'SYC', '248', 'Indian/Mahe', '-4.58333333', '55.66666666', '🇸🇨', 'SCR', NULL);
INSERT INTO `countries` VALUES (171, 'SA', 'SAU', '966', 'Asia/Riyadh', '25.00000000', '45.00000000', '🇸🇦', 'SAR', NULL);
INSERT INTO `countries` VALUES (172, 'CX', 'CXR', '61', 'Indian/Christmas', '-10.50000000', '105.66666666', '🇨🇽', 'AUD', NULL);
INSERT INTO `countries` VALUES (173, 'ST', 'STP', '239', 'Africa/Sao_Tome', '1.00000000', '7.00000000', '🇸🇹', 'STD', NULL);
INSERT INTO `countries` VALUES (174, 'SH', 'SHN', '290', 'Atlantic/St_Helena', '-15.95000000', '-5.70000000', '🇸🇭', 'SHP', NULL);
INSERT INTO `countries` VALUES (175, 'KN', 'KNA', '1-869', 'America/St_Kitts', '17.33333333', '-62.75000000', '🇰🇳', 'XCD', NULL);
INSERT INTO `countries` VALUES (176, 'LC', 'LCA', '1-758', 'America/St_Lucia', '13.88333333', '-60.96666666', '🇱🇨', 'XCD', NULL);
INSERT INTO `countries` VALUES (177, 'SM', 'SMR', '378', 'Europe/San_Marino', '43.76666666', '12.41666666', '🇸🇲', 'EUR', NULL);
INSERT INTO `countries` VALUES (178, 'PM', 'SPM', '508', 'America/Miquelon', '46.83333333', '-56.33333333', '🇵🇲', 'EUR', NULL);
INSERT INTO `countries` VALUES (179, 'VC', 'VCT', '1-784', 'America/St_Vincent', '13.25000000', '-61.20000000', '🇻🇨', 'XCD', NULL);
INSERT INTO `countries` VALUES (180, 'LK', 'LKA', '94', 'Asia/Colombo', '7.00000000', '81.00000000', '🇱🇰', 'LKR', NULL);
INSERT INTO `countries` VALUES (181, 'SK', 'SVK', '421', 'Europe/Bratislava', '48.66666666', '19.50000000', '🇸🇰', 'EUR', NULL);
INSERT INTO `countries` VALUES (182, 'SI', 'SVN', '386', 'Europe/Ljubljana', '46.11666666', '14.81666666', '🇸🇮', 'EUR', NULL);
INSERT INTO `countries` VALUES (183, 'SJ', 'SJM', '47', 'Arctic/Longyearbyen', '78.00000000', '20.00000000', '🇸🇯', 'NOK', NULL);
INSERT INTO `countries` VALUES (184, 'SZ', 'SWZ', '268', 'Africa/Mbabane', '-26.50000000', '31.50000000', '🇸🇿', 'SZL', NULL);
INSERT INTO `countries` VALUES (185, 'SD', 'SDN', '249', 'Africa/Khartoum', '15.00000000', '30.00000000', '🇸🇩', 'SDG', NULL);
INSERT INTO `countries` VALUES (186, 'SR', 'SUR', '597', 'America/Paramaribo', '4.00000000', '-56.00000000', '🇸🇷', 'SRD', NULL);
INSERT INTO `countries` VALUES (187, 'SB', 'SLB', '677', 'Pacific/Guadalcanal', '-8.00000000', '159.00000000', '🇸🇧', 'SBD', NULL);
INSERT INTO `countries` VALUES (188, 'SO', 'SOM', '252', 'Africa/Mogadishu', '10.00000000', '49.00000000', '🇸🇴', 'SOS', NULL);
INSERT INTO `countries` VALUES (189, 'TJ', 'TJK', '992', 'Asia/Dushanbe', '39.00000000', '71.00000000', '🇹🇯', 'TJS', NULL);
INSERT INTO `countries` VALUES (190, 'TH', 'THA', '66', 'Asia/Bangkok', '15.00000000', '100.00000000', '🇹🇭', 'THB', NULL);
INSERT INTO `countries` VALUES (191, 'TZ', 'TZA', '255', 'Africa/Dar_es_Salaam', '-6.00000000', '35.00000000', '🇹🇿', 'TZS', NULL);
INSERT INTO `countries` VALUES (192, 'TO', 'TON', '676', 'Pacific/Tongatapu', '-20.00000000', '-175.00000000', '🇹🇴', 'TOP', NULL);
INSERT INTO `countries` VALUES (193, 'TC', 'TCA', '1-649', 'America/Grand_Turk', '21.75000000', '-71.58333333', '🇹🇨', 'USD', NULL);
INSERT INTO `countries` VALUES (194, 'TT', 'TTO', '1-868', 'America/Port_of_Spain', '11.00000000', '-61.00000000', '🇹🇹', 'TTD', NULL);
INSERT INTO `countries` VALUES (195, 'TN', 'TUN', '216', 'Africa/Tunis', '34.00000000', '9.00000000', '🇹🇳', 'TND', NULL);
INSERT INTO `countries` VALUES (196, 'TV', 'TUV', '688', 'Pacific/Funafuti', '-8.00000000', '178.00000000', '🇹🇻', 'AUD', NULL);
INSERT INTO `countries` VALUES (197, 'TR', 'TUR', '90', 'Europe/Istanbul', '39.00000000', '35.00000000', '🇹🇷', 'TRY', NULL);
INSERT INTO `countries` VALUES (198, 'TM', 'TKM', '993', 'Asia/Ashgabat', '40.00000000', '60.00000000', '🇹🇲', 'TMT', NULL);
INSERT INTO `countries` VALUES (199, 'TK', 'TKL', '690', 'Pacific/Fakaofo', '-9.00000000', '-172.00000000', '🇹🇰', 'NZD', NULL);
INSERT INTO `countries` VALUES (200, 'WF', 'WLF', '681', 'Pacific/Wallis', '-13.30000000', '-176.20000000', '🇼🇫', 'XPF', NULL);
INSERT INTO `countries` VALUES (201, 'VU', 'VUT', '678', 'Pacific/Efate', '-16.00000000', '167.00000000', '🇻🇺', 'VUV', NULL);
INSERT INTO `countries` VALUES (202, 'GT', 'GTM', '502', 'America/Guatemala', '15.50000000', '-90.25000000', '🇬🇹', 'GTQ', NULL);
INSERT INTO `countries` VALUES (203, 'VI', 'VIR', '1-340', 'America/St_Thomas', '18.34000000', '-64.93000000', '🇻🇮', 'USD', NULL);
INSERT INTO `countries` VALUES (204, 'VG', 'VGB', '1-284', 'America/Tortola', '18.43138300', '-64.62305000', '🇻🇬', 'USD', NULL);
INSERT INTO `countries` VALUES (205, 'VE', 'VEN', '58', 'America/Caracas', '8.00000000', '-66.00000000', '🇻🇪', 'VEF', NULL);
INSERT INTO `countries` VALUES (206, 'BN', 'BRN', '673', 'Asia/Brunei', '4.50000000', '114.66666666', '🇧🇳', 'BND', NULL);
INSERT INTO `countries` VALUES (207, 'UG', 'UGA', '256', 'Africa/Kampala', '1.00000000', '32.00000000', '🇺🇬', 'UGX', NULL);
INSERT INTO `countries` VALUES (208, 'UA', 'UKR', '380', 'Europe/Kiev', '49.00000000', '32.00000000', '🇺🇦', 'UAH', NULL);
INSERT INTO `countries` VALUES (209, 'UY', 'URY', '598', 'America/Montevideo', '-33.00000000', '-56.00000000', '🇺🇾', 'UYU', NULL);
INSERT INTO `countries` VALUES (210, 'UZ', 'UZB', '998', 'Asia/Samarkand', '41.00000000', '64.00000000', '🇺🇿', 'UZS', NULL);
INSERT INTO `countries` VALUES (211, 'ES', 'ESP', '34', 'Africa/Ceuta', '40.00000000', '-4.00000000', '🇪🇸', 'EUR', NULL);
INSERT INTO `countries` VALUES (212, 'GR', 'GRC', '30', 'Europe/Athens', '39.00000000', '22.00000000', '🇬🇷', 'EUR', NULL);
INSERT INTO `countries` VALUES (213, 'SG', 'SGP', '65', 'Asia/Singapore', '1.36666666', '103.80000000', '🇸🇬', 'SGD', NULL);
INSERT INTO `countries` VALUES (214, 'NC', 'NCL', '687', 'Pacific/Noumea', '-21.50000000', '165.50000000', '🇳🇨', 'XPF', NULL);
INSERT INTO `countries` VALUES (215, 'NZ', 'NZL', '64', 'Pacific/Auckland', '-41.00000000', '174.00000000', '🇳🇿', 'NZD', NULL);
INSERT INTO `countries` VALUES (216, 'HU', 'HUN', '36', 'Europe/Budapest', '47.00000000', '20.00000000', '🇭🇺', 'HUF', NULL);
INSERT INTO `countries` VALUES (217, 'SY', 'SYR', '963', 'Asia/Damascus', '35.00000000', '38.00000000', '🇸🇾', 'SYP', NULL);
INSERT INTO `countries` VALUES (218, 'JM', 'JAM', '1-876', 'America/Jamaica', '18.25000000', '-77.50000000', '🇯🇲', 'JMD', NULL);
INSERT INTO `countries` VALUES (219, 'AM', 'ARM', '374', 'Asia/Yerevan', '40.00000000', '45.00000000', '🇦🇲', 'AMD', NULL);
INSERT INTO `countries` VALUES (220, 'YE', 'YEM', '967', 'Asia/Aden', '15.00000000', '48.00000000', '🇾🇪', 'YER', NULL);
INSERT INTO `countries` VALUES (221, 'IQ', 'IRQ', '964', 'Asia/Baghdad', '33.00000000', '44.00000000', '🇮🇶', 'IQD', NULL);
INSERT INTO `countries` VALUES (222, 'IR', 'IRN', '98', 'Asia/Tehran', '32.00000000', '53.00000000', '🇮🇷', 'IRR', NULL);
INSERT INTO `countries` VALUES (223, 'IL', 'ISR', '972', 'Asia/Jerusalem', '31.50000000', '34.75000000', '🇮🇱', 'ILS', NULL);
INSERT INTO `countries` VALUES (224, 'IT', 'ITA', '39', 'Europe/Rome', '42.83333333', '12.83333333', '🇮🇹', 'EUR', NULL);
INSERT INTO `countries` VALUES (225, 'IN', 'IND', '91', 'Asia/Kolkata', '20.00000000', '77.00000000', '🇮🇳', 'INR', NULL);
INSERT INTO `countries` VALUES (226, 'ID', 'IDN', '62', 'Asia/Jakarta', '-5.00000000', '120.00000000', '🇮🇩', 'IDR', NULL);
INSERT INTO `countries` VALUES (227, 'GB', 'GBR', '44', 'Europe/London', '54.00000000', '-2.00000000', '🇬🇧', 'GBP', NULL);
INSERT INTO `countries` VALUES (228, 'IO', 'IOT', '246', 'Indian/Chagos', '-6.00000000', '71.50000000', '🇮🇴', 'USD', NULL);
INSERT INTO `countries` VALUES (229, 'JO', 'JOR', '962', 'Asia/Amman', '31.00000000', '36.00000000', '🇯🇴', 'JOD', NULL);
INSERT INTO `countries` VALUES (230, 'VN', 'VNM', '84', 'Asia/Ho_Chi_Minh', '16.16666666', '107.83333333', '🇻🇳', 'VND', NULL);
INSERT INTO `countries` VALUES (231, 'ZM', 'ZMB', '260', 'Africa/Lusaka', '-15.00000000', '30.00000000', '🇿🇲', 'ZMW', NULL);
INSERT INTO `countries` VALUES (232, 'JE', 'JEY', '44-1534', 'Europe/Jersey', '49.25000000', '-2.16666666', '🇯🇪', 'GBP', NULL);
INSERT INTO `countries` VALUES (233, 'TD', 'TCD', '235', 'Africa/Ndjamena', '15.00000000', '19.00000000', '🇹🇩', 'XAF', NULL);
INSERT INTO `countries` VALUES (234, 'GI', 'GIB', '350', 'Europe/Gibraltar', '36.13333333', '-5.35000000', '🇬🇮', 'GIP', NULL);
INSERT INTO `countries` VALUES (235, 'CL', 'CHL', '56', 'America/Punta_Arenas', '-30.00000000', '-71.00000000', '🇨🇱', 'CLP', NULL);
INSERT INTO `countries` VALUES (236, 'CF', 'CAF', '236', 'Africa/Bangui', '7.00000000', '21.00000000', '🇨🇫', 'XAF', NULL);

-- ----------------------------
-- Table structure for countries_en
-- ----------------------------
DROP TABLE IF EXISTS `countries_en`;
CREATE TABLE `countries_en`  (
  `id` smallint(1) UNSIGNED NOT NULL,
  `name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of countries_en
-- ----------------------------
INSERT INTO `countries_en` VALUES (1, 'China');
INSERT INTO `countries_en` VALUES (2, 'Albania');
INSERT INTO `countries_en` VALUES (3, 'Algeria');
INSERT INTO `countries_en` VALUES (4, 'Afghanistan');
INSERT INTO `countries_en` VALUES (5, 'Argentina');
INSERT INTO `countries_en` VALUES (6, 'United Arab Emirates');
INSERT INTO `countries_en` VALUES (7, 'Aruba');
INSERT INTO `countries_en` VALUES (8, 'Oman');
INSERT INTO `countries_en` VALUES (9, 'Azerbaijan');
INSERT INTO `countries_en` VALUES (10, 'Egypt');
INSERT INTO `countries_en` VALUES (11, 'Ethiopia');
INSERT INTO `countries_en` VALUES (12, 'Ireland');
INSERT INTO `countries_en` VALUES (13, 'Estonia');
INSERT INTO `countries_en` VALUES (14, 'Andorra');
INSERT INTO `countries_en` VALUES (15, 'Angola');
INSERT INTO `countries_en` VALUES (16, 'Anguilla');
INSERT INTO `countries_en` VALUES (17, 'Antigua and Barbuda');
INSERT INTO `countries_en` VALUES (18, 'Australia');
INSERT INTO `countries_en` VALUES (19, 'Austria');
INSERT INTO `countries_en` VALUES (20, 'Åland Islands');
INSERT INTO `countries_en` VALUES (21, 'Barbados');
INSERT INTO `countries_en` VALUES (22, 'Papua New Guinea');
INSERT INTO `countries_en` VALUES (23, 'Bahamas');
INSERT INTO `countries_en` VALUES (24, 'Pakistan');
INSERT INTO `countries_en` VALUES (25, 'Paraguay');
INSERT INTO `countries_en` VALUES (26, 'Palestine');
INSERT INTO `countries_en` VALUES (27, 'Bahrain');
INSERT INTO `countries_en` VALUES (28, 'Panama');
INSERT INTO `countries_en` VALUES (29, 'Brazil');
INSERT INTO `countries_en` VALUES (30, 'Belarus');
INSERT INTO `countries_en` VALUES (31, 'Bermuda');
INSERT INTO `countries_en` VALUES (32, 'Bulgaria');
INSERT INTO `countries_en` VALUES (33, 'Northern Mariana Islands');
INSERT INTO `countries_en` VALUES (34, 'Benin');
INSERT INTO `countries_en` VALUES (35, 'Belgium');
INSERT INTO `countries_en` VALUES (36, 'Iceland');
INSERT INTO `countries_en` VALUES (37, 'Puerto Rico');
INSERT INTO `countries_en` VALUES (38, 'Poland');
INSERT INTO `countries_en` VALUES (39, 'Bolivia');
INSERT INTO `countries_en` VALUES (40, 'Bosnia and Herzegovina');
INSERT INTO `countries_en` VALUES (41, 'Botswana');
INSERT INTO `countries_en` VALUES (42, 'Belize');
INSERT INTO `countries_en` VALUES (43, 'Bhutan');
INSERT INTO `countries_en` VALUES (44, 'Burkina Faso');
INSERT INTO `countries_en` VALUES (45, 'Burundi');
INSERT INTO `countries_en` VALUES (46, 'Bouvet Island');
INSERT INTO `countries_en` VALUES (47, 'North Korea');
INSERT INTO `countries_en` VALUES (48, 'Denmark');
INSERT INTO `countries_en` VALUES (49, 'Germany');
INSERT INTO `countries_en` VALUES (50, 'East Timor');
INSERT INTO `countries_en` VALUES (51, 'togo');
INSERT INTO `countries_en` VALUES (52, 'Dominica');
INSERT INTO `countries_en` VALUES (53, 'Dominican Republic');
INSERT INTO `countries_en` VALUES (54, 'Russia');
INSERT INTO `countries_en` VALUES (55, 'Ecuador');
INSERT INTO `countries_en` VALUES (56, 'Eritrea');
INSERT INTO `countries_en` VALUES (57, 'France');
INSERT INTO `countries_en` VALUES (58, 'Faroe Islands');
INSERT INTO `countries_en` VALUES (59, 'French Polynesia');
INSERT INTO `countries_en` VALUES (60, 'French Guiana');
INSERT INTO `countries_en` VALUES (61, 'French Southern Territory');
INSERT INTO `countries_en` VALUES (62, 'Vatican');
INSERT INTO `countries_en` VALUES (63, 'the Philippines');
INSERT INTO `countries_en` VALUES (64, 'Fiji');
INSERT INTO `countries_en` VALUES (65, 'Finland');
INSERT INTO `countries_en` VALUES (66, 'Cape Verde');
INSERT INTO `countries_en` VALUES (67, 'frank islands');
INSERT INTO `countries_en` VALUES (68, 'Gambia');
INSERT INTO `countries_en` VALUES (69, 'Congo');
INSERT INTO `countries_en` VALUES (70, 'Democratic Republic of Congo');
INSERT INTO `countries_en` VALUES (71, 'Colombia');
INSERT INTO `countries_en` VALUES (72, 'Costa Rica');
INSERT INTO `countries_en` VALUES (73, 'Guernsey');
INSERT INTO `countries_en` VALUES (74, 'Grenada');
INSERT INTO `countries_en` VALUES (75, 'Greenland');
INSERT INTO `countries_en` VALUES (76, 'Cuba');
INSERT INTO `countries_en` VALUES (77, 'Guadeloupe');
INSERT INTO `countries_en` VALUES (78, 'Guam');
INSERT INTO `countries_en` VALUES (79, 'Guyana');
INSERT INTO `countries_en` VALUES (80, 'Kazakhstan');
INSERT INTO `countries_en` VALUES (81, 'Haiti');
INSERT INTO `countries_en` VALUES (82, 'South Korea');
INSERT INTO `countries_en` VALUES (83, 'Netherlands');
INSERT INTO `countries_en` VALUES (84, 'heard and macdonald islands');
INSERT INTO `countries_en` VALUES (85, 'Honduras');
INSERT INTO `countries_en` VALUES (86, 'Kiribati');
INSERT INTO `countries_en` VALUES (87, 'Djibouti');
INSERT INTO `countries_en` VALUES (88, 'Kyrgyzstan');
INSERT INTO `countries_en` VALUES (89, 'Guinea');
INSERT INTO `countries_en` VALUES (90, 'Guinea-Bissau');
INSERT INTO `countries_en` VALUES (91, 'Canada');
INSERT INTO `countries_en` VALUES (92, 'Ghana');
INSERT INTO `countries_en` VALUES (93, 'Gabon');
INSERT INTO `countries_en` VALUES (94, 'Cambodia');
INSERT INTO `countries_en` VALUES (95, 'Czech Republic');
INSERT INTO `countries_en` VALUES (96, 'Zimbabwe');
INSERT INTO `countries_en` VALUES (97, 'Cameroon');
INSERT INTO `countries_en` VALUES (98, 'Qatar');
INSERT INTO `countries_en` VALUES (99, 'Cayman Islands');
INSERT INTO `countries_en` VALUES (100, 'cocos islands');
INSERT INTO `countries_en` VALUES (101, 'Comoros');
INSERT INTO `countries_en` VALUES (102, 'Côte d\'Ivoire');
INSERT INTO `countries_en` VALUES (103, 'Kuwait');
INSERT INTO `countries_en` VALUES (104, 'Croatia');
INSERT INTO `countries_en` VALUES (105, 'Kenya');
INSERT INTO `countries_en` VALUES (106, 'Island');
INSERT INTO `countries_en` VALUES (107, 'Latvia');
INSERT INTO `countries_en` VALUES (108, 'Lesotho');
INSERT INTO `countries_en` VALUES (109, 'Laos');
INSERT INTO `countries_en` VALUES (110, 'Lebanon');
INSERT INTO `countries_en` VALUES (111, 'Liberia');
INSERT INTO `countries_en` VALUES (112, 'Libya');
INSERT INTO `countries_en` VALUES (113, 'Lithuania');
INSERT INTO `countries_en` VALUES (114, 'Liechtenstein');
INSERT INTO `countries_en` VALUES (115, 'Reunion');
INSERT INTO `countries_en` VALUES (116, 'Luxembourg');
INSERT INTO `countries_en` VALUES (117, 'Rwanda');
INSERT INTO `countries_en` VALUES (118, 'Romania');
INSERT INTO `countries_en` VALUES (119, 'Madagascar');
INSERT INTO `countries_en` VALUES (120, 'Maldives');
INSERT INTO `countries_en` VALUES (121, 'Malta');
INSERT INTO `countries_en` VALUES (122, 'Malawi');
INSERT INTO `countries_en` VALUES (123, 'Malaysia');
INSERT INTO `countries_en` VALUES (124, 'Mali');
INSERT INTO `countries_en` VALUES (125, 'Macedonia');
INSERT INTO `countries_en` VALUES (126, 'Marshall Islands');
INSERT INTO `countries_en` VALUES (127, 'Martinique');
INSERT INTO `countries_en` VALUES (128, 'mayotte');
INSERT INTO `countries_en` VALUES (129, 'isle of man');
INSERT INTO `countries_en` VALUES (130, 'Mauritius');
INSERT INTO `countries_en` VALUES (131, 'Mauritania');
INSERT INTO `countries_en` VALUES (132, 'U.S.');
INSERT INTO `countries_en` VALUES (133, 'American Samoa');
INSERT INTO `countries_en` VALUES (134, 'U.S. Outlying Islands');
INSERT INTO `countries_en` VALUES (135, 'Mongolia');
INSERT INTO `countries_en` VALUES (136, 'montserrat');
INSERT INTO `countries_en` VALUES (137, 'Bengal');
INSERT INTO `countries_en` VALUES (138, 'Micronesia');
INSERT INTO `countries_en` VALUES (139, 'Peru');
INSERT INTO `countries_en` VALUES (140, 'Myanmar');
INSERT INTO `countries_en` VALUES (141, 'Moldova');
INSERT INTO `countries_en` VALUES (142, 'Morocco');
INSERT INTO `countries_en` VALUES (143, 'Monaco');
INSERT INTO `countries_en` VALUES (144, 'Mozambique');
INSERT INTO `countries_en` VALUES (145, 'Mexico');
INSERT INTO `countries_en` VALUES (146, 'Namibia');
INSERT INTO `countries_en` VALUES (147, 'South Africa');
INSERT INTO `countries_en` VALUES (148, 'Antarctica');
INSERT INTO `countries_en` VALUES (149, 'South Georgia and the South Sandwich Islands');
INSERT INTO `countries_en` VALUES (150, 'Nauru');
INSERT INTO `countries_en` VALUES (151, 'Nepal');
INSERT INTO `countries_en` VALUES (152, 'Nicaragua');
INSERT INTO `countries_en` VALUES (153, 'Niger');
INSERT INTO `countries_en` VALUES (154, 'Nigeria');
INSERT INTO `countries_en` VALUES (155, 'Niue');
INSERT INTO `countries_en` VALUES (156, 'Norway');
INSERT INTO `countries_en` VALUES (157, 'Norfolk');
INSERT INTO `countries_en` VALUES (158, 'Palau Islands');
INSERT INTO `countries_en` VALUES (159, 'Pitcairn');
INSERT INTO `countries_en` VALUES (160, 'Portugal');
INSERT INTO `countries_en` VALUES (161, 'Georgia');
INSERT INTO `countries_en` VALUES (162, 'Japan');
INSERT INTO `countries_en` VALUES (163, 'Sweden');
INSERT INTO `countries_en` VALUES (164, 'Switzerland');
INSERT INTO `countries_en` VALUES (165, 'salvador');
INSERT INTO `countries_en` VALUES (166, 'Samoa');
INSERT INTO `countries_en` VALUES (167, 'Sierra Leone');
INSERT INTO `countries_en` VALUES (168, 'Senegal');
INSERT INTO `countries_en` VALUES (169, 'Cyprus');
INSERT INTO `countries_en` VALUES (170, 'Seychelles');
INSERT INTO `countries_en` VALUES (171, 'Saudi Arabia');
INSERT INTO `countries_en` VALUES (172, 'christmas island');
INSERT INTO `countries_en` VALUES (173, 'Sao Tome and Principe');
INSERT INTO `countries_en` VALUES (174, 'St. Helena');
INSERT INTO `countries_en` VALUES (175, 'Saint Kitts and Nevis');
INSERT INTO `countries_en` VALUES (176, 'Saint Lucia');
INSERT INTO `countries_en` VALUES (177, 'San Marino');
INSERT INTO `countries_en` VALUES (178, 'saint pierre and miklon islands');
INSERT INTO `countries_en` VALUES (179, 'Saint Vincent and the Grenadines');
INSERT INTO `countries_en` VALUES (180, 'Sri Lanka');
INSERT INTO `countries_en` VALUES (181, 'Slovakia');
INSERT INTO `countries_en` VALUES (182, 'Slovenia');
INSERT INTO `countries_en` VALUES (183, 'Svalbard and Jan Martin');
INSERT INTO `countries_en` VALUES (184, 'Swaziland');
INSERT INTO `countries_en` VALUES (185, 'Sudan');
INSERT INTO `countries_en` VALUES (186, 'Suriname');
INSERT INTO `countries_en` VALUES (187, 'Solomon Islands');
INSERT INTO `countries_en` VALUES (188, 'Somalia');
INSERT INTO `countries_en` VALUES (189, 'Tajikistan');
INSERT INTO `countries_en` VALUES (190, 'Thailand');
INSERT INTO `countries_en` VALUES (191, 'Tanzania');
INSERT INTO `countries_en` VALUES (192, 'Tonga');
INSERT INTO `countries_en` VALUES (193, 'Turks and Kectes Islands');
INSERT INTO `countries_en` VALUES (194, 'Trinidad and Tobago');
INSERT INTO `countries_en` VALUES (195, 'Tunisia');
INSERT INTO `countries_en` VALUES (196, 'Tuvalu');
INSERT INTO `countries_en` VALUES (197, 'Turkey');
INSERT INTO `countries_en` VALUES (198, 'Turkmenistan');
INSERT INTO `countries_en` VALUES (199, 'Tokelau');
INSERT INTO `countries_en` VALUES (200, 'wallis and fortuna');
INSERT INTO `countries_en` VALUES (201, 'Vanuatu');
INSERT INTO `countries_en` VALUES (202, 'Guatemala');
INSERT INTO `countries_en` VALUES (203, 'Virgin Islands, U.S.');
INSERT INTO `countries_en` VALUES (204, 'Virgin Islands, British');
INSERT INTO `countries_en` VALUES (205, 'Venezuela');
INSERT INTO `countries_en` VALUES (206, 'Brunei');
INSERT INTO `countries_en` VALUES (207, 'Uganda');
INSERT INTO `countries_en` VALUES (208, 'Ukraine');
INSERT INTO `countries_en` VALUES (209, 'Uruguay');
INSERT INTO `countries_en` VALUES (210, 'Uzbekistan');
INSERT INTO `countries_en` VALUES (211, 'Spain');
INSERT INTO `countries_en` VALUES (212, 'Greece');
INSERT INTO `countries_en` VALUES (213, 'Singapore');
INSERT INTO `countries_en` VALUES (214, 'new caledonia');
INSERT INTO `countries_en` VALUES (215, 'new Zealand');
INSERT INTO `countries_en` VALUES (216, 'Hungary');
INSERT INTO `countries_en` VALUES (217, 'Syria');
INSERT INTO `countries_en` VALUES (218, 'Jamaica');
INSERT INTO `countries_en` VALUES (219, 'Armenia');
INSERT INTO `countries_en` VALUES (220, 'Yemen');
INSERT INTO `countries_en` VALUES (221, 'Iraq');
INSERT INTO `countries_en` VALUES (222, 'Iran');
INSERT INTO `countries_en` VALUES (223, 'Israel');
INSERT INTO `countries_en` VALUES (224, 'Italy');
INSERT INTO `countries_en` VALUES (225, 'India');
INSERT INTO `countries_en` VALUES (226, 'Indonesia');
INSERT INTO `countries_en` VALUES (227, 'U.K.');
INSERT INTO `countries_en` VALUES (228, 'British Indian Ocean Territory');
INSERT INTO `countries_en` VALUES (229, 'Jordan');
INSERT INTO `countries_en` VALUES (230, 'Vietnam');
INSERT INTO `countries_en` VALUES (231, 'Zambia');
INSERT INTO `countries_en` VALUES (232, 'Jersey');
INSERT INTO `countries_en` VALUES (233, 'Chad');
INSERT INTO `countries_en` VALUES (234, 'Gibraltar');
INSERT INTO `countries_en` VALUES (235, 'Chile');
INSERT INTO `countries_en` VALUES (236, 'Central African Republic');

-- ----------------------------
-- Table structure for countries_zh-cn
-- ----------------------------
DROP TABLE IF EXISTS `countries_zh-cn`;
CREATE TABLE `countries_zh-cn`  (
  `id` smallint(1) UNSIGNED NOT NULL,
  `name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of countries_zh-cn
-- ----------------------------
INSERT INTO `countries_zh-cn` VALUES (1, '中国');
INSERT INTO `countries_zh-cn` VALUES (2, '阿尔巴尼亚');
INSERT INTO `countries_zh-cn` VALUES (3, '阿尔及利亚');
INSERT INTO `countries_zh-cn` VALUES (4, '阿富汗');
INSERT INTO `countries_zh-cn` VALUES (5, '阿根廷');
INSERT INTO `countries_zh-cn` VALUES (6, '阿拉伯联合酋长国');
INSERT INTO `countries_zh-cn` VALUES (7, '阿鲁巴');
INSERT INTO `countries_zh-cn` VALUES (8, '阿曼');
INSERT INTO `countries_zh-cn` VALUES (9, '阿塞拜疆');
INSERT INTO `countries_zh-cn` VALUES (10, '埃及');
INSERT INTO `countries_zh-cn` VALUES (11, '埃塞俄比亚');
INSERT INTO `countries_zh-cn` VALUES (12, '爱尔兰');
INSERT INTO `countries_zh-cn` VALUES (13, '爱沙尼亚');
INSERT INTO `countries_zh-cn` VALUES (14, '安道尔');
INSERT INTO `countries_zh-cn` VALUES (15, '安哥拉');
INSERT INTO `countries_zh-cn` VALUES (16, '安圭拉');
INSERT INTO `countries_zh-cn` VALUES (17, '安提瓜岛和巴布达');
INSERT INTO `countries_zh-cn` VALUES (18, '澳大利亚');
INSERT INTO `countries_zh-cn` VALUES (19, '奥地利');
INSERT INTO `countries_zh-cn` VALUES (20, '奥兰群岛');
INSERT INTO `countries_zh-cn` VALUES (21, '巴巴多斯岛');
INSERT INTO `countries_zh-cn` VALUES (22, '巴布亚新几内亚');
INSERT INTO `countries_zh-cn` VALUES (23, '巴哈马');
INSERT INTO `countries_zh-cn` VALUES (24, '巴基斯坦');
INSERT INTO `countries_zh-cn` VALUES (25, '巴拉圭');
INSERT INTO `countries_zh-cn` VALUES (26, '巴勒斯坦');
INSERT INTO `countries_zh-cn` VALUES (27, '巴林');
INSERT INTO `countries_zh-cn` VALUES (28, '巴拿马');
INSERT INTO `countries_zh-cn` VALUES (29, '巴西');
INSERT INTO `countries_zh-cn` VALUES (30, '白俄罗斯');
INSERT INTO `countries_zh-cn` VALUES (31, '百慕大');
INSERT INTO `countries_zh-cn` VALUES (32, '保加利亚');
INSERT INTO `countries_zh-cn` VALUES (33, '北马里亚纳群岛');
INSERT INTO `countries_zh-cn` VALUES (34, '贝宁');
INSERT INTO `countries_zh-cn` VALUES (35, '比利时');
INSERT INTO `countries_zh-cn` VALUES (36, '冰岛');
INSERT INTO `countries_zh-cn` VALUES (37, '波多黎各');
INSERT INTO `countries_zh-cn` VALUES (38, '波兰');
INSERT INTO `countries_zh-cn` VALUES (39, '玻利维亚');
INSERT INTO `countries_zh-cn` VALUES (40, '波斯尼亚和黑塞哥维那');
INSERT INTO `countries_zh-cn` VALUES (41, '博茨瓦纳');
INSERT INTO `countries_zh-cn` VALUES (42, '伯利兹');
INSERT INTO `countries_zh-cn` VALUES (43, '不丹');
INSERT INTO `countries_zh-cn` VALUES (44, '布基纳法索');
INSERT INTO `countries_zh-cn` VALUES (45, '布隆迪');
INSERT INTO `countries_zh-cn` VALUES (46, '布韦岛');
INSERT INTO `countries_zh-cn` VALUES (47, '朝鲜');
INSERT INTO `countries_zh-cn` VALUES (48, '丹麦');
INSERT INTO `countries_zh-cn` VALUES (49, '德国');
INSERT INTO `countries_zh-cn` VALUES (50, '东帝汶');
INSERT INTO `countries_zh-cn` VALUES (51, '多哥');
INSERT INTO `countries_zh-cn` VALUES (52, '多米尼加');
INSERT INTO `countries_zh-cn` VALUES (53, '多米尼加共和国');
INSERT INTO `countries_zh-cn` VALUES (54, '俄罗斯');
INSERT INTO `countries_zh-cn` VALUES (55, '厄瓜多尔');
INSERT INTO `countries_zh-cn` VALUES (56, '厄立特里亚');
INSERT INTO `countries_zh-cn` VALUES (57, '法国');
INSERT INTO `countries_zh-cn` VALUES (58, '法罗群岛');
INSERT INTO `countries_zh-cn` VALUES (59, '法属波利尼西亚');
INSERT INTO `countries_zh-cn` VALUES (60, '法属圭亚那');
INSERT INTO `countries_zh-cn` VALUES (61, '法属南部领地');
INSERT INTO `countries_zh-cn` VALUES (62, '梵蒂冈');
INSERT INTO `countries_zh-cn` VALUES (63, '菲律宾');
INSERT INTO `countries_zh-cn` VALUES (64, '斐济');
INSERT INTO `countries_zh-cn` VALUES (65, '芬兰');
INSERT INTO `countries_zh-cn` VALUES (66, '佛得角');
INSERT INTO `countries_zh-cn` VALUES (67, '弗兰克群岛');
INSERT INTO `countries_zh-cn` VALUES (68, '冈比亚');
INSERT INTO `countries_zh-cn` VALUES (69, '刚果');
INSERT INTO `countries_zh-cn` VALUES (70, '刚果民主共和国');
INSERT INTO `countries_zh-cn` VALUES (71, '哥伦比亚');
INSERT INTO `countries_zh-cn` VALUES (72, '哥斯达黎加');
INSERT INTO `countries_zh-cn` VALUES (73, '格恩西岛');
INSERT INTO `countries_zh-cn` VALUES (74, '格林纳达');
INSERT INTO `countries_zh-cn` VALUES (75, '格陵兰');
INSERT INTO `countries_zh-cn` VALUES (76, '古巴');
INSERT INTO `countries_zh-cn` VALUES (77, '瓜德罗普');
INSERT INTO `countries_zh-cn` VALUES (78, '关岛');
INSERT INTO `countries_zh-cn` VALUES (79, '圭亚那');
INSERT INTO `countries_zh-cn` VALUES (80, '哈萨克斯坦');
INSERT INTO `countries_zh-cn` VALUES (81, '海地');
INSERT INTO `countries_zh-cn` VALUES (82, '韩国');
INSERT INTO `countries_zh-cn` VALUES (83, '荷兰');
INSERT INTO `countries_zh-cn` VALUES (84, '赫德和麦克唐纳群岛');
INSERT INTO `countries_zh-cn` VALUES (85, '洪都拉斯');
INSERT INTO `countries_zh-cn` VALUES (86, '基里巴斯');
INSERT INTO `countries_zh-cn` VALUES (87, '吉布提');
INSERT INTO `countries_zh-cn` VALUES (88, '吉尔吉斯斯坦');
INSERT INTO `countries_zh-cn` VALUES (89, '几内亚');
INSERT INTO `countries_zh-cn` VALUES (90, '几内亚比绍');
INSERT INTO `countries_zh-cn` VALUES (91, '加拿大');
INSERT INTO `countries_zh-cn` VALUES (92, '加纳');
INSERT INTO `countries_zh-cn` VALUES (93, '加蓬');
INSERT INTO `countries_zh-cn` VALUES (94, '柬埔寨');
INSERT INTO `countries_zh-cn` VALUES (95, '捷克共和国');
INSERT INTO `countries_zh-cn` VALUES (96, '津巴布韦');
INSERT INTO `countries_zh-cn` VALUES (97, '喀麦隆');
INSERT INTO `countries_zh-cn` VALUES (98, '卡塔尔');
INSERT INTO `countries_zh-cn` VALUES (99, '开曼群岛');
INSERT INTO `countries_zh-cn` VALUES (100, '科科斯群岛');
INSERT INTO `countries_zh-cn` VALUES (101, '科摩罗');
INSERT INTO `countries_zh-cn` VALUES (102, '科特迪瓦');
INSERT INTO `countries_zh-cn` VALUES (103, '科威特');
INSERT INTO `countries_zh-cn` VALUES (104, '克罗地亚');
INSERT INTO `countries_zh-cn` VALUES (105, '肯尼亚');
INSERT INTO `countries_zh-cn` VALUES (106, '库克群岛');
INSERT INTO `countries_zh-cn` VALUES (107, '拉脱维亚');
INSERT INTO `countries_zh-cn` VALUES (108, '莱索托');
INSERT INTO `countries_zh-cn` VALUES (109, '老挝');
INSERT INTO `countries_zh-cn` VALUES (110, '黎巴嫩');
INSERT INTO `countries_zh-cn` VALUES (111, '利比里亚');
INSERT INTO `countries_zh-cn` VALUES (112, '利比亚');
INSERT INTO `countries_zh-cn` VALUES (113, '立陶宛');
INSERT INTO `countries_zh-cn` VALUES (114, '列支敦士登');
INSERT INTO `countries_zh-cn` VALUES (115, '留尼旺岛');
INSERT INTO `countries_zh-cn` VALUES (116, '卢森堡');
INSERT INTO `countries_zh-cn` VALUES (117, '卢旺达');
INSERT INTO `countries_zh-cn` VALUES (118, '罗马尼亚');
INSERT INTO `countries_zh-cn` VALUES (119, '马达加斯加');
INSERT INTO `countries_zh-cn` VALUES (120, '马尔代夫');
INSERT INTO `countries_zh-cn` VALUES (121, '马耳他');
INSERT INTO `countries_zh-cn` VALUES (122, '马拉维');
INSERT INTO `countries_zh-cn` VALUES (123, '马来西亚');
INSERT INTO `countries_zh-cn` VALUES (124, '马里');
INSERT INTO `countries_zh-cn` VALUES (125, '马其顿');
INSERT INTO `countries_zh-cn` VALUES (126, '马绍尔群岛');
INSERT INTO `countries_zh-cn` VALUES (127, '马提尼克');
INSERT INTO `countries_zh-cn` VALUES (128, '马约特岛');
INSERT INTO `countries_zh-cn` VALUES (129, '曼岛');
INSERT INTO `countries_zh-cn` VALUES (130, '毛里求斯');
INSERT INTO `countries_zh-cn` VALUES (131, '毛里塔尼亚');
INSERT INTO `countries_zh-cn` VALUES (132, '美国');
INSERT INTO `countries_zh-cn` VALUES (133, '美属萨摩亚');
INSERT INTO `countries_zh-cn` VALUES (134, '美属外岛');
INSERT INTO `countries_zh-cn` VALUES (135, '蒙古');
INSERT INTO `countries_zh-cn` VALUES (136, '蒙特塞拉特');
INSERT INTO `countries_zh-cn` VALUES (137, '孟加拉');
INSERT INTO `countries_zh-cn` VALUES (138, '密克罗尼西亚');
INSERT INTO `countries_zh-cn` VALUES (139, '秘鲁');
INSERT INTO `countries_zh-cn` VALUES (140, '缅甸');
INSERT INTO `countries_zh-cn` VALUES (141, '摩尔多瓦');
INSERT INTO `countries_zh-cn` VALUES (142, '摩洛哥');
INSERT INTO `countries_zh-cn` VALUES (143, '摩纳哥');
INSERT INTO `countries_zh-cn` VALUES (144, '莫桑比克');
INSERT INTO `countries_zh-cn` VALUES (145, '墨西哥');
INSERT INTO `countries_zh-cn` VALUES (146, '纳米比亚');
INSERT INTO `countries_zh-cn` VALUES (147, '南非');
INSERT INTO `countries_zh-cn` VALUES (148, '南极洲');
INSERT INTO `countries_zh-cn` VALUES (149, '南乔治亚和南桑德威奇群岛');
INSERT INTO `countries_zh-cn` VALUES (150, '瑙鲁');
INSERT INTO `countries_zh-cn` VALUES (151, '尼泊尔');
INSERT INTO `countries_zh-cn` VALUES (152, '尼加拉瓜');
INSERT INTO `countries_zh-cn` VALUES (153, '尼日尔');
INSERT INTO `countries_zh-cn` VALUES (154, '尼日利亚');
INSERT INTO `countries_zh-cn` VALUES (155, '纽埃');
INSERT INTO `countries_zh-cn` VALUES (156, '挪威');
INSERT INTO `countries_zh-cn` VALUES (157, '诺福克');
INSERT INTO `countries_zh-cn` VALUES (158, '帕劳群岛');
INSERT INTO `countries_zh-cn` VALUES (159, '皮特凯恩');
INSERT INTO `countries_zh-cn` VALUES (160, '葡萄牙');
INSERT INTO `countries_zh-cn` VALUES (161, '乔治亚');
INSERT INTO `countries_zh-cn` VALUES (162, '日本');
INSERT INTO `countries_zh-cn` VALUES (163, '瑞典');
INSERT INTO `countries_zh-cn` VALUES (164, '瑞士');
INSERT INTO `countries_zh-cn` VALUES (165, '萨尔瓦多');
INSERT INTO `countries_zh-cn` VALUES (166, '萨摩亚');
INSERT INTO `countries_zh-cn` VALUES (167, '塞拉利昂');
INSERT INTO `countries_zh-cn` VALUES (168, '塞内加尔');
INSERT INTO `countries_zh-cn` VALUES (169, '塞浦路斯');
INSERT INTO `countries_zh-cn` VALUES (170, '塞舌尔');
INSERT INTO `countries_zh-cn` VALUES (171, '沙特阿拉伯');
INSERT INTO `countries_zh-cn` VALUES (172, '圣诞岛');
INSERT INTO `countries_zh-cn` VALUES (173, '圣多美和普林西比');
INSERT INTO `countries_zh-cn` VALUES (174, '圣赫勒拿');
INSERT INTO `countries_zh-cn` VALUES (175, '圣基茨和尼维斯');
INSERT INTO `countries_zh-cn` VALUES (176, '圣卢西亚');
INSERT INTO `countries_zh-cn` VALUES (177, '圣马力诺');
INSERT INTO `countries_zh-cn` VALUES (178, '圣皮埃尔和米克隆群岛');
INSERT INTO `countries_zh-cn` VALUES (179, '圣文森特和格林纳丁斯');
INSERT INTO `countries_zh-cn` VALUES (180, '斯里兰卡');
INSERT INTO `countries_zh-cn` VALUES (181, '斯洛伐克');
INSERT INTO `countries_zh-cn` VALUES (182, '斯洛文尼亚');
INSERT INTO `countries_zh-cn` VALUES (183, '斯瓦尔巴和扬马廷');
INSERT INTO `countries_zh-cn` VALUES (184, '斯威士兰');
INSERT INTO `countries_zh-cn` VALUES (185, '苏丹');
INSERT INTO `countries_zh-cn` VALUES (186, '苏里南');
INSERT INTO `countries_zh-cn` VALUES (187, '所罗门群岛');
INSERT INTO `countries_zh-cn` VALUES (188, '索马里');
INSERT INTO `countries_zh-cn` VALUES (189, '塔吉克斯坦');
INSERT INTO `countries_zh-cn` VALUES (190, '泰国');
INSERT INTO `countries_zh-cn` VALUES (191, '坦桑尼亚');
INSERT INTO `countries_zh-cn` VALUES (192, '汤加');
INSERT INTO `countries_zh-cn` VALUES (193, '特克斯和凯克特斯群岛');
INSERT INTO `countries_zh-cn` VALUES (194, '特立尼达和多巴哥');
INSERT INTO `countries_zh-cn` VALUES (195, '突尼斯');
INSERT INTO `countries_zh-cn` VALUES (196, '图瓦卢');
INSERT INTO `countries_zh-cn` VALUES (197, '土耳其');
INSERT INTO `countries_zh-cn` VALUES (198, '土库曼斯坦');
INSERT INTO `countries_zh-cn` VALUES (199, '托克劳');
INSERT INTO `countries_zh-cn` VALUES (200, '瓦利斯和福图纳');
INSERT INTO `countries_zh-cn` VALUES (201, '瓦努阿图');
INSERT INTO `countries_zh-cn` VALUES (202, '危地马拉');
INSERT INTO `countries_zh-cn` VALUES (203, '维尔京群岛，美属');
INSERT INTO `countries_zh-cn` VALUES (204, '维尔京群岛，英属');
INSERT INTO `countries_zh-cn` VALUES (205, '委内瑞拉');
INSERT INTO `countries_zh-cn` VALUES (206, '文莱');
INSERT INTO `countries_zh-cn` VALUES (207, '乌干达');
INSERT INTO `countries_zh-cn` VALUES (208, '乌克兰');
INSERT INTO `countries_zh-cn` VALUES (209, '乌拉圭');
INSERT INTO `countries_zh-cn` VALUES (210, '乌兹别克斯坦');
INSERT INTO `countries_zh-cn` VALUES (211, '西班牙');
INSERT INTO `countries_zh-cn` VALUES (212, '希腊');
INSERT INTO `countries_zh-cn` VALUES (213, '新加坡');
INSERT INTO `countries_zh-cn` VALUES (214, '新喀里多尼亚');
INSERT INTO `countries_zh-cn` VALUES (215, '新西兰');
INSERT INTO `countries_zh-cn` VALUES (216, '匈牙利');
INSERT INTO `countries_zh-cn` VALUES (217, '叙利亚');
INSERT INTO `countries_zh-cn` VALUES (218, '牙买加');
INSERT INTO `countries_zh-cn` VALUES (219, '亚美尼亚');
INSERT INTO `countries_zh-cn` VALUES (220, '也门');
INSERT INTO `countries_zh-cn` VALUES (221, '伊拉克');
INSERT INTO `countries_zh-cn` VALUES (222, '伊朗');
INSERT INTO `countries_zh-cn` VALUES (223, '以色列');
INSERT INTO `countries_zh-cn` VALUES (224, '意大利');
INSERT INTO `countries_zh-cn` VALUES (225, '印度');
INSERT INTO `countries_zh-cn` VALUES (226, '印度尼西亚');
INSERT INTO `countries_zh-cn` VALUES (227, '英国');
INSERT INTO `countries_zh-cn` VALUES (228, '英属印度洋领地');
INSERT INTO `countries_zh-cn` VALUES (229, '约旦');
INSERT INTO `countries_zh-cn` VALUES (230, '越南');
INSERT INTO `countries_zh-cn` VALUES (231, '赞比亚');
INSERT INTO `countries_zh-cn` VALUES (232, '泽西岛');
INSERT INTO `countries_zh-cn` VALUES (233, '乍得');
INSERT INTO `countries_zh-cn` VALUES (234, '直布罗陀');
INSERT INTO `countries_zh-cn` VALUES (235, '智利');
INSERT INTO `countries_zh-cn` VALUES (236, '中非共和国');

-- ----------------------------
-- Table structure for currencies
-- ----------------------------
DROP TABLE IF EXISTS `currencies`;
CREATE TABLE `currencies`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `code` varchar(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `symbol` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `default` tinyint(1) NULL DEFAULT 0 COMMENT '是否默认币种',
  `rate` double(20, 10) NULL DEFAULT NULL COMMENT '对于default币种的汇率',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `code`(`code`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 159 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of currencies
-- ----------------------------
INSERT INTO `currencies` VALUES (1, 'AWG', 'ƒ', 0, NULL);
INSERT INTO `currencies` VALUES (2, 'AFN', '؋', 0, NULL);
INSERT INTO `currencies` VALUES (3, 'AOA', 'Kz', 0, NULL);
INSERT INTO `currencies` VALUES (4, 'XCD', 'EC$', 0, NULL);
INSERT INTO `currencies` VALUES (5, 'EUR', '€', 0, NULL);
INSERT INTO `currencies` VALUES (6, 'ALL', 'lek', 0, NULL);
INSERT INTO `currencies` VALUES (7, 'AED', '.د.ب', 0, NULL);
INSERT INTO `currencies` VALUES (8, 'ARS', '$', 0, NULL);
INSERT INTO `currencies` VALUES (9, 'AMD', '֏', 0, NULL);
INSERT INTO `currencies` VALUES (10, 'USD', '$', 1, NULL);
INSERT INTO `currencies` VALUES (11, 'AUD', '$', 0, NULL);
INSERT INTO `currencies` VALUES (12, 'AZN', 'ман', 0, NULL);
INSERT INTO `currencies` VALUES (13, 'BIF', 'FBu', 0, NULL);
INSERT INTO `currencies` VALUES (14, 'XOF', 'CFA', 0, NULL);
INSERT INTO `currencies` VALUES (15, 'BDT', 'Tk', 0, NULL);
INSERT INTO `currencies` VALUES (16, 'BGN', 'лв', 0, NULL);
INSERT INTO `currencies` VALUES (17, 'BHD', '.د.ب or BD', 0, NULL);
INSERT INTO `currencies` VALUES (18, 'BSD', 'B$', 0, NULL);
INSERT INTO `currencies` VALUES (19, 'BAM', 'KM', 0, NULL);
INSERT INTO `currencies` VALUES (20, 'BYR', 'р', 0, NULL);
INSERT INTO `currencies` VALUES (21, 'BZD', 'BZ$', 0, NULL);
INSERT INTO `currencies` VALUES (22, 'BMD', '$', 0, NULL);
INSERT INTO `currencies` VALUES (23, 'BOB', '$b', 0, NULL);
INSERT INTO `currencies` VALUES (24, 'BRL', 'R$', 0, NULL);
INSERT INTO `currencies` VALUES (25, 'BBD', '$', 0, NULL);
INSERT INTO `currencies` VALUES (26, 'BND', '$', 0, NULL);
INSERT INTO `currencies` VALUES (27, 'BTN', 'Nu.', 0, NULL);
INSERT INTO `currencies` VALUES (28, 'INR', '₹', 0, NULL);
INSERT INTO `currencies` VALUES (29, 'BWP', 'P', 0, NULL);
INSERT INTO `currencies` VALUES (30, 'CAD', '$', 0, NULL);
INSERT INTO `currencies` VALUES (31, 'CHF', 'CHF', 0, NULL);
INSERT INTO `currencies` VALUES (32, 'CLP', '$', 0, NULL);
INSERT INTO `currencies` VALUES (33, 'CNY', '¥', 0, NULL);
INSERT INTO `currencies` VALUES (34, 'NZD', '$', 0, NULL);
INSERT INTO `currencies` VALUES (35, 'COP', '$', 0, NULL);
INSERT INTO `currencies` VALUES (36, 'KMF', 'CF', 0, NULL);
INSERT INTO `currencies` VALUES (37, 'CVE', '$', 0, NULL);
INSERT INTO `currencies` VALUES (38, 'CRC', '₡', 0, NULL);
INSERT INTO `currencies` VALUES (39, 'CUC', '$', 0, NULL);
INSERT INTO `currencies` VALUES (40, 'CUP', '₱', 0, NULL);
INSERT INTO `currencies` VALUES (41, 'ANG', 'ƒ', 0, NULL);
INSERT INTO `currencies` VALUES (42, 'KYD', '$', 0, NULL);
INSERT INTO `currencies` VALUES (43, 'CZK', 'Kč', 0, NULL);
INSERT INTO `currencies` VALUES (44, 'DJF', 'fdj', 0, NULL);
INSERT INTO `currencies` VALUES (45, 'DKK', 'kr', 0, NULL);
INSERT INTO `currencies` VALUES (46, 'DOP', '$', 0, NULL);
INSERT INTO `currencies` VALUES (47, 'DZD', 'جد', 0, NULL);
INSERT INTO `currencies` VALUES (48, 'EGP', '£ ', 0, NULL);
INSERT INTO `currencies` VALUES (49, 'ERN', 'ናቕፋ', 0, NULL);
INSERT INTO `currencies` VALUES (50, 'MAD', 'م.د.', 0, NULL);
INSERT INTO `currencies` VALUES (51, 'MRO', 'UM', 0, NULL);
INSERT INTO `currencies` VALUES (52, 'ETB', 'Br', 0, NULL);
INSERT INTO `currencies` VALUES (53, 'FJD', '$', 0, NULL);
INSERT INTO `currencies` VALUES (54, 'FKP', '£', 0, NULL);
INSERT INTO `currencies` VALUES (55, 'GBP', '£', 0, NULL);
INSERT INTO `currencies` VALUES (56, 'GEL', 'ლ', 0, NULL);
INSERT INTO `currencies` VALUES (57, 'GHS', 'GH¢', 0, NULL);
INSERT INTO `currencies` VALUES (58, 'GIP', '£', 0, NULL);
INSERT INTO `currencies` VALUES (59, 'GNF', 'FG', 0, NULL);
INSERT INTO `currencies` VALUES (60, 'GMD', 'D', 0, NULL);
INSERT INTO `currencies` VALUES (61, 'GTQ', 'Q', 0, NULL);
INSERT INTO `currencies` VALUES (62, 'GYD', '$', 0, NULL);
INSERT INTO `currencies` VALUES (63, 'HKD', 'HK$', 0, NULL);
INSERT INTO `currencies` VALUES (64, 'HNL', 'L', 0, NULL);
INSERT INTO `currencies` VALUES (65, 'HRK', 'kn', 0, NULL);
INSERT INTO `currencies` VALUES (66, 'HTG', 'G', 0, NULL);
INSERT INTO `currencies` VALUES (67, 'HUF', 'Ft', 0, NULL);
INSERT INTO `currencies` VALUES (68, 'IDR', 'Rp', 0, NULL);
INSERT INTO `currencies` VALUES (69, 'IRR', '﷼', 0, NULL);
INSERT INTO `currencies` VALUES (70, 'IQD', 'ع.د', 0, NULL);
INSERT INTO `currencies` VALUES (71, 'ISK', 'kr', 0, NULL);
INSERT INTO `currencies` VALUES (72, 'ILS', '₪', 0, NULL);
INSERT INTO `currencies` VALUES (73, 'JMD', 'J$', 0, NULL);
INSERT INTO `currencies` VALUES (74, 'JOD', 'ا.د', 0, NULL);
INSERT INTO `currencies` VALUES (75, 'JPY', '¥', 0, NULL);
INSERT INTO `currencies` VALUES (76, 'KZT', '₸', 0, NULL);
INSERT INTO `currencies` VALUES (77, 'KES', 'KSh', 0, NULL);
INSERT INTO `currencies` VALUES (78, 'KGS', 'лв', 0, NULL);
INSERT INTO `currencies` VALUES (79, 'KHR', '៛', 0, NULL);
INSERT INTO `currencies` VALUES (80, 'KRW', '₩', 0, NULL);
INSERT INTO `currencies` VALUES (81, 'KWD', 'ك', 0, NULL);
INSERT INTO `currencies` VALUES (82, 'LAK', '₭', 0, NULL);
INSERT INTO `currencies` VALUES (83, 'LBP', 'ل.ل', 0, NULL);
INSERT INTO `currencies` VALUES (84, 'LRD', '$', 0, NULL);
INSERT INTO `currencies` VALUES (85, 'LYD', ' د.ل', 0, NULL);
INSERT INTO `currencies` VALUES (86, 'LKR', 'Rs', 0, NULL);
INSERT INTO `currencies` VALUES (87, 'LSL', 'L or M', 0, NULL);
INSERT INTO `currencies` VALUES (88, 'ZAR', 'R', 0, NULL);
INSERT INTO `currencies` VALUES (89, 'MOP', 'MOP$', 0, NULL);
INSERT INTO `currencies` VALUES (90, 'MDL', 'L', 0, NULL);
INSERT INTO `currencies` VALUES (91, 'MGA', 'Ar', 0, NULL);
INSERT INTO `currencies` VALUES (92, 'MVR', 'rf', 0, NULL);
INSERT INTO `currencies` VALUES (93, 'MXN', '$', 0, NULL);
INSERT INTO `currencies` VALUES (94, 'MKD', 'ден', 0, NULL);
INSERT INTO `currencies` VALUES (95, 'MMK', 'K', 0, NULL);
INSERT INTO `currencies` VALUES (96, 'MNT', '₮', 0, NULL);
INSERT INTO `currencies` VALUES (97, 'MZN', 'MT', 0, NULL);
INSERT INTO `currencies` VALUES (98, 'MUR', 'Rs', 0, NULL);
INSERT INTO `currencies` VALUES (99, 'MWK', 'MK', 0, NULL);
INSERT INTO `currencies` VALUES (100, 'MYR', 'RM', 0, NULL);
INSERT INTO `currencies` VALUES (101, 'NAD', '$', 0, NULL);
INSERT INTO `currencies` VALUES (102, 'XPF', '₣', 0, NULL);
INSERT INTO `currencies` VALUES (103, 'NGN', '₦', 0, NULL);
INSERT INTO `currencies` VALUES (104, 'NIO', 'C$', 0, NULL);
INSERT INTO `currencies` VALUES (105, 'NOK', 'kr', 0, NULL);
INSERT INTO `currencies` VALUES (106, 'NPR', 'Rs', 0, NULL);
INSERT INTO `currencies` VALUES (107, 'OMR', 'ع.ر.', 0, NULL);
INSERT INTO `currencies` VALUES (108, 'PKR', 'Rs', 0, NULL);
INSERT INTO `currencies` VALUES (109, 'PAB', 'B/', 0, NULL);
INSERT INTO `currencies` VALUES (110, 'PEN', 'S/', 0, NULL);
INSERT INTO `currencies` VALUES (111, 'PHP', '₱', 0, NULL);
INSERT INTO `currencies` VALUES (112, 'PGK', 'K', 0, NULL);
INSERT INTO `currencies` VALUES (113, 'PLN', 'zł', 0, NULL);
INSERT INTO `currencies` VALUES (114, 'KPW', '₩', 0, NULL);
INSERT INTO `currencies` VALUES (115, 'PYG', '₲', 0, NULL);
INSERT INTO `currencies` VALUES (116, 'QAR', 'ق.ر ', 0, NULL);
INSERT INTO `currencies` VALUES (117, 'RON', 'lei', 0, NULL);
INSERT INTO `currencies` VALUES (118, 'RUB', '₽', 0, NULL);
INSERT INTO `currencies` VALUES (119, 'RWF', 'FRw, RF, R₣', 0, NULL);
INSERT INTO `currencies` VALUES (120, 'SAR', 'ر.س', 0, NULL);
INSERT INTO `currencies` VALUES (121, 'SDG', '.س.ج', 0, NULL);
INSERT INTO `currencies` VALUES (122, 'SGD', '$', 0, NULL);
INSERT INTO `currencies` VALUES (123, 'SBD', 'SI$', 0, NULL);
INSERT INTO `currencies` VALUES (124, 'SLL', 'Le', 0, NULL);
INSERT INTO `currencies` VALUES (125, 'SOS', 'S', 0, NULL);
INSERT INTO `currencies` VALUES (126, 'RSD', 'РСД', 0, NULL);
INSERT INTO `currencies` VALUES (127, 'SSP', '£', 0, NULL);
INSERT INTO `currencies` VALUES (128, 'SRD', '$', 0, NULL);
INSERT INTO `currencies` VALUES (129, 'SEK', 'kr', 0, NULL);
INSERT INTO `currencies` VALUES (130, 'SZL', 'L or E', 0, NULL);
INSERT INTO `currencies` VALUES (131, 'SCR', 'Rs', 0, NULL);
INSERT INTO `currencies` VALUES (132, 'SYP', '£', 0, NULL);
INSERT INTO `currencies` VALUES (133, 'THB', '฿', 0, NULL);
INSERT INTO `currencies` VALUES (134, 'TJS', 'SM', 0, NULL);
INSERT INTO `currencies` VALUES (135, 'TMT', 'T', 0, NULL);
INSERT INTO `currencies` VALUES (136, 'TOP', 'T$', 0, NULL);
INSERT INTO `currencies` VALUES (137, 'TTD', 'TT$', 0, NULL);
INSERT INTO `currencies` VALUES (138, 'TND', 'ت.د', 0, NULL);
INSERT INTO `currencies` VALUES (139, 'TRY', '₺', 0, NULL);
INSERT INTO `currencies` VALUES (140, 'TWD', 'NT$', 0, NULL);
INSERT INTO `currencies` VALUES (141, 'TZS', 'Sh', 0, NULL);
INSERT INTO `currencies` VALUES (142, 'UGX', 'USh', 0, NULL);
INSERT INTO `currencies` VALUES (143, 'UAH', '₴', 0, NULL);
INSERT INTO `currencies` VALUES (144, 'UYU', '$U', 0, NULL);
INSERT INTO `currencies` VALUES (145, 'UZS', 'лв', 0, NULL);
INSERT INTO `currencies` VALUES (146, 'VEF', 'Bs', 0, NULL);
INSERT INTO `currencies` VALUES (147, 'VND', '₫', 0, NULL);
INSERT INTO `currencies` VALUES (148, 'VUV', 'VT', 0, NULL);
INSERT INTO `currencies` VALUES (149, 'WST', '$', 0, NULL);
INSERT INTO `currencies` VALUES (150, 'YER', '﷼', 0, NULL);
INSERT INTO `currencies` VALUES (151, 'ZMW', 'ZMK', 0, NULL);
INSERT INTO `currencies` VALUES (152, 'ZWL', '$', 0, NULL);
INSERT INTO `currencies` VALUES (153, 'XAF', 'FCFA', 0, NULL);
INSERT INTO `currencies` VALUES (154, 'STD', 'Db', 0, NULL);
INSERT INTO `currencies` VALUES (155, 'SHP', '£', 0, NULL);
INSERT INTO `currencies` VALUES (156, 'CDF', 'FC', 0, NULL);
INSERT INTO `currencies` VALUES (157, 'BYN', 'Br', 0, NULL);
INSERT INTO `currencies` VALUES (158, 'AAD', '$', 0, NULL);

-- ----------------------------
-- Table structure for currencies_en
-- ----------------------------
DROP TABLE IF EXISTS `currencies_en`;
CREATE TABLE `currencies_en`  (
  `id` smallint(1) NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of currencies_en
-- ----------------------------
INSERT INTO `currencies_en` VALUES (1, 'Arubin florin');
INSERT INTO `currencies_en` VALUES (2, 'Afghan Afghani');
INSERT INTO `currencies_en` VALUES (3, 'Angolan Kwanza');
INSERT INTO `currencies_en` VALUES (4, 'East Caribbean dollar');
INSERT INTO `currencies_en` VALUES (5, 'Euro');
INSERT INTO `currencies_en` VALUES (6, 'Albanian lek');
INSERT INTO `currencies_en` VALUES (7, 'Emirati Dirham');
INSERT INTO `currencies_en` VALUES (8, 'Argentine peso');
INSERT INTO `currencies_en` VALUES (9, 'Armenian dram');
INSERT INTO `currencies_en` VALUES (10, 'US Dollar');
INSERT INTO `currencies_en` VALUES (11, 'Australian Dollar');
INSERT INTO `currencies_en` VALUES (12, 'Azerbaijani manat');
INSERT INTO `currencies_en` VALUES (13, 'Burundian Franc');
INSERT INTO `currencies_en` VALUES (14, 'CFA Franc');
INSERT INTO `currencies_en` VALUES (15, 'Bangladeshi Taka');
INSERT INTO `currencies_en` VALUES (16, 'Bulgarian lev');
INSERT INTO `currencies_en` VALUES (17, 'Bahraini Dinar');
INSERT INTO `currencies_en` VALUES (18, 'Bahamian dollar');
INSERT INTO `currencies_en` VALUES (19, 'Bosnian Convertible Marka');
INSERT INTO `currencies_en` VALUES (20, 'Belarusian ruble');
INSERT INTO `currencies_en` VALUES (21, 'Belize dollar');
INSERT INTO `currencies_en` VALUES (22, 'Bermudian dollar');
INSERT INTO `currencies_en` VALUES (23, 'Bolivian Boliviano');
INSERT INTO `currencies_en` VALUES (24, 'Brazilian real');
INSERT INTO `currencies_en` VALUES (25, 'Barbadian dollar');
INSERT INTO `currencies_en` VALUES (26, 'Bruneian Dollar');
INSERT INTO `currencies_en` VALUES (27, 'Bhutanese Ngultrum');
INSERT INTO `currencies_en` VALUES (28, 'Indian Rupee');
INSERT INTO `currencies_en` VALUES (29, 'Botswana Pula');
INSERT INTO `currencies_en` VALUES (30, 'Canadian Dollar');
INSERT INTO `currencies_en` VALUES (31, 'Swiss Franc');
INSERT INTO `currencies_en` VALUES (32, 'Chilean Peso');
INSERT INTO `currencies_en` VALUES (33, 'Renminbi');
INSERT INTO `currencies_en` VALUES (34, 'New Zealand Dollar');
INSERT INTO `currencies_en` VALUES (35, 'Colombian peso');
INSERT INTO `currencies_en` VALUES (36, 'Comoran Franc');
INSERT INTO `currencies_en` VALUES (37, 'Cape Verdean Escudo');
INSERT INTO `currencies_en` VALUES (38, 'Costa Rican colón');
INSERT INTO `currencies_en` VALUES (39, 'Cuban convertible peso');
INSERT INTO `currencies_en` VALUES (40, 'Cuban peso');
INSERT INTO `currencies_en` VALUES (41, 'Dutch Guilder');
INSERT INTO `currencies_en` VALUES (42, 'Caymanian Dollar');
INSERT INTO `currencies_en` VALUES (43, 'Czech koruna');
INSERT INTO `currencies_en` VALUES (44, 'Djiboutian Franc');
INSERT INTO `currencies_en` VALUES (45, 'Danish krone');
INSERT INTO `currencies_en` VALUES (46, 'Dominican peso');
INSERT INTO `currencies_en` VALUES (47, 'Algerian Dinar');
INSERT INTO `currencies_en` VALUES (48, 'Egyptian Pound');
INSERT INTO `currencies_en` VALUES (49, 'Eritrean nakfa');
INSERT INTO `currencies_en` VALUES (50, 'Moroccan Dirham');
INSERT INTO `currencies_en` VALUES (51, 'Mauritanian Ouguiya');
INSERT INTO `currencies_en` VALUES (52, 'Ethiopian Birr');
INSERT INTO `currencies_en` VALUES (53, 'Fijian dollar');
INSERT INTO `currencies_en` VALUES (54, 'Falkland Island Pound');
INSERT INTO `currencies_en` VALUES (55, 'British Pound');
INSERT INTO `currencies_en` VALUES (56, 'Georgian lari');
INSERT INTO `currencies_en` VALUES (57, 'Ghanaian Cedi');
INSERT INTO `currencies_en` VALUES (58, 'Gibraltar pound');
INSERT INTO `currencies_en` VALUES (59, 'Guinean Franc');
INSERT INTO `currencies_en` VALUES (60, 'Gambian dalasi');
INSERT INTO `currencies_en` VALUES (61, 'Guatemalan Quetzal');
INSERT INTO `currencies_en` VALUES (62, 'Guyanese dollar');
INSERT INTO `currencies_en` VALUES (63, 'Hong Kong dollar');
INSERT INTO `currencies_en` VALUES (64, 'Honduran lempira');
INSERT INTO `currencies_en` VALUES (65, 'Croatian kuna');
INSERT INTO `currencies_en` VALUES (66, 'Haitian gourde');
INSERT INTO `currencies_en` VALUES (67, 'Hungarian forint');
INSERT INTO `currencies_en` VALUES (68, 'Indonesian Rupiah');
INSERT INTO `currencies_en` VALUES (69, 'Iranian Rial');
INSERT INTO `currencies_en` VALUES (70, 'Iraqi Dinar');
INSERT INTO `currencies_en` VALUES (71, 'Icelandic Krona');
INSERT INTO `currencies_en` VALUES (72, 'Israeli Shekel');
INSERT INTO `currencies_en` VALUES (73, 'Jamaican dollar');
INSERT INTO `currencies_en` VALUES (74, 'Jordanian Dinar');
INSERT INTO `currencies_en` VALUES (75, 'Japanese yen');
INSERT INTO `currencies_en` VALUES (76, 'Kazakhstani tenge');
INSERT INTO `currencies_en` VALUES (77, 'Kenyan Shilling');
INSERT INTO `currencies_en` VALUES (78, 'Kyrgyzstani som');
INSERT INTO `currencies_en` VALUES (79, 'Cambodian Riel');
INSERT INTO `currencies_en` VALUES (80, 'South Korean won');
INSERT INTO `currencies_en` VALUES (81, 'Kuwaiti Dinar');
INSERT INTO `currencies_en` VALUES (82, 'Lao or Laotian Kip');
INSERT INTO `currencies_en` VALUES (83, 'Lebanese Pound');
INSERT INTO `currencies_en` VALUES (84, 'Liberian Dollar');
INSERT INTO `currencies_en` VALUES (85, 'Libyan Dinar');
INSERT INTO `currencies_en` VALUES (86, 'Sri Lankan Rupee');
INSERT INTO `currencies_en` VALUES (87, 'Lesotho loti');
INSERT INTO `currencies_en` VALUES (88, 'South African Rand');
INSERT INTO `currencies_en` VALUES (89, 'Macau Pataca');
INSERT INTO `currencies_en` VALUES (90, 'Moldovan Leu');
INSERT INTO `currencies_en` VALUES (91, 'Malagasy Ariary');
INSERT INTO `currencies_en` VALUES (92, 'Maldivian Rufiyaa');
INSERT INTO `currencies_en` VALUES (93, 'Mexico Peso');
INSERT INTO `currencies_en` VALUES (94, 'Macedonian Denar');
INSERT INTO `currencies_en` VALUES (95, 'Burmese Kyat');
INSERT INTO `currencies_en` VALUES (96, 'Mongolian Tughrik');
INSERT INTO `currencies_en` VALUES (97, 'Mozambican Metical');
INSERT INTO `currencies_en` VALUES (98, 'Mauritian rupee');
INSERT INTO `currencies_en` VALUES (99, 'Malawian Kwacha');
INSERT INTO `currencies_en` VALUES (100, 'Malaysian Ringgit');
INSERT INTO `currencies_en` VALUES (101, 'Namibian Dollar');
INSERT INTO `currencies_en` VALUES (102, 'CFP Franc');
INSERT INTO `currencies_en` VALUES (103, 'Nigerian Naira');
INSERT INTO `currencies_en` VALUES (104, 'Nicaraguan córdoba');
INSERT INTO `currencies_en` VALUES (105, 'Norwegian krone');
INSERT INTO `currencies_en` VALUES (106, 'Nepalese Rupee');
INSERT INTO `currencies_en` VALUES (107, 'Omani Rial');
INSERT INTO `currencies_en` VALUES (108, 'Pakistani Rupee');
INSERT INTO `currencies_en` VALUES (109, 'Balboa panamérn');
INSERT INTO `currencies_en` VALUES (110, 'Peruvian nuevo sol');
INSERT INTO `currencies_en` VALUES (111, 'Philippine Peso');
INSERT INTO `currencies_en` VALUES (112, 'Papua New Guinean Kina');
INSERT INTO `currencies_en` VALUES (113, 'Polish złoty');
INSERT INTO `currencies_en` VALUES (114, 'North Korean won');
INSERT INTO `currencies_en` VALUES (115, 'Paraguayan guarani');
INSERT INTO `currencies_en` VALUES (116, 'Qatari Riyal');
INSERT INTO `currencies_en` VALUES (117, 'Romanian leu');
INSERT INTO `currencies_en` VALUES (118, 'Russian Rouble');
INSERT INTO `currencies_en` VALUES (119, 'Rwandan franc');
INSERT INTO `currencies_en` VALUES (120, 'Saudi Arabian Riyal');
INSERT INTO `currencies_en` VALUES (121, 'Sudanese Pound');
INSERT INTO `currencies_en` VALUES (122, 'Singapore Dollar');
INSERT INTO `currencies_en` VALUES (123, 'Solomon Islander Dollar');
INSERT INTO `currencies_en` VALUES (124, 'Sierra Leonean Leone');
INSERT INTO `currencies_en` VALUES (125, 'Somali Shilling');
INSERT INTO `currencies_en` VALUES (126, 'Serbian Dinar');
INSERT INTO `currencies_en` VALUES (127, 'South Sudanese pound');
INSERT INTO `currencies_en` VALUES (128, 'Surinamese dollar');
INSERT INTO `currencies_en` VALUES (129, 'Swedish krona');
INSERT INTO `currencies_en` VALUES (130, 'Swazi Lilangeni');
INSERT INTO `currencies_en` VALUES (131, 'Seychellois Rupee');
INSERT INTO `currencies_en` VALUES (132, 'Syrian Pound');
INSERT INTO `currencies_en` VALUES (133, 'Thai Baht');
INSERT INTO `currencies_en` VALUES (134, 'Tajikistani somoni');
INSERT INTO `currencies_en` VALUES (135, 'Turkmenistan manat');
INSERT INTO `currencies_en` VALUES (136, 'Tongan Pa\'anga');
INSERT INTO `currencies_en` VALUES (137, 'Trinidadian dollar');
INSERT INTO `currencies_en` VALUES (138, 'Tunisian Dinar');
INSERT INTO `currencies_en` VALUES (139, 'Turkish Lira');
INSERT INTO `currencies_en` VALUES (140, 'Taiwan New Dollar');
INSERT INTO `currencies_en` VALUES (141, 'Tanzanian Shilling');
INSERT INTO `currencies_en` VALUES (142, 'Ugandan Shilling');
INSERT INTO `currencies_en` VALUES (143, 'Ukrainian Hryvnia');
INSERT INTO `currencies_en` VALUES (144, 'Uruguayan peso');
INSERT INTO `currencies_en` VALUES (145, 'Uzbekistani som');
INSERT INTO `currencies_en` VALUES (146, 'Venezuelan bolivar');
INSERT INTO `currencies_en` VALUES (147, 'Vietnamese Dong');
INSERT INTO `currencies_en` VALUES (148, 'Ni-Vanuatu Vatu');
INSERT INTO `currencies_en` VALUES (149, 'Samoan Tālā');
INSERT INTO `currencies_en` VALUES (150, 'Yemeni Rial');
INSERT INTO `currencies_en` VALUES (151, 'Zambian Kwacha');
INSERT INTO `currencies_en` VALUES (152, 'Zimbabwe Dollar');
INSERT INTO `currencies_en` VALUES (153, 'Central African CFA franc');
INSERT INTO `currencies_en` VALUES (154, 'Dobra');
INSERT INTO `currencies_en` VALUES (155, 'Saint Helena pound');
INSERT INTO `currencies_en` VALUES (156, 'Congolese Franc');
INSERT INTO `currencies_en` VALUES (157, 'Belarusian ruble');
INSERT INTO `currencies_en` VALUES (158, 'Antarctican dollar');

-- ----------------------------
-- Table structure for currencies_zh-cn
-- ----------------------------
DROP TABLE IF EXISTS `currencies_zh-cn`;
CREATE TABLE `currencies_zh-cn`  (
  `id` smallint(1) NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of currencies_zh-cn
-- ----------------------------
INSERT INTO `currencies_zh-cn` VALUES (1, '阿鲁宾弗洛林');
INSERT INTO `currencies_zh-cn` VALUES (2, '阿富汗阿富汗尼');
INSERT INTO `currencies_zh-cn` VALUES (3, '安哥拉宽扎');
INSERT INTO `currencies_zh-cn` VALUES (4, '东加勒比元');
INSERT INTO `currencies_zh-cn` VALUES (5, '欧元');
INSERT INTO `currencies_zh-cn` VALUES (6, '阿尔巴尼亚列克');
INSERT INTO `currencies_zh-cn` VALUES (7, '阿联酋迪拉姆');
INSERT INTO `currencies_zh-cn` VALUES (8, '阿根廷比索');
INSERT INTO `currencies_zh-cn` VALUES (9, '亚美尼亚德拉姆');
INSERT INTO `currencies_zh-cn` VALUES (10, '美元');
INSERT INTO `currencies_zh-cn` VALUES (11, '澳元');
INSERT INTO `currencies_zh-cn` VALUES (12, '阿塞拜疆马纳特');
INSERT INTO `currencies_zh-cn` VALUES (13, '布隆迪法郎');
INSERT INTO `currencies_zh-cn` VALUES (14, '非洲金融共同体法郎');
INSERT INTO `currencies_zh-cn` VALUES (15, '孟加拉塔卡');
INSERT INTO `currencies_zh-cn` VALUES (16, '保加利亚列弗');
INSERT INTO `currencies_zh-cn` VALUES (17, '巴林第纳尔');
INSERT INTO `currencies_zh-cn` VALUES (18, '巴哈马元');
INSERT INTO `currencies_zh-cn` VALUES (19, '波斯尼亚敞篷马卡');
INSERT INTO `currencies_zh-cn` VALUES (20, '白俄罗斯卢布');
INSERT INTO `currencies_zh-cn` VALUES (21, '伯利兹元');
INSERT INTO `currencies_zh-cn` VALUES (22, '百慕大元');
INSERT INTO `currencies_zh-cn` VALUES (23, '玻利维亚玻利维亚诺');
INSERT INTO `currencies_zh-cn` VALUES (24, '巴西雷亚尔');
INSERT INTO `currencies_zh-cn` VALUES (25, '巴巴多斯元');
INSERT INTO `currencies_zh-cn` VALUES (26, '文莱元');
INSERT INTO `currencies_zh-cn` VALUES (27, '不丹努尔特鲁姆');
INSERT INTO `currencies_zh-cn` VALUES (28, '印度卢比');
INSERT INTO `currencies_zh-cn` VALUES (29, '博茨瓦纳普拉');
INSERT INTO `currencies_zh-cn` VALUES (30, '加元');
INSERT INTO `currencies_zh-cn` VALUES (31, '瑞士法郎');
INSERT INTO `currencies_zh-cn` VALUES (32, '智利比索');
INSERT INTO `currencies_zh-cn` VALUES (33, '人民币');
INSERT INTO `currencies_zh-cn` VALUES (34, '新西兰元');
INSERT INTO `currencies_zh-cn` VALUES (35, '哥伦比亚比索');
INSERT INTO `currencies_zh-cn` VALUES (36, '科摩罗法郎');
INSERT INTO `currencies_zh-cn` VALUES (37, '佛得角埃斯库多');
INSERT INTO `currencies_zh-cn` VALUES (38, '哥斯达黎加结肠');
INSERT INTO `currencies_zh-cn` VALUES (39, '古巴可兑换比索');
INSERT INTO `currencies_zh-cn` VALUES (40, '古巴比索');
INSERT INTO `currencies_zh-cn` VALUES (41, '荷兰盾');
INSERT INTO `currencies_zh-cn` VALUES (42, '开曼元');
INSERT INTO `currencies_zh-cn` VALUES (43, '捷克克朗');
INSERT INTO `currencies_zh-cn` VALUES (44, '吉布提法郎');
INSERT INTO `currencies_zh-cn` VALUES (45, '丹麦克朗');
INSERT INTO `currencies_zh-cn` VALUES (46, '多米尼加比索');
INSERT INTO `currencies_zh-cn` VALUES (47, '阿尔及利亚第纳尔');
INSERT INTO `currencies_zh-cn` VALUES (48, '埃及镑');
INSERT INTO `currencies_zh-cn` VALUES (49, '厄立特里亚纳克法');
INSERT INTO `currencies_zh-cn` VALUES (50, '摩洛哥迪拉姆');
INSERT INTO `currencies_zh-cn` VALUES (51, '毛里塔尼亚乌吉亚');
INSERT INTO `currencies_zh-cn` VALUES (52, '埃塞俄比亚比尔');
INSERT INTO `currencies_zh-cn` VALUES (53, '斐济元');
INSERT INTO `currencies_zh-cn` VALUES (54, '福克兰群岛镑');
INSERT INTO `currencies_zh-cn` VALUES (55, '英镑');
INSERT INTO `currencies_zh-cn` VALUES (56, '格鲁吉亚拉里');
INSERT INTO `currencies_zh-cn` VALUES (57, '加纳塞地');
INSERT INTO `currencies_zh-cn` VALUES (58, '直布罗陀镑');
INSERT INTO `currencies_zh-cn` VALUES (59, '几内亚法郎');
INSERT INTO `currencies_zh-cn` VALUES (60, '冈比亚达拉西');
INSERT INTO `currencies_zh-cn` VALUES (61, '危地马拉格查尔');
INSERT INTO `currencies_zh-cn` VALUES (62, '圭亚那元');
INSERT INTO `currencies_zh-cn` VALUES (63, '港币');
INSERT INTO `currencies_zh-cn` VALUES (64, '洪都拉斯伦皮拉');
INSERT INTO `currencies_zh-cn` VALUES (65, '克罗地亚库纳');
INSERT INTO `currencies_zh-cn` VALUES (66, '海天古德');
INSERT INTO `currencies_zh-cn` VALUES (67, '匈牙利福林');
INSERT INTO `currencies_zh-cn` VALUES (68, '印尼盾');
INSERT INTO `currencies_zh-cn` VALUES (69, '伊朗里亚尔');
INSERT INTO `currencies_zh-cn` VALUES (70, '伊拉克第纳尔');
INSERT INTO `currencies_zh-cn` VALUES (71, '冰岛克朗');
INSERT INTO `currencies_zh-cn` VALUES (72, '以色列谢克尔');
INSERT INTO `currencies_zh-cn` VALUES (73, '牙买加元');
INSERT INTO `currencies_zh-cn` VALUES (74, '约旦第纳尔');
INSERT INTO `currencies_zh-cn` VALUES (75, '日圆');
INSERT INTO `currencies_zh-cn` VALUES (76, '哈萨克斯坦坚戈');
INSERT INTO `currencies_zh-cn` VALUES (77, '肯尼亚先令');
INSERT INTO `currencies_zh-cn` VALUES (78, '吉尔吉斯斯坦索姆');
INSERT INTO `currencies_zh-cn` VALUES (79, '柬埔寨瑞尔');
INSERT INTO `currencies_zh-cn` VALUES (80, '韩元');
INSERT INTO `currencies_zh-cn` VALUES (81, '科威特第纳尔');
INSERT INTO `currencies_zh-cn` VALUES (82, '老挝或老挝基普');
INSERT INTO `currencies_zh-cn` VALUES (83, '黎巴嫩镑');
INSERT INTO `currencies_zh-cn` VALUES (84, '利比里亚元');
INSERT INTO `currencies_zh-cn` VALUES (85, '利比亚第纳尔');
INSERT INTO `currencies_zh-cn` VALUES (86, '斯里兰卡卢比');
INSERT INTO `currencies_zh-cn` VALUES (87, '莱索托洛蒂');
INSERT INTO `currencies_zh-cn` VALUES (88, '南非兰特');
INSERT INTO `currencies_zh-cn` VALUES (89, '澳门元');
INSERT INTO `currencies_zh-cn` VALUES (90, '摩尔多瓦列伊');
INSERT INTO `currencies_zh-cn` VALUES (91, '马达加斯加阿里亚里');
INSERT INTO `currencies_zh-cn` VALUES (92, '马尔代夫拉菲亚');
INSERT INTO `currencies_zh-cn` VALUES (93, '墨西哥比索');
INSERT INTO `currencies_zh-cn` VALUES (94, '马其顿代纳尔');
INSERT INTO `currencies_zh-cn` VALUES (95, '缅元');
INSERT INTO `currencies_zh-cn` VALUES (96, '蒙古图格里克语');
INSERT INTO `currencies_zh-cn` VALUES (97, '莫桑比克梅蒂卡尔');
INSERT INTO `currencies_zh-cn` VALUES (98, '毛里求斯卢比');
INSERT INTO `currencies_zh-cn` VALUES (99, '马拉维克瓦查');
INSERT INTO `currencies_zh-cn` VALUES (100, '马来西亚令吉');
INSERT INTO `currencies_zh-cn` VALUES (101, '纳米比亚元');
INSERT INTO `currencies_zh-cn` VALUES (102, '太平洋法郎');
INSERT INTO `currencies_zh-cn` VALUES (103, '尼日利亚奈拉');
INSERT INTO `currencies_zh-cn` VALUES (104, '尼加拉瓜科尔多瓦');
INSERT INTO `currencies_zh-cn` VALUES (105, '挪威克朗');
INSERT INTO `currencies_zh-cn` VALUES (106, '尼泊尔卢比');
INSERT INTO `currencies_zh-cn` VALUES (107, '阿曼里亚尔');
INSERT INTO `currencies_zh-cn` VALUES (108, '巴基斯坦卢比');
INSERT INTO `currencies_zh-cn` VALUES (109, '巴尔博亚巴拿马');
INSERT INTO `currencies_zh-cn` VALUES (110, '秘鲁新索尔');
INSERT INTO `currencies_zh-cn` VALUES (111, '菲律宾比索');
INSERT INTO `currencies_zh-cn` VALUES (112, '巴布亚新几内亚基纳');
INSERT INTO `currencies_zh-cn` VALUES (113, '波兰兹罗提');
INSERT INTO `currencies_zh-cn` VALUES (114, '朝鲜元');
INSERT INTO `currencies_zh-cn` VALUES (115, '巴拉圭瓜拉尼');
INSERT INTO `currencies_zh-cn` VALUES (116, '卡塔尔里亚尔');
INSERT INTO `currencies_zh-cn` VALUES (117, '罗马尼亚列伊');
INSERT INTO `currencies_zh-cn` VALUES (118, '俄罗斯卢布');
INSERT INTO `currencies_zh-cn` VALUES (119, '卢旺达法郎');
INSERT INTO `currencies_zh-cn` VALUES (120, '沙特阿拉伯里亚尔');
INSERT INTO `currencies_zh-cn` VALUES (121, '苏丹镑');
INSERT INTO `currencies_zh-cn` VALUES (122, '新加坡元');
INSERT INTO `currencies_zh-cn` VALUES (123, '所罗门群岛元');
INSERT INTO `currencies_zh-cn` VALUES (124, '塞拉利昂利昂');
INSERT INTO `currencies_zh-cn` VALUES (125, '索马里先令');
INSERT INTO `currencies_zh-cn` VALUES (126, '塞尔维亚第纳尔');
INSERT INTO `currencies_zh-cn` VALUES (127, '南苏丹镑');
INSERT INTO `currencies_zh-cn` VALUES (128, '苏里南元');
INSERT INTO `currencies_zh-cn` VALUES (129, '瑞典克朗');
INSERT INTO `currencies_zh-cn` VALUES (130, '斯威士兰埃马兰吉尼');
INSERT INTO `currencies_zh-cn` VALUES (131, '塞舌尔卢比');
INSERT INTO `currencies_zh-cn` VALUES (132, '叙利亚镑');
INSERT INTO `currencies_zh-cn` VALUES (133, '泰铢');
INSERT INTO `currencies_zh-cn` VALUES (134, '塔吉克斯坦索莫尼');
INSERT INTO `currencies_zh-cn` VALUES (135, '土库曼斯坦马纳特');
INSERT INTO `currencies_zh-cn` VALUES (136, '汤加潘加');
INSERT INTO `currencies_zh-cn` VALUES (137, '特立尼达元');
INSERT INTO `currencies_zh-cn` VALUES (138, '突尼斯第纳尔');
INSERT INTO `currencies_zh-cn` VALUES (139, '土耳其里拉');
INSERT INTO `currencies_zh-cn` VALUES (140, '新台币');
INSERT INTO `currencies_zh-cn` VALUES (141, '坦桑尼亚先令');
INSERT INTO `currencies_zh-cn` VALUES (142, '乌干达先令');
INSERT INTO `currencies_zh-cn` VALUES (143, '乌克兰格里夫纳');
INSERT INTO `currencies_zh-cn` VALUES (144, '乌拉圭比索');
INSERT INTO `currencies_zh-cn` VALUES (145, '乌兹别克斯坦索姆');
INSERT INTO `currencies_zh-cn` VALUES (146, '委内瑞拉玻利瓦尔');
INSERT INTO `currencies_zh-cn` VALUES (147, '越南盾');
INSERT INTO `currencies_zh-cn` VALUES (148, '尼瓦努阿图瓦图');
INSERT INTO `currencies_zh-cn` VALUES (149, '萨摩亚塔拉');
INSERT INTO `currencies_zh-cn` VALUES (150, '也门里亚尔');
INSERT INTO `currencies_zh-cn` VALUES (151, '赞比亚克瓦查');
INSERT INTO `currencies_zh-cn` VALUES (152, '津巴布韦元');
INSERT INTO `currencies_zh-cn` VALUES (153, '中非金融共同体法郎');
INSERT INTO `currencies_zh-cn` VALUES (154, '多布拉');
INSERT INTO `currencies_zh-cn` VALUES (155, '圣赫勒拿镑');
INSERT INTO `currencies_zh-cn` VALUES (156, '刚果法郎');
INSERT INTO `currencies_zh-cn` VALUES (157, '白俄罗斯卢布');
INSERT INTO `currencies_zh-cn` VALUES (158, '南极元');

-- ----------------------------
-- Table structure for languages
-- ----------------------------
DROP TABLE IF EXISTS `languages`;
CREATE TABLE `languages`  (
  `id` smallint(1) UNSIGNED NOT NULL AUTO_INCREMENT,
  `iso` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `sort` smallint(1) UNSIGNED NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `iso`(`iso`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of languages
-- ----------------------------
INSERT INTO `languages` VALUES (1, 'en', 'English', 999);
INSERT INTO `languages` VALUES (2, 'zh-CN', '简体中文', 998);
INSERT INTO `languages` VALUES (3, 'zh-TW', '繁体中文', 0);
INSERT INTO `languages` VALUES (4, 'ja', 'やまと', 0);
INSERT INTO `languages` VALUES (5, 'it', 'Italiano', 0);
INSERT INTO `languages` VALUES (6, 'fr', 'Français', 0);
INSERT INTO `languages` VALUES (7, 'de', 'Deutsch', 0);
INSERT INTO `languages` VALUES (8, 'ru', 'Русский', 0);

-- ----------------------------
-- Table structure for regions
-- ----------------------------
DROP TABLE IF EXISTS `regions`;
CREATE TABLE `regions`  (
  `id` smallint(1) UNSIGNED NOT NULL AUTO_INCREMENT,
  `lang` smallint(1) UNSIGNED NOT NULL DEFAULT 1,
  `pid` smallint(1) UNSIGNED NULL DEFAULT 0,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`, `lang`) USING BTREE,
  INDEX `id`(`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 29 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of regions
-- ----------------------------
INSERT INTO `regions` VALUES (1, 0, 0, 'Africa');
INSERT INTO `regions` VALUES (2, 0, 0, 'Americas');
INSERT INTO `regions` VALUES (3, 0, 0, 'Asia');
INSERT INTO `regions` VALUES (4, 0, 0, 'Europe');
INSERT INTO `regions` VALUES (5, 0, 0, 'Oceania');
INSERT INTO `regions` VALUES (6, 0, 0, 'Polar');
INSERT INTO `regions` VALUES (7, 0, 5, 'Australia and New Zealand');
INSERT INTO `regions` VALUES (8, 0, 2, 'Caribbean');
INSERT INTO `regions` VALUES (9, 0, 2, 'Central America');
INSERT INTO `regions` VALUES (10, 0, 3, 'Central Asia');
INSERT INTO `regions` VALUES (11, 0, 1, 'Eastern Africa');
INSERT INTO `regions` VALUES (12, 0, 3, 'Eastern Asia');
INSERT INTO `regions` VALUES (13, 0, 4, 'Eastern Europe');
INSERT INTO `regions` VALUES (14, 0, 5, 'Melanesia');
INSERT INTO `regions` VALUES (15, 0, 5, 'Micronesia');
INSERT INTO `regions` VALUES (16, 0, 1, 'Middle Africa');
INSERT INTO `regions` VALUES (17, 0, 1, 'Northern Africa');
INSERT INTO `regions` VALUES (18, 0, 2, 'Northern America');
INSERT INTO `regions` VALUES (19, 0, 4, 'Northern Europe');
INSERT INTO `regions` VALUES (20, 0, 5, 'Polynesia');
INSERT INTO `regions` VALUES (21, 0, 2, 'South America');
INSERT INTO `regions` VALUES (22, 0, 3, 'South-Eastern Asia');
INSERT INTO `regions` VALUES (23, 0, 1, 'Southern Africa');
INSERT INTO `regions` VALUES (24, 0, 3, 'Southern Asia');
INSERT INTO `regions` VALUES (25, 0, 4, 'Southern Europe');
INSERT INTO `regions` VALUES (26, 0, 1, 'Western Africa');
INSERT INTO `regions` VALUES (27, 0, 3, 'Western Asia');
INSERT INTO `regions` VALUES (28, 0, 4, 'Western Europe');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` bigint(1) UNSIGNED NOT NULL AUTO_INCREMENT,
  `pid` bigint(1) UNSIGNED NULL DEFAULT 0 COMMENT '推荐人id',
  `account` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '登录账号,有则唯一',
  `mail` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '邮箱号,有则唯一',
  `phone` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '手机号,有则唯一',
  `mailvery` tinyint(1) NULL DEFAULT 0 COMMENT '邮箱是否验证,1为已验证',
  `phonevery` tinyint(1) NULL DEFAULT 0 COMMENT '手机是否验证,1为已验证',
  `pwd` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '密码',
  `nickname` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '昵称',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '头像地址',
  `addtime` int(1) UNSIGNED NULL DEFAULT 0 COMMENT '注册时间',
  `status` tinyint(1) UNSIGNED NULL DEFAULT 1 COMMENT '账号状态,1为正常,其他值均为不正常',
  `sex` tinyint(1) NULL DEFAULT 0 COMMENT '性别,0保密,1男，2女',
  `height` tinyint(1) UNSIGNED NULL DEFAULT 0 COMMENT '身高cm',
  `weight` float(5, 2) UNSIGNED NULL DEFAULT 0.00 COMMENT '体重kg',
  `birth` int(1) UNSIGNED NULL DEFAULT NULL COMMENT '生日',
  `age` tinyint(1) UNSIGNED NULL DEFAULT NULL COMMENT '年龄',
  `job` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '职业',
  `income` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '收入',
  `emotion` tinyint(1) UNSIGNED NULL DEFAULT 0 COMMENT '情感状态',
  `star` tinyint(1) UNSIGNED NULL DEFAULT 0 COMMENT '星座',
  `ip` int(1) UNSIGNED NULL DEFAULT NULL COMMENT '注册时的ipv4地址',
  `country` smallint(1) UNSIGNED NULL DEFAULT 0 COMMENT '国家id',
  `city` smallint(1) NULL DEFAULT NULL COMMENT '城市id',
  `singleid` tinyint(1) UNSIGNED NULL DEFAULT 0 COMMENT '单点登录token id',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `account`(`account`) USING BTREE,
  UNIQUE INDEX `mail`(`mail`) USING BTREE,
  UNIQUE INDEX `phone`(`phone`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, 1111111111111111111, 'aaa', NULL, NULL, 0, NULL, '$2y$10$YtTGQpR89tBaHBJAR6uKsuIylIcGOBtER336oFPCq1i4q39uP32xW', NULL, '', 1654743271, 1, 0, 254, 900.56, NULL, NULL, NULL, NULL, NULL, NULL, 4294967295, NULL, NULL, NULL);
INSERT INTO `users` VALUES (2, 0, 'bbb', NULL, NULL, 0, 0, NULL, NULL, NULL, 0, 1, 0, 0, 0.00, NULL, NULL, NULL, NULL, 0, 0, NULL, 0, NULL, NULL);
INSERT INTO `users` VALUES (3, 0, '', 'dfsf@qc.cc', '', 1, 0, '$2a$10$s1RIiHFkcubdu7TeoWJjPeFhQ7v1PhZOhCrA754bXCYSept.dGOFi', '', '', 0, 0, 0, 0, 0.00, 0, 0, '', '', 0, 0, 0, 0, NULL, NULL);
INSERT INTO `users` VALUES (4, 0, NULL, 'dfsf@qc.ccz', NULL, 0, 0, '$2a$10$XvY0aIuaV0WUClH3uSDWYuzuZlnZoJqBKFK7HU85mvBzB5i5Fxe7S', NULL, NULL, 0, 1, 0, 0, 0.00, NULL, NULL, NULL, NULL, 0, 0, 2130706433, 0, NULL, NULL);
INSERT INTO `users` VALUES (5, 0, NULL, 'dfsfzz@qc.ccz', NULL, 0, 0, '$2a$10$nQyk5E6pXEucHryCYo9Rceyy2SBVb25auT8gHHNB1NaK1YSo4VgKm', NULL, NULL, 0, 1, 0, 0, 0.00, NULL, NULL, NULL, NULL, 0, 0, 2130706433, 0, NULL, NULL);
INSERT INTO `users` VALUES (6, 0, NULL, NULL, NULL, 0, 0, '$2a$10$Jx0HApS0UlgqpXndLJhCf.h3GZoP0/rm2hMkFtd1ZzeVY71dWL0Xi', NULL, NULL, 0, 1, 0, 0, 0.00, NULL, NULL, NULL, NULL, 0, 0, 2130706433, 0, NULL, NULL);
INSERT INTO `users` VALUES (7, 0, NULL, NULL, NULL, 0, 0, '$2a$10$Q4rL/UbA5TMN5IlZ66oDmOc3ZxrJGZqVYTCigiPW2UnGVLf7n2p2u', NULL, NULL, 0, 1, 0, 0, 0.00, NULL, NULL, NULL, NULL, 0, 0, 2130706433, 0, NULL, NULL);
INSERT INTO `users` VALUES (8, 0, NULL, NULL, NULL, 0, 0, NULL, NULL, NULL, 1654856133, 1, 0, 0, 0.00, NULL, NULL, NULL, NULL, 0, 0, 2130706433, 0, NULL, NULL);
INSERT INTO `users` VALUES (9, 0, NULL, NULL, NULL, 0, 0, NULL, NULL, NULL, 1654856143, 1, 0, 0, 0.00, NULL, NULL, NULL, NULL, 0, 0, 2130706433, 0, NULL, NULL);

SET FOREIGN_KEY_CHECKS = 1;
