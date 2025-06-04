/*
 Navicat Premium Dump SQL

 Source Server         : 47.103.98.61
 Source Server Type    : MySQL
 Source Server Version : 50736 (5.7.36)
 Source Host           : 47.103.98.61:3306
 Source Schema         : medicine

 Target Server Type    : MySQL
 Target Server Version : 50736 (5.7.36)
 File Encoding         : 65001

 Date: 04/06/2025 13:10:15
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `nickname` varchar(100) NOT NULL DEFAULT '' COMMENT '用户昵称',
  `image` varchar(255) DEFAULT NULL COMMENT '用户头像 URL',
  `phone_num` varchar(20) NOT NULL COMMENT '手机号',
  `huawei_id` varchar(100) DEFAULT NULL COMMENT '华为用户ID',
  `password` varchar(255) NOT NULL COMMENT '用户密码（加密存储）',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_phone_num` (`phone_num`),
  KEY `idx_huawei_id` (`huawei_id`)
) ENGINE=InnoDB AUTO_INCREMENT=101 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` (`id`, `nickname`, `image`, `phone_num`, `huawei_id`, `password`, `created_at`, `updated_at`) VALUES (1, 'HarmonyOS_2', '', '15056332824', NULL, 'mLf0p3ymtc38YTNyr3+U1bZ3xvSrq6dLPg+ecdTNDOeq2w==', '2025-05-27 10:23:59', '2025-05-29 17:49:28');
INSERT INTO `user` (`id`, `nickname`, `image`, `phone_num`, `huawei_id`, `password`, `created_at`, `updated_at`) VALUES (2, 'HarmonyOS_1', NULL, '13611965064', NULL, 'mLf0p3ymtc38YTNyr3+U1bZ3xvSrq6dLPg+ecdTNDOeq2w==', '2025-05-29 15:45:52', '2025-05-29 15:45:55');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
