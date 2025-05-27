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

 Date: 27/05/2025 16:18:37
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for medicine_plan_record
-- ----------------------------
DROP TABLE IF EXISTS `medicine_plan_record`;
CREATE TABLE `medicine_plan_record` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
  `medicine_name` varchar(255) NOT NULL COMMENT '药物名称',
  `actual_time` datetime NOT NULL COMMENT '实际用药时间',
  `memo` text COMMENT '打卡备注信息',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否异常（0: 正常；1: 异常），与计划时间相差15分钟记为异常',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用药打卡记录表';

-- ----------------------------
-- Records of medicine_plan_record
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
