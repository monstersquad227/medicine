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

 Date: 04/06/2025 13:10:02
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for medicine_plan_record
-- ----------------------------
DROP TABLE IF EXISTS `medicine_plan_record`;
CREATE TABLE `medicine_plan_record` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `user_id` int(10) unsigned NOT NULL COMMENT '用户ID',
  `plan_id` int(10) unsigned DEFAULT NULL COMMENT '用药计划ID，关联 medicine_plan.id',
  `medicine_name` varchar(255) NOT NULL COMMENT '药物名称',
  `actual_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '实际用药时间',
  `memo` text COMMENT '打卡备注信息',
  `is_checked` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否打卡（0=未打卡，1=已打卡）',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否异常（0: 正常；1: 异常），与计划时间相差15分钟记为异常',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_plan_id` (`plan_id`),
  CONSTRAINT `fk_medicine_plan_record_medicine_plan` FOREIGN KEY (`plan_id`) REFERENCES `medicine_plan` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=101 DEFAULT CHARSET=utf8mb4 COMMENT='用药打卡记录表';

SET FOREIGN_KEY_CHECKS = 1;
