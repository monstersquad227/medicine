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

 Date: 28/05/2025 15:53:10
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for medicine_course
-- ----------------------------
DROP TABLE IF EXISTS `medicine_course`;
CREATE TABLE `medicine_course` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `user_id` int unsigned NOT NULL COMMENT '用户ID',
  `medicine_name` varchar(255) NOT NULL COMMENT '药物名称',
  `medicine_image` varchar(512) DEFAULT NULL COMMENT '药物图片',
  `medicine_type` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '药物方式：0-内服；1-外用',
  `medicine_timing` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '用药时机：0-不限；1-饭前用药；2-饭后用药；3-随餐用药；4-睡前用药',
  `course_start_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '用药开始时间',
  `status` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '方案状态：0-生效；1-废弃',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用药方案表';

SET FOREIGN_KEY_CHECKS = 1;
