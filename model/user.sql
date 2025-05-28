/*
 Navicat Premium Dump SQL

 Source Server         : 192.168.1.87
 Source Server Type    : MySQL
 Source Server Version : 80031 (8.0.31)
 Source Host           : 192.168.1.87:3307
 Source Schema         : medicine

 Target Server Type    : MySQL
 Target Server Version : 80031 (8.0.31)
 File Encoding         : 65001

 Date: 28/05/2025 15:53:31
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户表';

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` (`id`, `nickname`, `image`, `phone_num`, `huawei_id`, `password`, `created_at`, `updated_at`) VALUES (1, 'HarmonyOS2', '', '15056332824', NULL, 'mLf0p3ymtc38YTNyr3+U1bZ3xvSrq6dLPg+ecdTNDOeq2w==', '2025-05-27 10:23:59', '2025-05-28 09:00:50');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
