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

 Date: 28/05/2025 15:53:18
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for medicine_plan
-- ----------------------------
DROP TABLE IF EXISTS `medicine_plan`;
CREATE TABLE `medicine_plan` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `medicine_id` int unsigned NOT NULL COMMENT '用药方案ID，关联 medicine_course.id',
  `amount` int NOT NULL COMMENT '用药数量',
  `type` varchar(16) NOT NULL COMMENT '剂量单位',
  `plan_time` varchar(64) NOT NULL COMMENT '用药时间',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_medicine_id` (`medicine_id`),
  CONSTRAINT `fk_medicine_plan_course` FOREIGN KEY (`medicine_id`) REFERENCES `medicine_course` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用药计划明细表';

SET FOREIGN_KEY_CHECKS = 1;
